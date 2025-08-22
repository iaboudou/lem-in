package helpers

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Romms struct {
	x, y int
}

type Lemin struct {
	Nmber_Ants int
	Start      string
	End        string
	Roms       map[string]Romms
	Links      map[string][]string
}

var Info = Lemin{
	Roms:  make(map[string]Romms),
	Links: make(map[string][]string),
}

func Parssing(textfile string) {
	file, err := os.ReadFile(textfile)
	if err != nil {
		fmt.Println("1 err reading file :", err)
		return
	}
	data := strings.Split(string(file), "\n")
	var Its_start bool
	var Its_end bool
	for j, v := range data {
		fields := strings.Fields(strings.TrimSpace(v))
		if strings.HasPrefix(v, "#") && !strings.HasPrefix(v, "##start") && !strings.HasPrefix(v, "##end") {
			continue
		}

		if j == 0 {
			Info.Nmber_Ants, _ = strconv.Atoi(fields[0])
		} else if len(fields) == 1 {
			if fields[0] == "##start" {
				Its_start = true
				continue
			} else if fields[0] == "##end" {
				Its_end = true
				continue
			} else {
				f := strings.Split(string(fields[0]), "-")

				Info.Links[f[0]] = append(Info.Links[f[0]], f[1])
			}
		} else if len(fields) == 3 {
			if Its_start {
				temprom := Append_Rooms(fields)
				Info.Start = fields[0]
				Info.Roms[fields[0]] = temprom
				Its_start = false
			} else if Its_end {
				temprom := Append_Rooms(fields)
				Info.End = fields[0]
				Info.Roms[fields[0]] = temprom
				Its_end = false
			} else {

				temprom := Append_Rooms(fields)
				Info.Roms[fields[0]] = temprom
			}
		}

	}
}

func Append_Rooms(fields []string) Romms {
	Xval, _ := strconv.Atoi(fields[1])
	Yval, _ := strconv.Atoi(fields[2])

	temprom := Romms{
		x: Xval,
		y: Yval,
	}
	return temprom
}
