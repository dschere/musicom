package main

import (
//    "fmt"
    "log"
)


type Measure struct {
	id   int 
	cmds map[string][]interface{}
	repeat bool
    start_measure_id int
	count int
}

func (s *Song_) measure(name ...string) *Song_ {
	m := Measure{}
	m.cmds = make(map[string][]interface{})
	m.id = len(s.mlist)
	m.repeat = false
	
	s.current_measure = &m
	s.mlist = append(s.mlist, m)
	
//    fmt.Printf("measure %d %p\n", len(s.mlist), m)
    
	if len(name) == 1 {
		s.mlookup[name[0]] = m.id
	}
	
	return s
}

func (s *Song_) repeat(name string, count int) *Song_ {
	m := Measure{}
    m.repeat = true
    m.count = count
    
    id, found := s.mlookup[name]
    if found == false {
        log.Fatal("repeat(%s, %d) -> no matching measure for %s", 
            name, count, name)
    }

    m.start_measure_id = id
	s.mlist = append(s.mlist, m)

//    fmt.Printf("repeat %d %p\n", len(s.mlist), m)

    return s
}

func (s *Song_) track(trackId string, clist ...interface{}) *Song_ {
    _, found := s.trackTable[trackId]
    if found == false {
        log.Fatal("track(%s, %p) -> unknown track id %s",
            trackId, clist, trackId )
    }
	s.current_measure.cmds[trackId] = clist	
	return s
}
