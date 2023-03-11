package main

import (
    "time"
)


// builtin effects
/*
const BEND    = "bend"
const VIBRATO = "vibrato"
const SLIDE   = "slide"
const REVERB  = "reverb"
const CHORUS  = "chorus"
const PAN     = "pan" 
*/


const (
   EFF_CUSTOM = -1
   BEND_HALF = 0
   BEND_WHOLE = 1
   BEND_RELEASE_HALF = 2
   BEND_RELEASE_WHOLE = 3
   PREBEND_RELEASE_HALF = 4
   PREBEND_RELEASE_WHOLE = 5
   
   REVERB = 6
)

const BEND_TIME_SLICE_RATIO_DENOMITOR = 3

type CustomEffectCb func (seq *Sequencer, c *Channel, te* TickEvent)


type Effect struct {
    op int
    custom_before_note CustomEffectCb 
    custom_after_note  CustomEffectCb 

    reverb_val float32
}

func Reverb(val float32) *Effect {
    e := &Effect{ op: REVERB }
    e.reverb_val = val
    return e
}

func (e *Effect) execute(seq *Sequencer, c *Channel) {
    switch e.op {
        case REVERB:
            seq.s.fs.reverb( c.midi_chan, e.reverb_val)
    }
}

func (e *Effect) before_note(seq *Sequencer, c *Channel, te* TickEvent){
    switch e.op {
        case EFF_CUSTOM:
            if e.custom_before_note != nil {
                e.custom_before_note(seq, c, te)
            } 
        case PREBEND_RELEASE_HALF:
            seq.s.fs.bend( c.midi_chan, 1.0 )
        case PREBEND_RELEASE_WHOLE:
            seq.s.fs.bend( c.midi_chan, 2.0 )        
    }
}

func _bend(seq *Sequencer, c *Channel, te* TickEvent, start float32, end float32) time.Duration {
    num_time_slices := 13
    d := te.duration / time.Duration(BEND_TIME_SLICE_RATIO_DENOMITOR * num_time_slices)
    step := (end - start) / float32(num_time_slices)
    
    val := start
    
    for i := 0; i < num_time_slices; i ++ {
        val += step
        seq.s.fs.bend( c.midi_chan, val )
        time.Sleep( d * time.Millisecond ) 
    }
    
    return te.duration / BEND_TIME_SLICE_RATIO_DENOMITOR
}

func (e *Effect) after_note(seq *Sequencer, c *Channel, te* TickEvent) {
    
    
    switch e.op {
        case EFF_CUSTOM:
            if e.custom_after_note != nil {
                e.custom_after_note(seq, c, te)
            }
        case PREBEND_RELEASE_HALF:
            d := te.duration 
            d -= _bend(seq, c, te, 1.0, 0.0)
            time.Sleep(d * time.Millisecond)
            seq.s.fs.bend( c.midi_chan, 0.0)
            
        case PREBEND_RELEASE_WHOLE:
            d := te.duration 
            d -= _bend(seq, c, te, 2.0, 0.0)
            time.Sleep(d * time.Millisecond)
            seq.s.fs.bend( c.midi_chan, 0.0)

        case BEND_HALF:
            d := te.duration 
            d -= _bend(seq, c, te, 0.0, 1.0)
            time.Sleep(d * time.Millisecond)
            seq.s.fs.bend( c.midi_chan, 0.0)
            
        case BEND_WHOLE:
            d := te.duration 
            d -= _bend(seq, c, te, 0.0, 2.0)
            time.Sleep(d * time.Millisecond)
            seq.s.fs.bend( c.midi_chan, 0.0)
   
        case BEND_RELEASE_HALF:
            d := te.duration
            d3 := d/time.Duration(BEND_TIME_SLICE_RATIO_DENOMITOR) 
            _bend(seq, c, te, 0.0, 1.0)
            time.Sleep(d3 * time.Millisecond)
            _bend(seq, c, te, 1.0, 0.0)
            seq.s.fs.bend( c.midi_chan, 0.0)

        case BEND_RELEASE_WHOLE:
            d := te.duration
            d3 := d/time.Duration(BEND_TIME_SLICE_RATIO_DENOMITOR) 
            _bend(seq, c, te, 0.0, 2.0)
            time.Sleep(d3 * time.Millisecond)
            _bend(seq, c, te, 2.0, 0.0)
            seq.s.fs.bend( c.midi_chan, 0.0)
                 
    }
}





