package main

type Note_ struct {
	rest  bool
	mcode string // midi value
	fret  int
	gstr  int

    effects []Effect
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


