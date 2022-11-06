package main

import (
	"flag"
	"fmt"
	"math"
)

// Handicap Differential
//
// (adjusted gross score minus course rating) times (113 divided by slope rating)
//
// (score - cr) * (113 / sr)

// (83-72.6) * (113/135)

const standardSlope = 113

// presidio (72.6/135), metro (73.5/129), harding(72.7/125), corica (73/128), peacock (69.9/125)

func main() {
	score := flag.Float64("s", 0.0, "adjusted gross score")
	courseRating := flag.Float64("cr", 72.6, "course rating")
	slopeRating := flag.Float64("sr", 135.0, "slope rating")
	flag.Parse()
	differential := (*score - *courseRating) * (standardSlope / *slopeRating)

	fmt.Printf("Score: %v\n", int(*score))
	fmt.Printf("Course Rating: %v\n", *courseRating)
	fmt.Printf("Slope Rating: %v\n", int(*slopeRating))
	// little hack or go quark to get rounding to the tenth place
	// https://stackoverflow.com/questions/52048218/round-all-decimal-points-in-golang
	fmt.Printf("Differential: %v\n", math.Round(differential*10)/10)
}
