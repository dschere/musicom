package musicom

type Song_ struct {
	name  string
	ts    TimeSignature
	key_  string
	bpm   int  
	trackTable map[string]Track

}

func (s *Song_) Song(v string) *Song_{
	s.name = v
	s.ts = TimeSignature{}
	s.trackTable = make(map[string]Track)
	
	
	// set defaults
	s.ts.parse("4/4")
	s.key_ = "C"
	
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


