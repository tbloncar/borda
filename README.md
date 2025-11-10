# Borda

A simple implementation of the [Borda count](https://en.wikipedia.org/wiki/Borda_count) ranked voting system.

## Usage

This code creates a new Borda contest with 3 candidates and ranks them based on the provided ballots.

```go
contest, err := NewBordaContest(3)
if err != nil {
    panic(err)
}
ca, cb, cc := Candidate{"A"}, Candidate{"B"}, Candidate{"C"}
ballots := [][]Candidate{
    {ca, cb, cc},
    {cb, cc, ca},
    {cc, ca, cb},
    {cc, ca, cb},
}
results, err := contest.Rank(ballots)
if err != nil {
    panic(err)
}
fmt.Println(results)
```

### Options

Pass options to the `NewBordaContest` function to configure the contest.

```go
contest, err := NewBordaContest(
  3,
  WithRequireFullBallot(false),
  WithRankScores([]int{4, 3, 2}),
)
```
