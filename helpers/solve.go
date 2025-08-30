package helpers

import (
	"fmt"
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

		if last == Data.Start {
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
	IdNmla   int = 1
	finished int = 0
)

func PirntNmla(Path []*Room) {
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
		// fmt.Print("\n",Path[i].Name,"\n")
		Path[i+1].Nmla = fromIdNmla
		
		fmt.Printf("L%d-%s ", fromIdNmla, to)
		if i == 1{
			// fmt.Println("is 1")
		}
	}

	if len(Path) == 2 && IdNmla <= Data.Nmber_Ants && Path[1].Nmla == 0 {
		// fmt.Println("here")
		fmt.Printf("L%d-%s ", IdNmla, Path[1].Name)
		IdNmla++
		finished++
		return
	}

	// fmt.Print("\nkahsha tkon 0 -> ",Path[1].Nmla, " \n")
	if IdNmla <= Data.Nmber_Ants && Path[1].Nmla == 0 {
		Path[1].Nmla = IdNmla
		if Path[1].Name == Data.End {
			finished++
		}
		fmt.Printf("L%d-%s ", IdNmla, Path[1].Name)
		IdNmla++
		// fmt.Println("dkhel ")
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
