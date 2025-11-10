# Borda

A simple implementation of the [Borda count](https://en.wikipedia.org/wiki/Borda_count) ranked voting system.

## Usage

```go
contest, err := NewBordaContest(3)
if err != nil {
    panic(err)
}
ca, cb, cc := Candidate{Id: "A"}, Candidate{Id: "B"}, Candidate{Id: "C"}
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
