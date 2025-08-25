package helpers

import (
	"fmt"
	"strings"
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
	ants := 0
	print := []string{}
	// moveAnts := []string{}
	for ants < Data.Nmber_Ants {
		for _, v := range paths {
			if ants >= Data.Nmber_Ants {
				return
			}
			ants++
			print = append(print, fmt.Sprintf("L%d-%s", ants, v[1]))
			fmt.Println(strings.Join(print, " "))
		}
	}
}
