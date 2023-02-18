package musicom

type Song_ struct {
	name string
}

func (s *Song_) Song(n string) *Song_{
	s.name = n
	return s
}


