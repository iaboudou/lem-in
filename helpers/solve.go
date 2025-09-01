package bfs

import (
	"fmt"

	"lemin/config"
)

var (
	Ta = []string{}

	Visited = make(map[string]bool)
	Exists  = make(map[string]bool)

	AllPaths = [][]string{}
	Levels   = make(map[string]int)
)

func BFS() {
	if len(config.Data.Links) == 0 {
		return
	}
	q := [][]string{{config.Data.Start}}

	visited := map[string]int{}
	for len(q) > 0 {
		path := q[0]
		last := path[len(path)-1]
		q = q[1:]
		for _, n := range config.Data.Links[last] {
			if visited[n] > 50 {
				continue
			}
			if n == config.Data.Start {
				continue
			}
			if Exists[n] {
				continue
			}
			temp := append([]string{}, path...)
			temp = append(temp, n)
			//	fmt.Println("***", temp)

			if n == config.Data.End {
				AllPaths = append(AllPaths, temp)
			} else {
				q = append(q, temp)
				visited[n]++
			}
		}
	}

	/*for i := 0; i < len(AllPaths); i++ {
	loop:
		for j := 1; j < len(AllPaths[i])-1; j++ {
			for k := 1; k != i && k < len(AllPaths); k++ {
				if len(AllPaths[i]) <= len(AllPaths[k]) && AllPaths[i][j] == AllPaths[k][j] {
					AllPaths = append(AllPaths[:k], AllPaths[k+1:]...)
					i--
					break loop
				}
			}
		}
	}*/
}

func Delet_Duplicated_Pather() {
	alli := [][]string{}
	allf := [][]string{}

	for k, path := range AllPaths {
		alli = append(alli, path)
		Exists = make(map[string]bool)

		for i := 1; i < len(path)-1; i++ {
			Exists[path[i]] = true
		}
		for j, v := range AllPaths {
			if j != k {

				for i := 0; i < len(v); i++ {

					if Exists[v[i]] {
						break
					}
					Ta = append(Ta, v[i])
				}
				if len(Ta) == len(v) {
					alli = append(alli, Ta)

					for i := 1; i < len(v)-1; i++ {
						Exists[v[i]] = true
					}

				}

			}
			Ta = []string{}
		}
		if len(allf) < len(alli) {
			allf = alli
		}
		alli = [][]string{}
	}
	AllPaths = allf
}

func FindPaths() {
	BFS()
	// fmt.Println("----",Data.Links[Data.Start])
	// fmt.Println("****************************")
	// fmt.Println(Data.Links[Data.End])

	Delet_Duplicated_Pather()
	for _, v := range AllPaths {
		fmt.Println(v)
	}
}

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

		if fromIdNmla == 0 || Path[i+1].Nmla != 0 && to != config.Data.End {
			continue
		}

		if to == config.Data.End {
			finished++
		}

		Path[i].Nmla = 0
		Path[i+1].Nmla = fromIdNmla

		fmt.Printf("L%d-%s ", fromIdNmla, to)

		if i == 1 && config.Data.Nmber_Ants-IdNmla < len(Path)-2 {
			if indexPath != 0 && (len(NewPaths[indexPath])-len(NewPaths[indexPath-1]) > 1) {
				return
			}
		}
	}

	if IdNmla <= config.Data.Nmber_Ants && Path[1].Nmla == 0 {
		if len(Path) == 2 {
			fmt.Printf("L%d-%s ", IdNmla, Path[1].Name)
			IdNmla++
			finished++
			return
		}
		Path[1].Nmla = IdNmla
		if Path[1].Name == config.Data.End {
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

	for finished < config.Data.Nmber_Ants {
		for i := 0; i < len(NewPaths); i++ {
			PirntNmla(NewPaths[i], i)
		}
		fmt.Println()
	}
}
