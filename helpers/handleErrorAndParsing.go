package helpers

import (
	"errors"
	"os"
	"strconv"
	"strings"
)

const (
	PathFiles = "./tests/"
)

type Romms struct {
	X, Y int
}

type Lemin struct {
	Nmber_Ants int
	Start      string
	End        string
	Roms       map[string]Romms
	Links      map[string][]string
}

var Data = Lemin{
	Roms:  make(map[string]Romms),
	Links: make(map[string][]string),
}

// function to check if the room is valid or no 
func IsValidRoom(line *string) bool {
	if (*line)[0] == 'L' {
		return false
	}
	Chunks := strings.Split(*line, " ")

	if len(Chunks) != 3 {
		return false
	}
	if len(Chunks[0]) > 50 {
		return false
	}

	if _, err := strconv.Atoi(Chunks[1]); err != nil {
		return false
	}
	if _, err := strconv.Atoi(Chunks[2]); err != nil {
		return false
	}
	if _, Exists := Data.Roms[Chunks[0]]; Exists {
		return false
	}

	return true
}

// function to extract the name of the room and the position from the currnt line
func ExtractDataRoom(line *string) (string, int, int) {
	var NameRoom string
	var x, y int
	Arr := strings.Split(*line, " ")
	NameRoom = Arr[0]
	x, _ = strconv.Atoi(Arr[1])
	y, _ = strconv.Atoi(Arr[2])
	return NameRoom, x, y
}

// function to check if the link between rooms if is valid or no 
func IsValidLink(line *string) bool {
	Chunks := strings.Split(*line, "-")
	if len(Chunks) != 2 || len(Chunks[0]) > 50 || len(Chunks[1]) > 50 {
		return false
	}
	if Chunks[0] == Chunks[1] {
		return false
	}

	_, Room1Exists := Data.Roms[Chunks[0]]
	_, Room2Exists := Data.Roms[Chunks[1]]
	if !Room1Exists || !Room2Exists {
		return false
	}

	return true
}

// function to extract the name of the rooms witch they are connected 
func ExtractDataLink(line *string) (string, string) {
	var NameRoom1, NameRoom2 string
	Arr := strings.Split(*line, "-")
	NameRoom1 = Arr[0]
	NameRoom2 = Arr[1]

	return NameRoom1, NameRoom2
}

// this function handles all the logic of pars and return the errors
func HandleError(file string) error {
	ArrBytes, err := os.ReadFile(PathFiles + file)
	if err != nil {
		return errors.New("ERROR: " + "Invalid path of file or error in Reading file")
	}

	if len(ArrBytes) == 0 {
		return errors.New("ERROR: " + "You have an empty file")
	}

	StrFile := string(ArrBytes)
	ArrFile := strings.Split(StrFile, "\n")

	// Check the ants number
	NumberAnts, err := strconv.Atoi(ArrFile[0])
	if err != nil || NumberAnts <= 0 {
		return errors.New("ERROR: " + "Invalid number of ants")
	}

	// Insert number of ants
	Data.Nmber_Ants = NumberAnts
	ArrFile = ArrFile[1:]

	FoundStart := 0  // Found Start
	FoundEnd := 0    // Found End
	InRooms := true  // If you are in rooms
	InLinks := false // If you are in links

	// Split Rooms and Links + Check New Lines
	for IndexOfLine, line := range ArrFile {
		// Check new lines
		if line == "" {
			return errors.New("ERROR: " + "You add an '\\n' in the file")
		}

		IsValidRoomV := IsValidRoom(&line)
		IsValidLinkV := IsValidLink(&line)

		// Skip commants and ##start + ##end
		if line == "##start" && IndexOfLine+1 < len(ArrFile) && IsValidRoom(&ArrFile[IndexOfLine+1]) {
			Data.Start, _, _ = ExtractDataRoom(&ArrFile[IndexOfLine+1])
			FoundStart++
			continue
		}
		if line == "##end" && IndexOfLine+1 < len(ArrFile) && IsValidRoom(&ArrFile[IndexOfLine+1]) {
			Data.End, _, _ = ExtractDataRoom(&ArrFile[IndexOfLine+1])
			FoundEnd++
			continue
		}
		if line[0] == '#' && !(line == "##start" || line == "##end") {
			continue
		}

		// Check error start and end
		if line == "##start" && IndexOfLine+1 < len(ArrFile) && !IsValidRoom(&ArrFile[IndexOfLine+1]) ||
			line == "##end" && IndexOfLine+1 < len(ArrFile) && !IsValidRoom(&ArrFile[IndexOfLine+1]) ||
			line == "##start" && FoundStart == 1 ||
			line == "##end" && FoundStart == 1 ||
			line == "##start" && IndexOfLine+1 >= len(ArrFile) ||
			line == "##end" && IndexOfLine+1 >= len(ArrFile) {
			return errors.New("ERROR: " + "Problem In your ##start or ##end")
		}

		// Check the lines if links or rooms
		if !IsValidRoomV && !IsValidLinkV {
			return errors.New("ERROR: " + "You have problem in this line : " + line)
		}
		if IsValidRoomV && InRooms {
			NameRoom, x, y := ExtractDataRoom(&line)
			Data.Roms[NameRoom] = Romms{X: x, Y: y}
			continue
		} else if IsValidLinkV && InRooms {
			InRooms = false
			InLinks = true

		}
		if IsValidLinkV && InLinks {

			NameRoom1, NameRoom2 := ExtractDataLink(&line)
			if !contains(Data.Links[NameRoom2], NameRoom1) {
				Data.Links[NameRoom1] = append(Data.Links[NameRoom1], NameRoom2)
			}
			if !contains(Data.Links[NameRoom2], NameRoom1) {
				Data.Links[NameRoom2] = append(Data.Links[NameRoom2], NameRoom1)
			}

			continue
		} else if IsValidRoomV && InLinks {
			return errors.New("ERROR: " + "You place this room after link(s) : " + line)
		}
	}

	// Check start and end
	if FoundStart < 1 {
		return errors.New("ERROR: " + "You don't have ##start in your file")
	} else if FoundStart > 1 {
		return errors.New("ERROR: " + "You have more than one ##start in your file")
	}
	if FoundEnd < 1 {
		return errors.New("ERROR: " + "You don't have ##end in your file")
	} else if FoundStart > 1 {
		return errors.New("ERROR: " + "You have more than one ##end in your file")
	}

	return nil
}

// contains we use it for check if the target exists in the array we give it and returns bool
func contains(s []string, target string) bool {
	for _, v := range s {
		if v == target {
			return true
		}
	}
	return false
}
