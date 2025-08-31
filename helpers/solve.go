package helpers

import "fmt"

var (
	Ta = []string{}

	Visited = make(map[string]bool)
	Exists  = make(map[string]bool)

	AllPaths = [][]string{}
	Levels   = make(map[string]int)
)

func Levelchecker() {
	visited := make(map[string]bool)

	// BFS queue
	queue := []string{Data.Start}
	Levels[Data.Start] = 0
	visited[Data.Start] = true

	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]

		for _, neighbor := range Data.Links[current] {
			if !visited[neighbor] {
				visited[neighbor] = true
				Levels[neighbor] = Levels[current] + 1
				queue = append(queue, neighbor)
			}
		}
	}
}

func SV(end string, Ta []string) {
	if _, ok := Data.Links[end]; !ok {
		return
	}

	for _, v := range Data.Links[end] {

		if Exists[Ta[len(Ta)-1]] {
			return
		}
		if Exists[v] {
			continue
		}
		if v == Data.End {

			Ta = append(Ta, v)

			for j := 1; j < len(Ta)-1; j++ {
				Exists[Ta[j]] = true
			}

			AllPaths = append(AllPaths, Ta)

			return

		}

		if Visited[v] || Levels[end]>Levels[v]{
			continue
		}

		// Chana<-v
		Visited[v] = true
		newTa := make([]string, len(Ta))
		copy(newTa, Ta)
		newTa = append(newTa, v)
		SV(v, newTa)
		Visited[v] = false
	}
}

func FindPaths() {
	Levelchecker()
	Ta = []string{Data.Start}

	Visited[Data.Start] = true

	SV(Data.Start, Ta)
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
		if i == 1 {
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
