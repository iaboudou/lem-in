package helpers

import (
	"fmt"
	"strings"
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

type Ant struct {
	ID       int
	Path     []string
	Position int
}
type Room struct {
	Name    string
	Channel chan *Ant
}

var ChMoves = make(chan string, 1000) // will take the string into print

func Solve(paths [][]string) {
	if len(paths) == 0 {
		return
	}

	// here we build channel for each path
	channs := [][]*Room{}
	for _, path := range paths {
		channs = append(channs, buildchannel(path))
	}

	antID := 1
	for i := 0; i < len(channs) && antID <= Data.Nmber_Ants; i++ {
		rooms := channs[i]

		ant := &Ant{ID: antID, Path: paths[i], Position: 0}
		rooms[0].Channel <- ant

		go moveAnt(ant, rooms)
		antID++

		time.Sleep(50 * time.Millisecond)
	}

	done := 0
	roundMoves := []string{}

	for done < Data.Nmber_Ants {
		for len(ChMoves) > 0 {
			move := <-ChMoves
			roundMoves = append(roundMoves, move)

			for _, path := range paths {
				if strings.Contains(move, path[len(path)-1]) {
					done++
					break
				}
			}
		}

		if len(roundMoves) > 0 {
			fmt.Println(strings.Join(roundMoves, " "))
			roundMoves = []string{}
		}

		time.Sleep(1 * time.Second)
	}
}

// will build a channel for each room in path that has size= 1
func buildchannel(path []string) []*Room {
	rooms := make([]*Room, len(path))

	for i, name := range path {
		size := 1
		if i == 0 || i == len(path)-1 {
			size = Data.Nmber_Ants
		}
		rooms[i] = &Room{Name: name, Channel: make(chan *Ant, size)}
	}
	return rooms
}

// just move ants don't be afraid, the room will take just one ant ////////////////////////
func moveAnt(a *Ant, rooms []*Room) {
	for i := 1; i < len(rooms); {
		rooms[i].Channel <- a
		ChMoves <- fmt.Sprintf("L%d-%s", a.ID, rooms[i].Name)

		if i > 0 {
			<-rooms[i-1].Channel
		}

		a.Position = i
		i++
		time.Sleep(30 * time.Millisecond)
	}
}
