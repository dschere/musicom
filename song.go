package main


type Song_ struct {
	name  string
	ts    TimeSignature
	key_  string
	bpm   int  
	trackTable map[string]Track
    mlist []Measure
    current_measure* Measure
    mlookup map[string]int
}

func Song(v string) *Song_ {
	s := &Song_{}
	return s.init(v)
}

func (s *Song_) init(v string) *Song_{
	s.name = v
	s.ts = TimeSignature{}
	s.trackTable = make(map[string]Track)
	s.mlist = make([]Measure, 0)
	s.mlookup = make(map[string]int)
	
	
	// set defaults
	s.ts.parse("4/4")
	s.key_ = "C"
	s.bpm = 120
	
	return s
}

func (s *Song_) timesig(v string) *Song_{
	s.ts.parse(v)
	return s
}

func (s *Song_) beat(v int) *Song_ {
	s.bpm = v
	return s
}

func (s *Song_) key(v string) *Song_{
	s.key_ = v
	return s
}

func (s* Song_) play() {
    if s.current_measure != nil {
	    s.mlist = append(s.mlist, *s.current_measure)
        s.current_measure = nil
    }
    seq := Sequencer{}

    seq.compile(s)
    seq.pretty_print()
}




