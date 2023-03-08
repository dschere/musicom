package main

import (
    "log"
    "strconv"
)

type Note_ struct {
	rest  bool
	mcode int // midi value
	fret  int
	gstr  int
    
    // set to true and this will override any track dynamic
    // that is set. 
    custom_dynamic bool
    
    dynamic Dynamic

    effects []Effect
}


func NoteFromCode(mcode int) *Note_ {
    n := &Note_{}
    n.mcode = mcode
    n.custom_dynamic = false
    n.dynamic = MP
    n.effects = make([]Effect,0)
    n.rest = false
    return n
}

func Note(args ...int) *Note_ {
    n := &Note_{}
    
    n.effects = make([]Effect,0)
    if len(args) == 0 {
        n.rest = true
    } else {
        n.fret = args[0]
        n.gstr = args[1]
    }
    n.custom_dynamic = false
    n.dynamic = MP
    n.mcode = 0
    
    return n
}

var StepTable = map[string]int {
    "A": 0,
    "B": 2,
    "C": 3,
    "D": 5,
    "E": 7,
    "F": 8,
    "G": 10,
}


// use the tuning of the track and fret/gstr values to compute the
// mcode, if the mcode is zero (uninitialized)
func (n *Note_) ComputeMidiCode(opts TrackOpts) {
    tuning := ""

    if n.mcode == 0 {
        switch n.gstr {
            case 1:
                tuning = opts.string1
            case 2:
                tuning = opts.string2
            case 3:
                tuning = opts.string3
            case 4:
                tuning = opts.string4
            case 5:
                tuning = opts.string5
            case 6:
                tuning = opts.string6
        }
    }
    
    tlen := len(tuning)

    if tlen > 0 {        
        octave, err1 := strconv.Atoi(tuning[tlen-1:tlen]) 
        if err1 != nil {
            log.Fatal("Invalid octave value for tuning %s", tuning)
            return
        }
        step, found := StepTable[tuning[0:1]]
        if found == false {
            log.Fatal("Invalid note  value for tuning %s bust by A-G", 
                tuning)
            return
        }
        if tlen == 3 {
            switch tuning[1:2] {
                case "#":
                    step++
                case "b":
                    step--
                default:
                    log.Fatal("Invalid note value for tuning %s expected #/b", 
                        tuning)
                    return
            }
        }
        fret_0_value := (octave * 12) + step
        n.mcode = fret_0_value + n.fret 
    }
}


// change dynamic for this note.
func (n *Note_) d(dyn Dynamic) *Note_ {
    n.custom_dynamic = true
    n.dynamic = dyn
    return n
}

func (n *Note_) b(args ...int) *Note_ {
    e := Effect{}
    e.name = BEND
    e.i1 = 1
    
    if len(args) == 1 {
        e.i1 = args[1]
    }
    
    n.effects = append(n.effects, e)
    return n
}

func (n *Note_) v(args ...int) *Note_ {
    e := Effect{}
    e.name = VIBRATO
    e.i1 = 64 
    
    if len(args) == 1 {
        e.i1 = args[1]
    }
    
    n.effects = append(n.effects, e)
    return n
}

func (n *Note_) e(name string) *Note_ {
    // todo lookup effect append to effects list
    return n
}


// global presets

var R = Note()

