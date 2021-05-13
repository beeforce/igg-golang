module github.com/nitipatl/igg-golang-workshop

go 1.16

// replace github.com/nitipatl/go-math => ../math

require (
	github.com/nitipatl/go-deck v0.0.0-00010101000000-000000000000
	// github.com/nitipatl/go-math v0.0.0-00010101000000-000000000000
)

replace github.com/nitipatl/go-deck => ./deck