package main


const (
	Whole      int = 64
	Half           = 32
	Quarter        = 16
	Eighth         = 8
	Sixteenth      = 4
	ThirtySec      = 2
	Sixtyforth     = 1
)

type Duration struct {
	etype   string `default:"Duration"`	
	dotted  bool
	v       int
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

