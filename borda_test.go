package borda

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBordaRank_FullBallotRequired(t *testing.T) {
	ballots := [][]Candidate{
		{{Id: "A"}, {Id: "B"}, {Id: "C"}},
		{{Id: "B"}, {Id: "C"}, {Id: "A"}},
		{{Id: "C"}, {Id: "A"}, {Id: "B"}},
		{{Id: "C"}, {Id: "A"}, {Id: "B"}},
	}
	contest, err := NewBordaContest(3, WithRequireFullBallot(true))
	assert.NoError(t, err)
	results, err := contest.Rank(ballots)
	assert.NoError(t, err)
	assert.Equal(t, results, []Result{
		{Candidate: Candidate{Id: "C"}, Score: 9},
		{Candidate: Candidate{Id: "A"}, Score: 8},
		{Candidate: Candidate{Id: "B"}, Score: 7},
	})
}

func TestBordaRank_FullBallotNotRequired(t *testing.T) {
	ballots := [][]Candidate{
		{{Id: "A"}, {Id: "B"}, {Id: "C"}},
		{{Id: "B"}, {Id: "C"}, {Id: "A"}},
		{{Id: "C"}, {Id: "A"}},
		{{Id: "C"}, {Id: "A"}},
	}
	contest, err := NewBordaContest(3, WithRequireFullBallot(false))
	assert.NoError(t, err)
	results, err := contest.Rank(ballots)
	assert.NoError(t, err)
	assert.Equal(t, results, []Result{
		{Candidate: Candidate{Id: "C"}, Score: 9},
		{Candidate: Candidate{Id: "A"}, Score: 8},
		{Candidate: Candidate{Id: "B"}, Score: 5},
	})
}

func TestBordaRank_IncompleteBallot(t *testing.T) {
	ballots := [][]Candidate{
		{{Id: "A"}, {Id: "B"}, {Id: "C"}},
		{{Id: "B"}, {Id: "C"}},
		{{Id: "C"}, {Id: "A"}, {Id: "B"}},
	}
	contest, err := NewBordaContest(3, WithRequireFullBallot(true))
	assert.NoError(t, err)
	_, err = contest.Rank(ballots)
	assert.ErrorContains(t, err, "detected incomplete ballot")
}
