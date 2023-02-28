package text

import (
	_ "embed"
	"strings"

	"github.com/samber/lo"
)

//go:embed heart.txt
var heartText string

func HeartLines() []string {
	return strings.Split(heartText, "\n")
}

// ShouldBeEroded decide if the point should be "eroded" in the next cycle
// by checking if any of it adjacent points (up, down, left, and right)
// is a blank space.
func ShouldBeEroded(lines []string, x int, y int) bool {
	if x == 0 || y == 0 || y == len(lines)-1 {
		return true
	}

	checkBound := func(line string, x int) bool {
		return x >= len(line)-1
	}
	checkCharacter := func(line string, x int) bool {
		return line[x] == ' '
	}
	checkLeftRight := func(line string, x int) bool {
		return checkCharacter(line, x+1) ||
			checkCharacter(line, x-1)
	}
	lineUp := lines[y+1]
	line := lines[y]
	lineDown := lines[y-1]

	final := checkBound(line, x) || checkLeftRight(line, x) ||
		checkBound(lineUp, x) || checkCharacter(lineUp, x) ||
		checkBound(lineDown, x) || checkCharacter(lineDown, x)

	return final
}

// Erode turns the outermost characters of an ASCII figure into blank spaces.
func Erode(lines []string) []string {
	newLines := make([]string, 0, len(lines))
	for y := 0; y < len(lines); y++ {
		line := lines[y]
		newLine := ""
		for x := 0; x < len(line); x++ {
			if ShouldBeEroded(lines, x, y) {
				newLine += " "
			} else {
				newLine += "*"
			}
		}
		newLines = append(newLines, newLine)
	}
	return newLines
}

func AllBlank(lines []string) bool {
	for _, line := range lines {
		if strings.TrimSpace(line) != "" {
			return false
		}
	}
	return true
}

func HeartStates() [][]string {
	firstState := HeartLines()
	states := [][]string{firstState}

	for {
		lastState := states[len(states)-1]
		if AllBlank(lastState) {
			break
		}
		states = append(states, Erode(lastState))
	}

	return lo.Reverse(states)
}
