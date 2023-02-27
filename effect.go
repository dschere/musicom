package main


// builtin effects
const BEND    = "bend"
const VIBRATO = "vibrato"
const SLIDE   = "slide"
const REVERB  = "reverb"
const CHORUS  = "chorus"
const PAN     = "pan" 


type Effect struct {
    name string
    
    f1   float32
    i1   int
    f2   float32
    i2   int
}




