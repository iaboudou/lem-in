package helpers

import "fmt"

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
				visited[n] = true
				q = append(q, temp)
			}
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
		if Path[i].Nmla == 0 {
			continue
		}
		from, to := Path[i].Name, Path[i+1].Name
		fromIdNmla := Path[i].Nmla

		Path[i].Nmla, Path[i+1].Nmla = 0, fromIdNmla

		// check if "to" is "room3amra"
		if to != Data.End && to != Data.Start {
			if room3amra[to] {
				continue
			}
		}

		if to != Data.End {
			room3amra[to] = true
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

	// for i := len(Path) - 2; i >= 1; i-- {
	// 	if Path[i].Nmla == 0 {
	// 		continue
	// 	}
	// 	willgoto := Path[i+1].Name
	// 	willgofrom := Path[i].Name
	// 	id := Path[i].Nmla

	// 	Path[i].Nmla = 0
	// 	Path[i+1].Nmla = id

	// 	if willgoto != Data.Start && willgoto != Data.End {
	// 		if room3amra[willgoto] {
	// 			continue
	// 		}
	// 	}

	// 	if willgoto != Data.Start && willgoto != Data.End {
	// 		room3amra[willgoto] = true
	// 	}
	// 	if willgofrom != Data.Start && willgofrom != Data.End {
	// 		room3amra[willgofrom] = false
	// 	}

	// 	if willgoto == Data.End {
	// 		finished++
	// 	}

	// 	fmt.Printf("L%d-%s ", id, willgoto)
	// }

	// if done < Data.Nmber_Ants && Path[1].Nmla == 0 {
	// 	if Path[1].Name == Data.Start || Path[1].Name == Data.End || !room3amra[Path[1].Name] {
	// 		done++
	// 		Path[1].Nmla = done
	// 		if Path[1].Name != Data.Start && Path[1].Name != Data.End {
	// 			room3amra[Path[1].Name] = true
	// 		}
	// 		fmt.Printf("L%d-%s ", done, Path[1].Name)
	// 	}
	// }
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

// type Room struct {
// 	Name string
// 	Nmla int
// }

// var IdNmla int = 0

// func PirntNmla(Path []*Room) {
// 	IdNmla++
// 	if IdNmla > Data.Nmber_Ants {
// 		return
// 	}
// 	for i := len(Path) - 2; i >= 1; i-- {
// 		if Path[i].Nmla == 0 {
// 			continue
// 		}

// 		temp := Path[i].Nmla
// 		Path[i].Nmla = 0
// 		Path[i-1].Nmla = temp
// 		if i != 1 {
// 			fmt.Println("here")
// 			fmt.Printf("L%d-%s ", IdNmla, Path[i].Name)
// 		}
// 	}
// 	Path[1].Nmla = IdNmla
// 	fmt.Printf("L%d-%s ", IdNmla, Path[1].Name)
// }

// func Solve(Paths [][]string) {
// 	// start := make([]int, Data.Nmber_Ants)
// 	// for i := 1; i <= Data.Nmber_Ants; i++ {
// 	// 	start[i-1] = i
// 	// }

// 	NewPaths := make([][]*Room, len(Paths))
// 	for i, Path := range Paths {
// 		Pth := make([]*Room, len(Path))
// 		for j, roomName := range Path {
// 			Pth[j] = &Room{Name: roomName, Nmla: 0}
// 		}
// 		NewPaths[i] = Pth
// 	}

// 	for IdNmla <= Data.Nmber_Ants {
// 		for i := 0; i < len(NewPaths); i++ {
// 			PirntNmla(NewPaths[i])
// 			if IdNmla > Data.Nmber_Ants {
// 				break
// 			}
// 		}
// 		fmt.Println()
// 		// for i := 0; i < len(NewPaths) ; i++ {
// 		// 	for j := 1 ; j < len(NewPaths[i]) ; j++ {
// 		// 		if NewPaths[i][j].Nmla != 0 {
// 		// 			fmt.Print(NewPaths[i][j].Name," -> ",NewPaths[i][j].Nmla," ")
// 		// 		}
// 		// 	}
// 		// 	fmt.Println()
// 		// }
// 	}
// }

// func Solve(paths [][]string) {
// 	if len(paths) == 0 {
// 		return
// 	}

// 	// here we build channel for each path
// 	channs := [][]*Room{}
// 	for _, path := range paths {
// 		channs = append(channs, buildchannel(path))
// 	}

// 	antID := 1
// 	for i := 0; i < len(channs) && antID <= Data.Nmber_Ants; i++ {
// 		rooms := channs[i]

// 		ant := &Ant{ID: antID, Path: paths[i], Position: 0}
// 		rooms[0].Channel <- ant

// 		go moveAnt(ant, rooms)
// 		antID++

// 		// time.Sleep(50 * time.Millisecond)
// 	}

// 	done := 0
// 	roundMoves := []string{}

// 	for done < Data.Nmber_Ants {
// 		for len(ChMoves) > 0 {
// 			move := <-ChMoves
// 			roundMoves = append(roundMoves, move)

// 			for _, path := range paths {
// 				if strings.Contains(move, path[len(path)-1]) {
// 					done++
// 					break

// 				}
// 			}
// 		}

// 		if len(roundMoves) > 0 {
// 			fmt.Println(strings.Join(roundMoves, " "))
// 			roundMoves = []string{}
// 		}

// 		// time.Sleep(1 * time.Second)
// 	}
// }

// // will build a channel for each room in path that has size= 1
// func buildchannel(path []string) []*Room {
// 	rooms := make([]*Room, len(path))

// 	for i, name := range path {
// 		size := 1
// 		if i == 0 || i == len(path)-1 {
// 			size = Data.Nmber_Ants
// 		}
// 		rooms[i] = &Room{Name: name, Channel: make(chan *Ant, size)}
// 	}
// 	return rooms
// }

// // just move ants don't be afraid, the room will take just one ant ////////////////////////
// func moveAnt(a *Ant, rooms []*Room) {
// 	for i := 1; i < len(rooms); {
// 		rooms[i].Channel <- a
// 		ChMoves <- fmt.Sprintf("L%d-%s", a.ID, rooms[i].Name)

// 		if i > 0 {
// 			<-rooms[i-1].Channel
// 		}

// 		a.Position = i
// 		i++
// 		// time.Sleep(30 * time.Millisecond)
// 	}
// }
