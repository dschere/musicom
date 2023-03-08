package main

/*
 * fluid synth client 
 */

import (
    "fmt"
    "time"
    "net"
    "log"
)

type FluidSynth struct {
    port          int
    msgQueue chan string   // sending data to fluidsyn via TCP
    ready    chan bool
}

const QUEUE_DEPTH = 10
const MAX_FLUIDSYNTH_RETRY = 15
const MAX_FLUIDSYNTH_RETRY_INTERVAL = time.Second

const FLUIDSYNTH_PORT = 2112


func run_fluidsynth_client(fs *FluidSynth) {
    // establish TCP connection to daemon
    // launch coroutine to handle messages
    var conn *net.TCPConn
    connected := false
    servAddr := fmt.Sprintf("localhost:%d", fs.port)
    retryConnCount := 0
    
    tcpAddr, err := net.ResolveTCPAddr("tcp", servAddr)
    if err != nil {
        log.Fatal(err)
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
                log.Fatal(err)
                return
            }
        }
    }
    fmt.Printf("Connected to %s\n", servAddr)
    // unblock start()
    fs.ready <- true
    
    // close when coroutine ends 
    defer conn.Close()

    // while connected service buffered event queue
    for connected == true {
        message := <- fs.msgQueue
        _, err := conn.Write([]byte( message ))
        if err != nil {
            log.Panic(err)
            return
        }
    }
}

func (fs *FluidSynth) start() {
    // create message queue for outbound commands
    fs.msgQueue = make(chan string, QUEUE_DEPTH)
    fs.ready = make(chan bool, QUEUE_DEPTH)

    go run_fluidsynth_client(fs)
    
    // wait for connection to localhost 
    <- fs.ready    
}

func (fs *FluidSynth) set_instrument(chn int, midi_instrument_code int) {
    message := fmt.Sprintf("prog %d %d\n",chn, midi_instrument_code) 
    fmt.Printf("FluidSynth: set_instrument %s\n", message)
    fs.msgQueue <- message    
}

func (fs *FluidSynth) noteon(chn int, midi_note_code int, dynamic int)  {
    message := fmt.Sprintf("noteon %d %d %d\n",chn,midi_note_code, dynamic) 
    fmt.Printf("FluidSynth: noteon %s\n", message)
    fs.msgQueue <- message   
} 

func (fs *FluidSynth) noteoff(chn int, midi_note_code int)  {
    message := fmt.Sprintf("noteoff %d %d\n",chn,midi_note_code) 
    fmt.Printf("FluidSynth: noteoff %s\n", message)
    fs.msgQueue <- message   
} 