var N1_0 = Note(0, 1)
var N1_1 = Note(1, 1)
var N1_2 = Note(2, 1)
var N1_3 = Note(3, 1)
var N1_4 = Note(4, 1)
var N1_5 = Note(5, 1)
var N1_6 = Note(6, 1)
var N1_7 = Note(7, 1)
var N1_8 = Note(8, 1)
var N1_9 = Note(9, 1)
var N1_10 = Note(10, 1)
var N1_11 = Note(11, 1)
var N1_12 = Note(12, 1)
var N1_13 = Note(13, 1)
var N1_14 = Note(14, 1)
var N1_15 = Note(15, 1)
var N1_16 = Note(16, 1)
var N1_17 = Note(17, 1)
var N1_18 = Note(18, 1)
var N1_19 = Note(19, 1)
var N1_20 = Note(20, 1)
var N1_21 = Note(21, 1)
var N1_22 = Note(22, 1)
var N1_23 = Note(23, 1)
var N1_24 = Note(24, 1)
var N2_0 = Note(0, 2)
var N2_1 = Note(1, 2)
var N2_2 = Note(2, 2)
var N2_3 = Note(3, 2)
var N2_4 = Note(4, 2)
var N2_5 = Note(5, 2)
var N2_6 = Note(6, 2)
var N2_7 = Note(7, 2)
var N2_8 = Note(8, 2)
var N2_9 = Note(9, 2)
var N2_10 = Note(10, 2)
var N2_11 = Note(11, 2)
var N2_12 = Note(12, 2)
var N2_13 = Note(13, 2)
var N2_14 = Note(14, 2)
var N2_15 = Note(15, 2)
var N2_16 = Note(16, 2)
var N2_17 = Note(17, 2)
var N2_18 = Note(18, 2)
var N2_19 = Note(19, 2)
var N2_20 = Note(20, 2)
var N2_21 = Note(21, 2)
var N2_22 = Note(22, 2)
var N2_23 = Note(23, 2)
var N2_24 = Note(24, 2)
var N3_0 = Note(0, 3)
var N3_1 = Note(1, 3)
var N3_2 = Note(2, 3)
var N3_3 = Note(3, 3)
var N3_4 = Note(4, 3)
var N3_5 = Note(5, 3)
var N3_6 = Note(6, 3)
var N3_7 = Note(7, 3)
var N3_8 = Note(8, 3)
var N3_9 = Note(9, 3)
var N3_10 = Note(10, 3)
var N3_11 = Note(11, 3)
var N3_12 = Note(12, 3)
var N3_13 = Note(13, 3)
var N3_14 = Note(14, 3)
var N3_15 = Note(15, 3)
var N3_16 = Note(16, 3)
var N3_17 = Note(17, 3)
var N3_18 = Note(18, 3)
var N3_19 = Note(19, 3)
var N3_20 = Note(20, 3)
var N3_21 = Note(21, 3)
var N3_22 = Note(22, 3)
var N3_23 = Note(23, 3)
var N3_24 = Note(24, 3)
var N4_0 = Note(0, 4)
var N4_1 = Note(1, 4)
var N4_2 = Note(2, 4)
var N4_3 = Note(3, 4)
var N4_4 = Note(4, 4)
var N4_5 = Note(5, 4)
var N4_6 = Note(6, 4)
var N4_7 = Note(7, 4)
var N4_8 = Note(8, 4)
var N4_9 = Note(9, 4)
var N4_10 = Note(10, 4)
var N4_11 = Note(11, 4)
var N4_12 = Note(12, 4)
var N4_13 = Note(13, 4)
var N4_14 = Note(14, 4)
var N4_15 = Note(15, 4)
var N4_16 = Note(16, 4)
var N4_17 = Note(17, 4)
var N4_18 = Note(18, 4)
var N4_19 = Note(19, 4)
var N4_20 = Note(20, 4)
var N4_21 = Note(21, 4)
var N4_22 = Note(22, 4)
var N4_23 = Note(23, 4)
var N4_24 = Note(24, 4)
var N5_0 = Note(0, 5)
var N5_1 = Note(1, 5)
var N5_2 = Note(2, 5)
var N5_3 = Note(3, 5)
var N5_4 = Note(4, 5)
var N5_5 = Note(5, 5)
var N5_6 = Note(6, 5)
var N5_7 = Note(7, 5)
var N5_8 = Note(8, 5)
var N5_9 = Note(9, 5)
var N5_10 = Note(10, 5)
var N5_11 = Note(11, 5)
var N5_12 = Note(12, 5)
var N5_13 = Note(13, 5)
var N5_14 = Note(14, 5)
var N5_15 = Note(15, 5)
var N5_16 = Note(16, 5)
var N5_17 = Note(17, 5)
var N5_18 = Note(18, 5)
var N5_19 = Note(19, 5)
var N5_20 = Note(20, 5)
var N5_21 = Note(21, 5)
var N5_22 = Note(22, 5)
var N5_23 = Note(23, 5)
var N5_24 = Note(24, 5)
var N6_0 = Note(0, 6)
var N6_1 = Note(1, 6)
var N6_2 = Note(2, 6)
var N6_3 = Note(3, 6)
var N6_4 = Note(4, 6)
var N6_5 = Note(5, 6)
var N6_6 = Note(6, 6)
var N6_7 = Note(7, 6)
var N6_8 = Note(8, 6)
var N6_9 = Note(9, 6)
var N6_10 = Note(10, 6)
var N6_11 = Note(11, 6)
var N6_12 = Note(12, 6)
var N6_13 = Note(13, 6)
var N6_14 = Note(14, 6)
var N6_15 = Note(15, 6)
var N6_16 = Note(16, 6)
var N6_17 = Note(17, 6)
var N6_18 = Note(18, 6)
var N6_19 = Note(19, 6)
var N6_20 = Note(20, 6)
var N6_21 = Note(21, 6)
var N6_22 = Note(22, 6)
var N6_23 = Note(23, 6)
var N6_24 = Note(24, 6)

