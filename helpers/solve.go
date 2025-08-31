package helpers

import (
	"fmt"
)

// this fucntion loop over the "Data.Links" and  return all the exists paths in a graph
func FindPaths() [][]string {
	AllPaths := [][]string{}
	if len(Data.Links) == 0 {
		return nil
	}
	q := [][]string{{Data.Start}}
	visited := map[string]bool{Data.Start: true}
	for len(q) > 0 {
		path := q[0]
		last := path[len(path)-1]
		q = q[1:]

		for _, n := range Data.Links[last] {
			if visited[n] {
				continue
			}
			temp := append([]string{}, path...)
			temp = append(temp, n)

			if n == Data.End {
				AllPaths = append(AllPaths, temp)
			} else {
				q = append(q, temp)
				visited[n] = true
			}
		}
	}
	return AllPaths
}

// define Room in A path
type Room struct {
	Name string
	Nmla int
}

var (
	IdNmla   int = 1
	finished int = 0
)

var NewPaths [][]*Room

// the fucntion move ants forward along the path, starting from the last to the first room
func PirntNmla(Path []*Room, indexPath int) {
	if len(Path) < 2 {
		return
	}

	for i := len(Path) - 2; i >= 1; i-- {
		to := Path[i+1].Name
		fromIdNmla := Path[i].Nmla

		if fromIdNmla == 0 || Path[i+1].Nmla != 0 && to != Data.End {
			continue
		}

		if to == Data.End {
			finished++
		}

		Path[i].Nmla = 0
		Path[i+1].Nmla = fromIdNmla

		fmt.Printf("L%d-%s ", fromIdNmla, to)

		if i == 1 && Data.Nmber_Ants-IdNmla < len(Path)-2 {
			if indexPath != 0 && (len(NewPaths[indexPath])-len(NewPaths[indexPath-1]) > 1) {
				return
			}
		}
	}

	if IdNmla <= Data.Nmber_Ants && Path[1].Nmla == 0 {
		if len(Path) == 2 {
			fmt.Printf("L%d-%s ", IdNmla, Path[1].Name)
			IdNmla++
			finished++
			return
		}
		Path[1].Nmla = IdNmla
		if Path[1].Name == Data.End {
			finished++
		}
		fmt.Printf("L%d-%s ", IdNmla, Path[1].Name)
		IdNmla++
	}
}

// this fucntion try to loop over the path it print all the moves did to reach the end, it take a matrix of string the initlialize a room in every case
func Solve(Paths [][]string) {
	NewPaths = make([][]*Room, len(Paths))
	for i, Path := range Paths {
		Pth := make([]*Room, len(Path))
		for j, roomName := range Path {
			Pth[j] = &Room{Name: roomName, Nmla: 0}
		}
		NewPaths[i] = Pth
	}

	for finished < Data.Nmber_Ants {
		for i := 0; i < len(NewPaths); i++ {
			PirntNmla(NewPaths[i], i)
		}
		fmt.Println()
	}
}
