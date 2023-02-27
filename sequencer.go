package main

import (
   "time"
)

type TickEvent struct {
    ref interface{}
    duration time.Duration
}

type Channel struct {
    // event tick -> list of events for that tick
    events map[time.Duration] []TickEvent
    midi_chan int
    last_tick time.Duration
    note_duration Duration
    triplet bool
     
    t Track
}

func (c *Channel) addEvent(te TickEvent) {
    teList, found := c.events[c.last_tick]
    if found == false {
        teList = make([]TickEvent,0)
    }
    teList = append(teList, te) 
    c.events[c.last_tick] = teList
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
        }
    }     
}

func _repeat_loop(seq *Sequencer, s_measure_id int, e_measure_id int, count int) {
    for c := 0; c < count; c++ {
        for i := s_measure_id; i < e_measure_id; i++ {
            m := seq.s.mlist[i]
            _process_measure(seq, m)
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
        c.events = make(map[time.Duration] []TickEvent)
        
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
    _repeat_loop(seq , 0, len(seq.channels), 1)
} 



