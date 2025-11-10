package borda

import (
	"errors"
	"sort"
)

type Contest struct {
	NVotes            int
	RequireFullBallot bool
	RankScores        []int
}

type Option func(*Contest) error

func WithRequireFullBallot(requireFullBallot bool) Option {
	return func(c *Contest) (err error) {
		c.RequireFullBallot = requireFullBallot
		return
	}
}

func WithRankScores(rankScores []int) Option {
	return func(c *Contest) (err error) {
		if len(rankScores) != c.NVotes {
			return errors.New("a score must be provided for each rank")
		}
		c.RankScores = rankScores
		return
	}
}

type Candidate struct {
	Id string
}

type Result struct {
	Candidate Candidate
	Score     int
}

func NewBordaContest(nVotes int, options ...Option) (*Contest, error) {
	if nVotes <= 0 {
		return nil, errors.New("vote count must be greater than 0")
	}
	rankScores := make([]int, nVotes)
	for i := 0; i < nVotes; i++ {
		rankScores[i] = nVotes - i
	}
	c := &Contest{
		NVotes:            nVotes,
		RequireFullBallot: true,
		RankScores:        rankScores,
	}
	for _, option := range options {
		if err := option(c); err != nil {
			return nil, err
		}
	}
	return c, nil
}

func (c *Contest) Rank(
	ballots [][]Candidate,
) ([]Result, error) {
	if c.RequireFullBallot {
		for _, ballot := range ballots {
			if len(ballot) != c.NVotes {
				return nil, errors.New("detected incomplete ballot")
			}
		}
	}

	scores := make(map[Candidate]int)

	for _, ballot := range ballots {
		for i, candidate := range ballot {
			scores[candidate] += c.RankScores[i]
		}
	}

	results := make([]Result, 0, len(scores))
	for candidate, score := range scores {
		results = append(results, Result{Candidate: candidate, Score: score})
	}

	sort.Slice(results, func(i, j int) bool { return results[i].Score > results[j].Score })

	return results, nil
}
