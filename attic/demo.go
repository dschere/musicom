package main


func main() {
  Song("stuff").
    timesig("4/4").
    key("C").
    beat(120).
    
    define_effect("dreamy").

    define_track("g1", ACOUSTIC_STEEL_GUITAR).
    define_track("p", DRUMS).
 
    measure("m1").
       track("g1", LegatoOn, Effect("dreamy"), Dq, N2_0, De,N5_1,N4_2,N3_2,N5_1,N4_2,N3_2).
       track("p" , Dw, ACOUSTIC_SNARE).
    measure().
       track("g1", Effect(""), Dq, N2_0, De,N5_0,N4_2,N3_2,N5_0,N4_2,N3_2).
       track("p" , Dw, ACOUSTIC_SNARE).
    repeat("m1",1).
  play() 
}