var  BASS_DRUM_1 = NoteFromCode(36)
var  SIDE_STICK = NoteFromCode(37)
var  ACOUSTIC_SNARE = NoteFromCode(38)
var  HAND_CLAP = NoteFromCode(39)
var  ELECTRIC_SNARE = NoteFromCode(40)
var  LOW_FLOOR_TOM = NoteFromCode(41)
var  CLOSED_HIGH_HAT = NoteFromCode(42)
var  HIGH_FLOOR_TOM = NoteFromCode(43)
var  PEDAL_HIGH_HAT = NoteFromCode(44)
var  LOW_TOM = NoteFromCode(45)
var  OPEN_HIGH_HAT = NoteFromCode(46)
var  LOW_MID_TOM = NoteFromCode(47)
var  HIGH_MID_TOM = NoteFromCode(48)
var  CRASH_CYMBAL_1 = NoteFromCode(49)
var  HIGH_TOM = NoteFromCode(50)
var  CHINESE_CYMBAL = NoteFromCode(52)
var  RIDE_BELL = NoteFromCode(53)
var  TAMBOURINE = NoteFromCode(54)
var  SPLASH_CYMBAL = NoteFromCode(55)
var  COWBELL = NoteFromCode(56)
var  CRASH_CYMBAL_2 = NoteFromCode(57)
var  VIBRASLAP = NoteFromCode(58)
var  RIDE_CYMBAL_2 = NoteFromCode(59)
var  HIGH_BONGO = NoteFromCode(60)
var  LOW_BONGO = NoteFromCode(61)
var  MUTE_HIGH_CONGA = NoteFromCode(62)
var  OPEN_HIGH_CONGA = NoteFromCode(63)
var  LOW_CONGA = NoteFromCode(64)
var  HIGH_TIMBALE = NoteFromCode(65)
var  LOW_TIMBALE = NoteFromCode(66)
var  LOW_AGOGO = NoteFromCode(68)
var  CABASA = NoteFromCode(69)
var  MARACAS = NoteFromCode(70)
var  SHORT_WHISTLE = NoteFromCode(71)
var  LONG_WHISTLE = NoteFromCode(72)
var  SHORT_GUIRO = NoteFromCode(73)
var  LONG_GUIRO = NoteFromCode(74)
var  CLAVES = NoteFromCode(75)
var  HIGH_WOOD_BLOCK = NoteFromCode(76)
var  LOW_WOOD_BLOCK = NoteFromCode(77)
var  MUTE_CUICA = NoteFromCode(78)
var  OPEN_CUICA = NoteFromCode(79)
var  MUTE_TRIANGLE = NoteFromCode(80)
var  OPEN_TRIANGLE = NoteFromCode(81)


