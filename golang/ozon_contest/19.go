package ozoncontest

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

var scan = bufio.NewScanner(os.Stdout)

func task19XYZCount(msg string) (int, int, int) {
	x, y, z := 0, 0, 0
	for _, s := range msg {
		switch s {
		case 'X':
			x++
		case 'Y':
			y++
		case 'Z':
			z++
		}
	}

	return x, y, z
}

func Task19Handler(msg string, length int) bool {
	x, y, z := task19XYZCount(msg)
	temp := make([]bool, length)
	queue := make([]int, 0, length)

	for i := 0; i < length; i++ {
		if temp[i] {
			continue
		}
		switch msg[i] {
		case 'Z':
			if len(queue) == 0 {
				return false
			}
			index := queue[0]
			if index < i {
				sort.Slice(queue, func(i, j int) bool {
					return queue[i] > queue[j]
				})
				index = queue[0]
				if index < i {
					return false
				}
			}
			queue = queue[1:]
			temp[index] = false
			y++
			z--
		case 'Y':
			if z <= 0 || y+z <= x {
				if len(queue) == 0 {
					return false
				}
				index := queue[0]
				if index < i {
					return false
				}
				queue = queue[1:]
				temp[index] = false
			} else {
				for j := i + 1; j < length; j++ {
					if !temp[j] && msg[j] == 'Z' {
						temp[j] = true
						z--
						y--
						break
					}
					if j == length-1 {
						return false
					}
				}
			}
		case 'X':
			if x >= y+z {
				for j := i + 1; j < length; j++ {
					if !temp[j] {
						if msg[j] == 'Z' {
							temp[j] = true
							z--
							break
						} else if msg[j] == 'Y' {
							temp[j] = true
							y--
							break
						}
					}
					if j == length-1 {
						return false
					}
				}
			} else if z > y {
				for j := i + 1; j < length; j++ {
					if !temp[j] && msg[j] == 'Z' {
						temp[j] = true
						z--
						break
					}
					if j == length-1 {
						return false
					}
				}
			} else {
				for j := length - 1; j > i; j-- {
					if !temp[j] && msg[j] == 'Y' {
						temp[j] = true
						queue = append(queue, j)
						y--
						break
					}
					if j == i+1 {
						return false
					}
				}
			}
			x--
		}
	}

	return true
}

func Task19() {
	f, _ := os.Open("tmp3")
	var scan = bufio.NewScanner(f)
	scan.Scan()
	count, _ := strconv.Atoi(scan.Text())
	results := make([]bool, count)

	for i := 0; i < count; i++ {
		scan.Scan()
		length, _ := strconv.Atoi(scan.Text())
		scan.Scan()
		results[i] = Task19Handler(scan.Text(), length)
	}

	for i := 0; i < count; i++ {
		if results[i] {
			fmt.Print("Yes")
		} else {
			fmt.Print("No")
		}
		if i != count-1 {
			fmt.Print("\n")
		}
	}
}
