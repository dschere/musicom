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

