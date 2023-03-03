package main

import (
   "fmt"
   "time"
)

type TickEvent struct {
    ref interface{}
    duration time.Duration
    moment time.Duration // current tick value in channel
}

type Channel struct {
    // event tick -> list of events for that tick
    events []TickEvent
    midi_chan int
    last_tick time.Duration
    note_duration Duration
    triplet bool
     
    t Track
}

func (c *Channel) addEvent(te TickEvent) {
    te.moment = c.last_tick
    c.events = append(c.events, te)
    c.last_tick += te.duration
}



type Sequencer struct {
    channels map[string]Channel
    s* Song_ //reduce the number of args that have to be passed around
    bpm int
    ts TimeSignature
}


func _compute_duration(seq *Sequencer, c *Channel) time.Duration {
    incr := c.note_duration.v * (60.0 / float32(seq.bpm))
    if c.note_duration.dotted == true {
        incr += (incr/2)
    }
    if c.triplet == true {
        incr *= 0.66
    }
    
    // duration expressed as a millisecond integer
    return time.Duration(incr * 1000)
}

func _process_channel_event(seq *Sequencer, c *Channel, i interface{}){
    switch i.(type) {
        case TimeSignature:
            seq.ts = i.(TimeSignature)
        case Duration:
            c.note_duration = i.(Duration)
        case TripletDuration:
            c.triplet = i.(TripletDuration).v
        default:
            te := TickEvent{ref: i}
            te.duration = _compute_duration(seq, c)
            c.addEvent(te)
    }
}

func _process_channel(seq *Sequencer, c *Channel,ilist []interface{}){
    // compute the tick value for each event
    // this is the millisecond value when a note or effect is to be played
    // or stopped.
    for i := 0; i < len(ilist); i++ {
        _process_channel_event(seq , c, ilist[i])
    }      
}

func _process_measure(seq *Sequencer, m Measure) {
    if m.repeat == true {
        _repeat_loop(seq, m.start_measure_id, m.id, m.count)
    } else {
        for trackName, ilist := range m.cmds {
            c := seq.channels[trackName]
            _process_channel(seq, &c, ilist)
            seq.channels[trackName] = c
        }
    }     
}

func _repeat_loop(seq *Sequencer, s_measure_id int, e_measure_id int, count int) {
    //fmt.Printf("s_measure_id %d, e_measure_id %d, count %d\n",  
    //     s_measure_id , e_measure_id , count)
         
    for c := 0; c < count; c++ {
        for i := s_measure_id; i < e_measure_id; i++ {
            m := seq.s.mlist[i]
            _process_measure(seq, m)
            seq.s.mlist[i] = m
        }
    }
}


func _create_channels(seq *Sequencer) {
    seq.channels = make(map[string]Channel)

    channel := 0
    
    // assign midi channels per track, normally this 1/1 however in the
    // case of a distortion guitar its multiple channels.
    //   distorion + muted and harmonics channels 
    for tname, tval := range seq.s.trackTable {
        c := Channel{}
        c.midi_chan = channel
        c.t = tval
        c.last_tick = 0
        c.note_duration = Dw
        c.triplet = false    
        c.events = make([]TickEvent, 0)
        
        seq.channels[tname] = c    
        if tval.instrument == DISTORTION_GUITAR || 
           tval.instrument == OVERDRIVEN_GUITAR   {
            // channel + 1 -> muted
            // channel + 2 -> harmonics
            channel += 3       
        } else {
            channel += 1
        }
    }
}

/*
 * Use the song data structure to generate a sequence of events
 * that will be converted to commands for fluidsynth. 
 */
func (seq *Sequencer) compile(s *Song_) {  
    seq.s = s  
    seq.bpm = s.bpm
    seq.ts = s.ts
    _create_channels(seq)
    // treat the entire peace as inside a single repeat
    // loop
    _repeat_loop(seq , 0, len(s.mlist), 1)
} 

func (seq *Sequencer) pretty_print() {
    fmt.Printf("song bpm=%s, ts=%d/%d\n", seq.bpm, 
    	seq.ts.beatsPerBar, seq.ts.beatUnit)
    for name, channel := range seq.channels {
        fmt.Printf("    %s track{ %p }\n", name, channel.t )
        fmt.Printf("    %d\n", len(channel.events))
        for i := 0; i < len(channel.events); i++ {
            te := channel.events[i]
            fmt.Printf("        [%ld] duration=%d, %p \n", 
                te.moment, te.duration, te.ref )
        }
    }
}

