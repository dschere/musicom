package player


/*
# example fluidsynth commands
ladspa_effect e1 /usr/lib/ladspa/delay.so
ladspa_link e1 Input Main:L
ladspa_link e1 Output Main:L
ladspa_set e1 Delay 0.5

ladspa_effect e2 /usr/lib/ladspa/delay.so
ladspa_link e2 Input Main:R
ladspa_link e2 Output Main:R
ladspa_set e2 Delay 1.5

 * 
 */


type BeatEvent struct {
	// port number -> list of fluidsynth commands to execute.
	fs_cmds map[int][]string
	// TODO ffplay commands
}

type Player struct {
	// beat associated with an event that triggers fluidsynth
	// The score is the timeline representated as a sparce
	// matrix. we walk from beat 0 till the end of the peice and
	// if there a BeatEvent for that beat we execute commands. 
	score map[int]BeatEvent
	
	
}
