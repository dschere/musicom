package main


type TrackOpts struct {
	legato bool
	stacatto bool
	string1 string  `standard:"E3"  dropd:"D3" `
	string2 string  `standard:"A3"  dropd:"A3" `  
	string3 string  `standard:"D4"  dropd:"D4" `
	string4 string  `standard:"G4"  dropd:"G4" `
	string5 string  `standard:"B4"  dropd:"B4" `
	string6 string  `standard:"E5"  dropd:"E5" `
	tuning_name string
    
    // create a beat based on varying loudness of notes.
    dynamic []Dynamic
}

type Track struct {
	instrument int
    opt TrackOpts	
}



func (s *Song_) define_track(
  name string, instrument int, opts ...TrackOpts) *Song_ {

    t := Track{}
    t.instrument = instrument

    t.opt.legato = false
    t.opt.stacatto = false
    
    t.opt.string1 = "E3"
    t.opt.string2 = "A3"
    t.opt.string3 = "D4"
    t.opt.string4 = "G4"
    t.opt.string5 = "B4"
    t.opt.string6 = "E5"
    
    if len(opts) == 1 {
		
		if opts[0].tuning_name == "dropd" {
			opts[0].string1 = "D3"
		}
		
		if opts[0].tuning_name == "dadgad" {
			opts[0].string1 = "D3"
            opts[0].string6 = "D5"			
		}
		
        t.opt.legato = opts[0].legato
        t.opt.stacatto = opts[0].stacatto
        t.opt.dynamic = opts[0].dynamic
        if len(opts[0].string1) > 0 { t.opt.string1 = opts[0].string1}    
        if len(opts[0].string2) > 0 { t.opt.string2 = opts[0].string2}    
        if len(opts[0].string3) > 0 { t.opt.string3 = opts[0].string3}    
        if len(opts[0].string4) > 0 { t.opt.string4 = opts[0].string4}    
        if len(opts[0].string5) > 0 { t.opt.string5 = opts[0].string5}    
        if len(opts[0].string6) > 0 { t.opt.string6 = opts[0].string6}
	} 
	
	s.trackTable[name] = t
    
    return s
}

