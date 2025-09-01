package helpers

import "fmt"

var (
	Ta = []string{}

	Visited = make(map[string]bool)
	Exists  = make(map[string]bool)

	AllPaths = [][]string{}
	Levels   = make(map[string]int)
)

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

	Ta = []string{Data.Start}


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
