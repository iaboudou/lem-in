package helpers

import (
	"fmt"
	"time"
)

func FindPaths() [][]string {
	AllPaths := [][]string{}
	if len(Data.Links) == 0 {
		return nil
	}

	contains := func(path []string, node string) bool {
		for _, ele := range path {
			if ele == node {
				return true
			}
		}
		return false
	}

	q := [][]string{{Data.Start}}

	for len(q) > 0 {
		path := q[0]
		lastElement := path[len(path)-1]
		q = q[1:]

		for _, n := range Data.Links[lastElement] {
			if contains(path, n) {
				continue
			}

			temp := append([]string{}, path...)
			temp = append(temp, n)

			if n == Data.End {
				AllPaths = append(AllPaths, temp)
			} else {
				q = append(q, temp)
			}
		}
	}
	return AllPaths
}

func Solve(paths [][]string) {
	if len(paths) == 0 {
		return
	}

	type Ant struct {
		ID       int
		Path     []string
		Position int
	}
	antID := 0
	ants := make([]*Ant, 0, Data.Nmber_Ants)

	room3amra := make(map[string]bool) // we put every buzy room here to true
	round := 0                         // it count rounds in eatch iteration

	for {
		printmoves := []string{} // this is will use for appending output string
		round++
		newAnts := []*Ant{} // will append ants moves at level

		for p, path := range paths {
			for room := range path {
				// move one ant
				newAnt := &Ant{ID: antID + 1, Position: room, Path: paths[p]}
				ants = append(ants, newAnt)
				printmoves = append(printmoves, fmt.Sprintf("L%d-%s", antID, path[room]))
				room3amra[path[room]] = true
				fmt.Println(printmoves)
				time.Sleep(1 * time.Second)
			}
		}
	}
}
