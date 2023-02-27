/*
   Interface for a child process
   
   * midi is handled by fluidsynth
   * live audio ffplay

For ladspa and fluidsynth: 
   
https://github.com/FluidSynth/fluidsynth/blob/master/doc/ladspa.md
*/
package player

import (
    "log"
    "fmt"
    "os/exec"
    "github.com/go-cmd/cmd"
    "strconv"
    "time"
)

const MidiTrackPlayer      = "/usr/local/bin/miditrackplayer.sh"
const DefaultSoundFontFile = "/data/sound-fonts/27mg_Symphony_Hall_Bank.SF2"
const StartingPortNumber   = 10000

const QUEUE_DEPTH = 10;
const MAX_FLUIDSYNTH_RETRY = 15
const MAX_FLUIDSYNTH_RETRY_INTERVAL = time.Second

const (
	MIDI_TRACK int  = 0
	AUDIO_TRACK     = 1
	MODULATOR_TRACK = 2
	LIVE_TRACK      = 3
)

type TrackPlayer struct {
    proc      *cmd.Cmd     // child process doing the work. 
    tracknum  int           // track in song this player is dedicated to
    tracktype int           // live|midi     
    ladspa_enabled bool     // update ladspa or enable ?
    msgQueue  chan string   // sending data to fluidsyn via TCP, I want no delay
    fsError   error         // flag and message that fluid synth coroutine is dead.
    onCmdShutdown chan bool // event used to synnchonize shutdown of child process
}


func (player *TrackPlayer) _launch_child_process() int {
    path := ""
    
    switch player.tracktype {
        case MIDI_TRACK:
            path = MidiTrackPlayer        
	//case AUDIO_TRACK
	//case MODULATOR_TRACK
	//case LIVE_TRACK
        default:
            log.Panic("Unknown track type")
            return -1
    }
    path, err := exec.LookPath(path)
    if err != nil {
        log.Panic(err)
        return -1
    }
    port := StartingPortNumber + player.tracknum
    
    player.proc = cmd.NewCmd(
        path,
        strconv.Itoa(port),
        DefaultSoundFontFile,  
    )
    player.proc.Start()
    fmt.Printf("Starting fluidsynth on port %d\n", port)
    
    player.msgQueue = make(chan string, QUEUE_DEPTH)
    player.onCmdShutdown = make(chan bool)
    
    if player.tracktype == MIDI_TRACK {
        go handle_fluidsynth(player, port)
    }
    return 0
}

func (player *TrackPlayer) init(tracknum int, tracktype int) int {
    player.tracknum  = tracknum
    player.tracktype = tracktype
    
    switch player.tracktype {
        case MIDI_TRACK:
            player.msgQueue = make(chan string, 10)
            player.ladspa_enabled = false
    } 
    
    return player._launch_child_process()
}







