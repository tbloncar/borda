package borda

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBordaContest_Rank(t *testing.T) {
	ca, cb, cc := Candidate{"A"}, Candidate{"B"}, Candidate{"C"}

	tests := []struct {
		name            string
		ballots         [][]Candidate
		requireFull     bool
		rankScores      []int
		expectedResults []Result
		expectErr       string
	}{
		{
			name: "FullBallotRequired",
			ballots: [][]Candidate{
				{ca, cb, cc},
				{cb, cc, ca},
				{cc, ca, cb},
				{cc, ca, cb},
			},
			requireFull: true,
			rankScores:  nil,
			expectedResults: []Result{
				{Candidate: cc, Score: 9},
				{Candidate: ca, Score: 8},
				{Candidate: cb, Score: 7},
			},
			expectErr: "",
		},
		{
			name: "FullBallotNotRequired",
			ballots: [][]Candidate{
				{ca, cb, cc},
				{cb, cc, ca},
				{cc, ca},
				{cc, ca},
			},
			rankScores:  nil,
			requireFull: false,
			expectedResults: []Result{
				{Candidate: cc, Score: 9},
				{Candidate: ca, Score: 8},
				{Candidate: cb, Score: 5},
			},
			expectErr: "",
		},
		{
			name: "WithRankScores",
			ballots: [][]Candidate{
				{ca, cb, cc},
				{cb, cc, ca},
				{cc, ca, cb},
				{cc, ca, cb},
			},
			requireFull: true,
			rankScores:  []int{4, 3, 2},
			expectedResults: []Result{
				{Candidate: cc, Score: 13},
				{Candidate: ca, Score: 12},
				{Candidate: cb, Score: 11},
			},
			expectErr: "",
		},
		{
			name: "IncompleteBallotError",
			ballots: [][]Candidate{
				{ca, cb, cc},
				{cb, cc},
				{cc, ca, cb},
			},
			rankScores:      nil,
			requireFull:     true,
			expectedResults: nil,
			expectErr:       "detected incomplete ballot",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			options := []Option{WithRequireFullBallot(tt.requireFull)}
			if tt.rankScores != nil {
				options = append(options, WithRankScores(tt.rankScores))
			}
			contest, err := NewBordaContest(3, options...)
			assert.NoError(t, err)
			results, err := contest.Rank(tt.ballots)
			if tt.expectErr != "" {
				assert.ErrorContains(t, err, tt.expectErr)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.expectedResults, results)
			}
		})
	}
}
