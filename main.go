package main

/*
 */
func main() {
    
  Song("stuff").
    timesig("4/4").
    key("C").
    beat(120).

    define_track("g1", ACOUSTIC_STEEL_GUITAR, TrackOpts{ 
        legato: true,
        dynamic: []Dynamic{F,MP,MP,MP,MF,MP,MP},
    }).
    define_track("p", DRUMS).
 
    measure("m0").
       track("g1", Reverb(100.0), Dw, N6_3.b(BEND_RELEASE_WHOLE), Reverb(0.0)).
       track("p", Dw, R).
 
    measure("m1").
       track("g1", Dq, N2_0, De,N5_1,N4_2,N3_2,N5_1,N4_2,N3_2).
       track("p" , Dw, ACOUSTIC_SNARE).
    measure().
       track("g1", Dq, N2_0, De,N5_0,N4_2,N3_2,N5_0,N4_2,N3_2).
       track("p" , Dw, ACOUSTIC_SNARE).
    repeat("m1",1).
  play() 

}
