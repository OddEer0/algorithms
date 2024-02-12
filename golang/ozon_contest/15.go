package ozoncontest

import (
	"encoding/json"
	"fmt"
	"strconv"
	"strings"
	"sync"
)

type (
	Dirs struct {
		Dir     string   `json:"dir"`
		Files   []string `json:"files,omitempty"`
		Folders []Dirs   `json:"folders,omitempty"`
	}

	Task15Result struct {
		index int
		res   int
	}
)

func Task15CheckHack(files []string) bool {
	for _, file := range files {
		if strings.Contains(file, ".hack") {
			return true
		}
	}

	return false
}

func Task15Reduce(dir *Dirs, isHack bool) int {
	res := 0
	if !isHack && len(dir.Files) > 0 {
		isHack = Task15CheckHack(dir.Files)
	}

	if isHack {
		res = len(dir.Files)
	}

	if len(dir.Folders) > 0 {
		for _, d := range dir.Folders {
			res += Task15Reduce(&d, isHack)
		}
	}
	return res
}

func Task15Input(count int) []string {
	jsons := make([]string, count)
	for i := 0; i < count; i++ {
		scan.Scan()
		line, _ := strconv.Atoi(scan.Text())
		var s strings.Builder
		for j := 0; j < line; j++ {
			scan.Scan()
			s.WriteString(scan.Text())
		}
		jsons[i] = s.String()
	}
	return jsons
}

func Task15Handle(wg *sync.WaitGroup, index int, results chan<- Task15Result, jsonStr string) {
	defer wg.Done()
	var root Dirs
	err := json.Unmarshal([]byte(jsonStr), &root)
	if err != nil {
		fmt.Println("error unmarshal")
	}

	results <- Task15Result{res: Task15Reduce(&root, false), index: index}
}

func Task15() {
	scan.Scan()
	count, _ := strconv.Atoi(scan.Text())
	results := make([]int, count)
	resultsChan := make(chan Task15Result, count)
	defer close(resultsChan)

	jsons := Task15Input(count)

	var wg sync.WaitGroup
	for i := 0; i < count; i++ {
		wg.Add(1)
		go Task15Handle(&wg, i, resultsChan, jsons[i])
	}

	go func() {
		for i := 0; i < count; i++ {
			result := <-resultsChan
			results[result.index] = result.res
		}
	}()

	wg.Wait()

	for i := 0; i < count; i++ {
		fmt.Print(results[i])
		if i != count-1 {
			fmt.Print("\n")
		}
	}
}
