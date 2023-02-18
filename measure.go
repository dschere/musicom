package musicom


type Measure struct {
	id   int 
	cmds map[string][]interface{}
	repeat bool
	count int
}

func (s *Song_) measure(name ...string) *Song_ {
	m := Measure{}
	m.cmds = make(map[string][]interface{})
	m.id = len(s.mlist)
	m.repeat = false
	
	s.current_measure = m
	
	s.mlist = append(s.mlist, m)
	
	if len(name) == 1 {
		s.mlookup[name[0]] = m.id
	}
	
	return s
}

func (s *Song_) repeat(name string, count int) *Song_ {
	m := Measure{}
    m.repeat = true
    m.count = count
	s.mlist = append(s.mlist, m)
    return s        	
}

func (s *Song_) track(trackId string, clist ...interface{}) *Song_ {
	s.current_measure.cmds[trackId] = clist	
	return s
}
