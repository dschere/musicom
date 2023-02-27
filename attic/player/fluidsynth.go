/**
fluid synth controller. 

Maintains a TCP connection to fluid synth.
Abstracts the low level commands to create higher 
 level commands.
*/
package player

import (
    "fmt"
    "time"
    "net"
    "log"
)



func (player *TrackPlayer) set_instrument(midi_instrument_code int) error {
    message := fmt.Sprintf("prog 0 %d\n",midi_instrument_code) 
    return player.send_fluidsynth_cmd(message)    
}

func (player *TrackPlayer) noteon(midi_note_code int, dynamic int) error {
    message := fmt.Sprintf("noteon 0 %d %d\n",midi_note_code, dynamic) 
    return player.send_fluidsynth_cmd(message)    
} 

func (player *TrackPlayer) quit() error {
    return player.send_fluidsynth_cmd("quit\n")    
}

// coroutine that interfaces with the fluidsynth that
// handles midi events uses ladspa effects.
func (player *TrackPlayer) send_fluidsynth_cmd(message string) error {
    if (player.fsError != nil) {
        return player.fsError
    } 
    player.msgQueue <- message 
    return nil
}

func (player *TrackPlayer) shutdown_fluidsynth()  {
    if (player.fsError == nil) {
        player.msgQueue <- "quit\n"
        <-player.onCmdShutdown     
    }
    // quit will do a graceful shutdown, this might fail because
    // the process may have exited.
    player.proc.Stop()

    player.fsError = nil
    player.proc = nil
}    
    
func handle_fluidsynth(player *TrackPlayer, port int) {
    var conn *net.TCPConn
    connected := false
    servAddr := fmt.Sprintf("localhost:%d", port)
    retryConnCount := 0
    
    player.fsError = nil
    
    tcpAddr, err := net.ResolveTCPAddr("tcp", servAddr)
    if err != nil {
        log.Panic(err)
        player.fsError = err
        return
    }

    fmt.Printf("Trying to connect to %s\n", servAddr)
    for connected == false {
        conn, err = net.DialTCP("tcp", nil, tcpAddr)
        if err == nil {
            connected = true
        } else {
            time.Sleep(MAX_FLUIDSYNTH_RETRY_INTERVAL); // waut for fluidsynth to come up.
            retryConnCount++;
            
            if retryConnCount > MAX_FLUIDSYNTH_RETRY {
                log.Panic(err)
                player.fsError = err
                return
            }
        }
    }
    fmt.Printf("Connected to %s\n", servAddr)
    
    // close when coroutine ends 
    defer conn.Close()
    
    // load default soundfont
    conn.Write([]byte( "load " + DefaultSoundFontFile + "\n" ))
    
    
    // while connected service buffered event queue
    for connected == true {
        message := <- player.msgQueue
        _, err := conn.Write([]byte( message ))
        
        if message == "quit\n" {
            fmt.Printf("quit received for player[%d]\n", player.tracknum)
            connected = false
        } else if err != nil {
            log.Panic(err)
            player.fsError = err
            return
        }
    }
    player.proc.Stop()
    
    // graceful shutdown
    player.onCmdShutdown <- true
}

