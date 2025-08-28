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
				}
			}
			visited[last] = true
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
			time.Sleep(1 * time.Second)
			if Path[i].Nmla == 0 {
				continue
			}
			from, to := Path[i].Name, Path[i+1].Name
			fromIdNmla := Path[i].Nmla
			Path[i].Nmla, Path[i+1].Nmla = 0, fromIdNmla
			
			if room3amra[to] {
				continue
			}
			if to == Data.End {
				finished++
			}else {
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
