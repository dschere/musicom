package main


type Dynamic struct {
    val int
}

var PPP = Dynamic{ val: 16 } // piano-pianissimo
var PP  = Dynamic{ val: 32 } // pianissimo
var P   = Dynamic{ val: 48 } // piano
var MP  = Dynamic{ val: 64 } // mezzo-piano
var MF  = Dynamic{ val: 80 } // mezzo-forte
var F   = Dynamic{ val: 96 } // forte
var FF  = Dynamic{ val: 112} // fortissimo
var FFF = Dynamic{ val: 127} // forte-fortissimo

