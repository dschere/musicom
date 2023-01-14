
adding a simple but important feature to ffplay. The ability to suspend
until SIGUSR2 is issues to the process.

The play is to use ffplay to play recorded audio in sync with midi messages
issued to fluidsynth. Timing is crucial. I want to preload ffplay and have it 
loaded with the correct confuration then issue a SIGUSR2 at the moment a noteon
event is issued to fluidsynth. 

Note:
   ffplay can play from a start -> end time and loop 
   meaning it can be synced to measure of music 

   ffplay can apply ladspa effects (echo,delay,flanger etc.)


