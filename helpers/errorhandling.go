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

func IsValidRoom(line string) bool {
	if line[0] == 'L' { return false }
	Chunks := strings.Split(line, " ")
	if len(Chunks) != 3 { return false }
	if len(Chunks[0]) > 50 { return false }
	if _, err := strconv.Atoi(Chunks[1]); err != nil { return false }
	if _, err := strconv.Atoi(Chunks[2]); err != nil { return false }
	return true
}

func IsValidLink(line string) bool {
	Chunks := strings.Split(line, "-")
	if len(Chunks) != 2 { return false }
	if len(Chunks[0]) > 50 || len(Chunks[1]) > 50 { return false }
	return true
}

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
	ArrFile = ArrFile[1:]

	FoundStart := false // Found Start 
	FoundEnd   := false // Found End 
	
	// Split Rooms and Links + Check New Lines
	for IndexOfLine, line := range ArrFile {

		// Check new lines
		if line == "" { return errors.New("ERROR: " + "You add an '\\n' in the file") }

		// Skip commants and ##start + ##end
		if line == "##start" { FoundStart = true ; continue }
		if line == "##end" {   FoundEnd   = true ; continue }
		if line[0] == '#' && !(line == "##start" || line == "##end") { continue }

		// Check the lines if links or rooms
		
	}


	// Check Rooms


	return nil
}
