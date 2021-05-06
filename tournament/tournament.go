package tournament

import (
	"bufio"
	"fmt"
	"io"
	"sort"
	"strings"
)

type team struct {
	name   string
	played int
	win    int
	lost   int
	drawn  int
}

type teams map[string]*team

func (t *team) won() {
	t.played++
	t.win++
}

func (t *team) lose() {
	t.played++
	t.lost++
}

func (t *team) draw() {
	t.played++
	t.drawn++
}

func (t *team) points() int {
	return 3*t.win + 1*t.drawn
}

// getOrCreateTeam looks up a team by name, creating it if it does not exist yet
func (t teams) getOrCreateTeam(n string) *team {
	if t[n] == nil {
		t[n] = &team{name: n}
	}
	return t[n]
}

// sort returns a list of teams, sorted by points and then team name alphabetically
func (t teams) sort() []*team {
	var teams []*team
	for _, team := range t {
		teams = append(teams, team)
	}
	sort.Slice(teams, func(i, j int) bool {
		if teams[i].points() == teams[j].points() {
			return teams[i].name < teams[j].name
		}
		return teams[i].points() > teams[j].points()
	})
	
	return teams
}

func (t teams) fromResults(r io.Reader) error {
	s := bufio.NewScanner(r)
	for s.Scan() {
		l := s.Text()
		// bypass empty line and comments
		if l == "" || strings.HasPrefix(l, "#") {
			continue
		}

		tokens := strings.Split(l, ";")
		if len(tokens) != 3 {
			return fmt.Errorf("wrong field count -> line: %s (got %d fields)",
				l, len(tokens))
		}

		team1, team2, outcome := tokens[0], tokens[1], tokens[2]
		switch outcome {
		case "win":
			t.getOrCreateTeam(team1).won()
			t.getOrCreateTeam(team2).lose()
		case "loss":
			t.getOrCreateTeam(team1).lose()
			t.getOrCreateTeam(team2).won()
		case "draw":
			t.getOrCreateTeam(team1).draw()
			t.getOrCreateTeam(team2).draw()
		default:
			return fmt.Errorf("invalid outcome %s", outcome)
		}
	}

	return nil
}

// Tally reads a newline-separated list of match outcomes from r and writes
// a summary of the result to w
func Tally(r io.Reader, w io.Writer) error {
	t := make(teams)
	err := t.fromResults(r)
	if err != nil {
		return err
	}

	// the header is constant, using a format string make it easier to change the formatting
	_, err = fmt.Fprintf(w, "%-30s | %2s | %2s | %2s | %2s | %2s\n",
		"Team", "MP", "W", "D", "L", "P")
	if err != nil {
		return err
	}

	for _, team := range t.sort() {
		_, err := fmt.Fprintf(w, "%-30s | %2d | %2d | %2d | %2d | %2d\n",
			team.name, team.played, team.win, team.drawn, team.lost, team.points())
		if err != nil {
			return err
		}
	}

	return nil
}
