package musicom


type LadspaEffect struct {
	ffmpeg_cli string
	fluidsynth_code string
}

type Effects struct {
	preset string
	
}


func (s *Song_) define_effect(
   preset string, elist ...interface{}) *Song_ {
   
   //TODO
   return s
} 

func (s *Song_) Effect(preset string) *Song_ {
   
   //TODO 
   return s
} 
