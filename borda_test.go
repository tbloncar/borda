package borda_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/tbloncar/borda"
)

func TestBordaContest_Rank(t *testing.T) {
	ca, cb, cc := borda.Candidate{"A"}, borda.Candidate{"B"}, borda.Candidate{"C"}

	tests := []struct {
		name            string
		ballots         [][]borda.Candidate
		requireFull     bool
		rankScores      []int
		expectedResults []borda.Result
		expectErr       string
	}{
		{
			name: "FullBallotRequired",
			ballots: [][]borda.Candidate{
				{ca, cb, cc},
				{cb, cc, ca},
				{cc, ca, cb},
				{cc, ca, cb},
			},
			requireFull: true,
			rankScores:  nil,
			expectedResults: []borda.Result{
				{Candidate: cc, Score: 9},
				{Candidate: ca, Score: 8},
				{Candidate: cb, Score: 7},
			},
			expectErr: "",
		},
		{
			name: "FullBallotNotRequired",
			ballots: [][]borda.Candidate{
				{ca, cb, cc},
				{cb, cc, ca},
				{cc, ca},
				{cc, ca},
			},
			rankScores:  nil,
			requireFull: false,
			expectedResults: []borda.Result{
				{Candidate: cc, Score: 9},
				{Candidate: ca, Score: 8},
				{Candidate: cb, Score: 5},
			},
			expectErr: "",
		},
		{
			name: "WithRankScores",
			ballots: [][]borda.Candidate{
				{ca, cb, cc},
				{cb, cc, ca},
				{cc, ca, cb},
				{cc, ca, cb},
			},
			requireFull: true,
			rankScores:  []int{4, 3, 2},
			expectedResults: []borda.Result{
				{Candidate: cc, Score: 13},
				{Candidate: ca, Score: 12},
				{Candidate: cb, Score: 11},
			},
			expectErr: "",
		},
		{
			name: "IncompleteBallotError",
			ballots: [][]borda.Candidate{
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
			options := []borda.Option{borda.WithRequireFullBallot(tt.requireFull)}
			if tt.rankScores != nil {
				options = append(options, borda.WithRankScores(tt.rankScores))
			}
			contest, err := borda.NewBordaContest(3, options...)
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
