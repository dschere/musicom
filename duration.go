package main


const (
	Whole      float32 = 4
	Half           = 2
	Quarter        = 1
	Eighth         = 0.5
	Sixteenth      = 0.25
	ThirtySec      = 0.125
	Sixtyforth     = 0.0625
)

type Duration struct {
	dotted  bool
	v       float32
}


type Legato struct {
    enabled bool
}
var LegatoOn = Legato{enabled: true}
var LegatoOff = Legato{enabled: true}
 

type Staccato struct {
	enabled bool
}
var StaccatoOn = Staccato{enabled: true}
var StaccatoOff = Staccato{enabled: true}


var Dw  = Duration{v: Whole}
var Dh  = Duration{v: Half}
var Dq  = Duration{v: Quarter}
var De  = Duration{v: Eighth}
var Ds  = Duration{v: Sixteenth}
var Dt  = Duration{v: ThirtySec}
var Dx  = Duration{v: Sixtyforth}

func Dot(d Duration) Duration {
    d.dotted = true;
    return d;
}


type TripletDuration struct {
    v bool
}

var StartTriplet = TripletDuration{v: true}
var EndTriplet   = TripletDuration{v: false}

