package main


type LadspaEffect struct {
	ffmpeg_cli string
	fluidsynth_code string
}

type Effect_ struct {
	preset string	
}

func Effect(v string) Effect_{
	return Effect_{preset: v}
}


func (s *Song_) define_effect(
   preset string, elist ...interface{}) *Song_ {
   
   //TODO
   return s
} 

