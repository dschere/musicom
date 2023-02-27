module main

go 1.18

replace example.com/player => ./player

require example.com/player v0.0.0-00010101000000-000000000000

require github.com/go-cmd/cmd v1.4.1 // indirect
