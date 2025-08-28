package helpers

import (
	"fmt"
	"time"
)

type visit struct {
	vI int
	vB bool
}

func FindPaths() [][]string {
	AllPaths := [][]string{}
	if len(Data.Links) == 0 {
		return nil
	}
	q := [][]string{{Data.Start}}

	visited := map[string]visit{Data.Start: {vI: 1, vB: false}}
	for len(q) > 0 {
		path := q[0]
		last := path[len(path)-1]
		q = q[1:]

		for _, n := range Data.Links[last] {
			if visited[n].vI > ((1 * Data.Nmber_Ants) / 100) {
				continue
			}
			if visited[n].vB {
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
		
		if last == Data.Start{
			visited[last] = visit{vI: visited[last].vI + 1, vB: true}
		}
		visited[last] = visit{vI: visited[last].vI + 1, vB: visited[last].vB}


	}

	for i := 0; i < len(AllPaths)-2; i++ {
		for j := 1; j < len(AllPaths[i])-1; j++ {
			if len(AllPaths[i]) <= len(AllPaths[i+1]) && AllPaths[i][j] == AllPaths[i+1][j] {
				AllPaths = append(AllPaths[:i+1], AllPaths[i+2:]...)
				i--
				break
			}
		}
		if i >= Data.Nmber_Ants {
			break
		}
	}
	return AllPaths
}

type Room struct {
	Name string
	Nmla int
}

var (
	room3amra     = make(map[string]bool)
	done      int = 0
	finished  int = 0
)

func PirntNmla(Path []*Room) {
	if len(Path) < 2 {
		return
	}
	for i := len(Path) - 2; i >= 1; i-- {
		time.Sleep(90 * time.Millisecond)
		if Path[i].Nmla == 0 {
			continue
		}
		from, to := Path[i].Name, Path[i+1].Name
		fromIdNmla := Path[i].Nmla

		if room3amra[to] && to != Data.End {
			continue
		}
		Path[i].Nmla = 0

		if to == Data.End {
			finished++
		} else {
			Path[i+1].Nmla = fromIdNmla
			room3amra[to] = true
		}

		if from != Data.End {
			room3amra[from] = false
		}
		fmt.Printf("L%d-%s ", fromIdNmla, to)
	}
	// put the next ants in the begining
	if done < Data.Nmber_Ants && Path[1].Nmla == 0 {
		if !room3amra[Path[1].Name] {
			done++
			Path[1].Nmla = done
			room3amra[Path[1].Name] = true
			fmt.Printf("L%d-%s ", done, Path[1].Name)
		}
	}
}

func Solve(Paths [][]string) {
	NewPaths := make([][]*Room, len(Paths))
	for i, Path := range Paths {
		Pth := make([]*Room, len(Path))
		for j, roomName := range Path {
			Pth[j] = &Room{Name: roomName, Nmla: 0}
		}
		NewPaths[i] = Pth
	}

	for finished < Data.Nmber_Ants {
		for i := 0; i < len(NewPaths); i++ {
			PirntNmla(NewPaths[i])
		}
		fmt.Println()
	}
}
