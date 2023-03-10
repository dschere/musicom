package main

import (
   "strings"
   "strconv"
   "log"
)

type TimeSignature struct {
	beatsPerBar int `default:4`
	beatUnit    int `default:4`
}

var T5_4 = TimeSignature{beatsPerBar: 5, beatUnit: 4}
var T4_4 = TimeSignature{beatsPerBar: 4, beatUnit: 4}
var T3_4 = TimeSignature{beatsPerBar: 3, beatUnit: 4}
var T6_8 = TimeSignature{beatsPerBar: 6, beatUnit: 8}
var T7_8 = TimeSignature{beatsPerBar: 7, beatUnit: 8}


func (ts *TimeSignature) parse (c string) {
	parts := strings.Split(c,"/")
	ts.beatUnit = 4
	
	if len(parts) == 1 {
		i, err := strconv.Atoi(c)
		if err != nil {
			log.Panic("Invalid time signature " + c)
		} else {
		    ts.beatsPerBar = i
	    }  
	} else {
		n, err := strconv.Atoi(parts[0])
		if err != nil {
			log.Panic("Invalid time signature " + c)
		} else {
		    d, err := strconv.Atoi(parts[1])
			if err != nil {
				log.Panic("Invalid time signature " + c)
			} else {
		        ts.beatsPerBar = n
		        ts.beatUnit = d		 
		    }			
	    }
	}
}

