package ozoncontest

import (
	"fmt"
)

const (
	Right rune = 'R'
	Left  rune = 'L'
	Up    rune = 'U'
	Down  rune = 'D'
	Home  rune = 'B'
	End   rune = 'E'
	Enter rune = 'N'
)

func insertAtIndex(str string, char rune, index int) string {
	if index < 0 || index > len(str) {
		return str
	}
	return str[:index] + string(char) + str[index:]
}

func terminal(str string) string {
	lengths := make([]int, 1, 50)
	result := ""
	currentIndex := 0
	currentLine := 0
	currentColumn := 0

	for _, char := range str {
		switch char {
		case Left:
			if currentColumn > 0 {
				currentColumn--
				currentIndex--
			}
		case Right:
			if currentColumn < lengths[currentLine] {
				currentColumn++
				currentIndex++
				if currentLine < len(lengths)-1 && currentColumn == lengths[currentLine] {
					currentColumn--
					currentIndex--
				}
			}
		case Home:
			currentIndex -= currentColumn
			currentColumn = 0
		case End:
			currentIndex += lengths[currentLine] - currentColumn
			currentColumn = lengths[currentLine]
			if currentLine < len(lengths)-1 {
				currentColumn--
				currentIndex--
			}
		case Enter:
			result = insertAtIndex(result, '\n', currentIndex)
			if currentLine == len(lengths)-1 {
				lengths = append(lengths, 0)
				prevLen := lengths[currentLine]
				lengths[currentLine] = currentColumn + 1
				currentLine++
				lengths[currentLine] = prevLen - currentColumn
			} else {
				prevLen := lengths[currentLine]
				lengths[currentLine] = currentColumn + 1
				lengths = append(lengths[:currentLine], append([]int{0}, lengths[currentLine:]...)...)
				lengths[currentLine] = currentColumn + 1
				currentLine++
				lengths[currentLine] = prevLen - currentColumn
			}
			currentColumn = 0
			currentIndex++
		case Up:
			if currentLine > 0 {
				currentLine--
				if currentColumn > lengths[currentLine]-1 {
					currentIndex -= (currentColumn + 1)
					currentColumn = lengths[currentLine]
					if currentLine < len(lengths)-1 {
						currentColumn--
					}
				} else {
					currentIndex -= lengths[currentLine]
				}
			}
		case Down:
			if currentLine < len(lengths)-1 {
				currentLen := lengths[currentLine]
				currentLine++
				if currentColumn > lengths[currentLine]-1 {
					currentIndex += ((currentLen - currentColumn) + lengths[currentLine])
					currentColumn = lengths[currentLine]
					if currentLine < len(lengths)-1 {
						currentIndex--
						currentColumn--
					}
				} else {
					currentIndex += currentLen
				}
			}
		default:
			result = insertAtIndex(result, char, currentIndex)
			currentIndex++
			currentColumn++
			lengths[currentLine]++
		}

	}

	return result
}

func task6() {
	var count int
	fmt.Scan(&count)
	results := make([]string, count)

	for i := 0; i < count; i++ {
		var text string
		fmt.Scan(&text)
		results[i] = terminal(text)
	}

	for i := 0; i < count; i++ {
		fmt.Print(results[i], "\n-")
		if i != len(results)-1 {
			fmt.Print("\n")
		}
	}
}
