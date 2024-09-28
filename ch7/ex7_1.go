package main

import (
	"io"
	"os"
	"sort"
)

type Team struct {
	Name string
	Players []string
}

type League struct {
	Name string
	Teams map[string]Team
	Wins map[string]int
}

func (l *League) MatchResult(team1 string, score1 int, team2 string, score2 int) {
	if _, ok := l.Teams[team1]; !ok {
		return
	}
	if _, ok := l.Teams[team2]; !ok {
		return
	}
	if score1 > score2 {
		l.Wins[team1]++
	} else if score2 > score1 {
		l.Wins[team2]++
	}
}

func (l League) Ranking() []string {
	var res []string

	for k := range l.Wins {
		res = append(res, k)
	}

	sort.Slice(res, func(i, j int) bool {
		return l.Wins[res[i]] > l.Wins[res[j]]
	})

	return res
}

type Ranker interface {
	Ranking() []string
}

func RankPrinter(r Ranker, out io.Writer) {
	for _, team := range r.Ranking() {
		_, err := io.WriteString(out, team + "\n")
		if err != nil {
			return
		}
	}
}

func main() {
	league := League{
		Name: "NBA",
		Teams: map[string]Team{
			"LA": {
				Name: "LA",
				Players: []string{"James", "Davis"},
			},
			"Denver": {
				Name: "Denver",
				Players: []string{"Jokic", "Murray"},
			},
			"NY": {
				Name: "NY",
				Players: []string{"Hart", "Bronson"},
			},
		},
		Wins: map[string]int{
			"LA": 0,
			"Denver": 0,
			"NY": 0,
		},
	}

	league.MatchResult("LA", 100, "Denver", 110)
	league.MatchResult("LA", 100, "Denver", 120)
	league.MatchResult("NY", 100, "Denver", 110)
	league.MatchResult("NY", 110, "Denver", 110)
	league.MatchResult("LA", 120, "NY", 110)
	league.MatchResult("LA", 120, "NY", 110)

	//fmt.Println(league.Ranking())
	RankPrinter(league, os.Stdout)
}
