package main

import (
	"fmt"

	"github.com/tbloncar/borda"
)

func main() {
	contest, err := borda.NewBordaContest(3)
	if err != nil {
		panic(err)
	}
	ca, cb, cc := borda.Candidate{Id: "A"}, borda.Candidate{Id: "B"}, borda.Candidate{Id: "C"}
	ballots := [][]borda.Candidate{
		{ca, cb, cc},
		{cb, cc, ca},
		{cc, ca, cb},
		{cc, ca, cb},
	}
	results, err := contest.Rank(ballots)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%+v\n", results)
}
