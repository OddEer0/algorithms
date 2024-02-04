package ozoncontest

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type (
	Data struct {
		parentId int
		message  string
	}

	Tree struct {
		root   *Node
		length int
	}

	Node struct {
		key      int
		message  string
		children []*Node
	}
)

func (n *Node) appendManyData(data map[int]*Data) {
	if len(data) != 0 {
		for key, item := range data {
			if n.key == item.parentId {
				n.children = append(n.children, &Node{
					key:      key,
					message:  item.message,
					children: make([]*Node, 0, len(data)/2),
				})
				delete(data, key)
			}
		}

		sort.Slice(n.children, func(i, j int) bool {
			return n.children[i].key < n.children[j].key
		})
		for _, node := range n.children {
			node.appendManyData(data)
		}
	}
}

func buildTree(data map[int]*Data, length int) *Tree {
	result := &Tree{
		root: &Node{
			key:      -1,
			message:  "",
			children: make([]*Node, 0, len(data)/2),
		},
		length: length,
	}

	result.root.appendManyData(data)

	return result
}

var scanner = bufio.NewScanner(os.Stdin)

func messageInput() map[int]*Data {
	scanner.Scan()
	messageCount, _ := strconv.Atoi(scanner.Text())
	data := make(map[int]*Data, messageCount)

	for i := 0; i < messageCount; i++ {
		var (
			id       int
			dataItem = &Data{}
		)
		scanner.Scan()
		line := strings.Split(scanner.Text(), " ")
		id, _ = strconv.Atoi(line[0])
		dataItem.parentId, _ = strconv.Atoi(line[1])
		dataItem.message = strings.Join(line[2:], " ")
		data[id] = dataItem
	}

	return data
}

func printHelper(lines []int) {
	for i := 0; lines[i] != -1; i++ {
		if i == 0 {
			if lines[i] == 1 {
				fmt.Print("|")
			} else {
				fmt.Print(" ")
			}
		} else {
			if lines[i] == 1 {
				fmt.Print("  |")
			} else {
				fmt.Print("   ")
			}
		}
	}
}

func (n *Node) print(lines []int, index, brotherCount int) {
	if index == 0 {
		fmt.Println(n.message)
	} else {
		printHelper(lines)
		fmt.Print("--")
		fmt.Print(n.message, "\n")
	}

	if brotherCount == 0 && index > 0 {
		lines[index-1] = 0
	}

	if len(n.children) > 0 {
		lines[index] = 1
		lines[index+1] = -1
	}

	for i, node := range n.children {
		if i == 0 {
			printHelper(lines)
			fmt.Print("\n")
		}

		node.print(lines, index+1, len(n.children)-1-i)
		if i == len(n.children)-1 {
			lines[index] = -1
		} else {
			printHelper(lines)
			fmt.Print("\n")
		}
	}
}

func (t *Tree) print(lines []int) {
	for i, node := range t.root.children {
		node.print(lines, 0, len(t.root.children)-1)
		if i != len(t.root.children)-1 {
			fmt.Print("\n")
		}
	}
}

func task10() {
	scanner.Scan()
	count, _ := strconv.Atoi(scanner.Text())
	results := make([]*Tree, count)

	for i := 0; i < count; i++ {
		messagesData := messageInput()
		results[i] = buildTree(messagesData, len(messagesData))
	}

	for i, tree := range results {
		lines := make([]int, tree.length)
		if tree.length > 0 {
			lines[0] = -1
			tree.print(lines)
			if i != len(results)-1 {
				fmt.Print("\n")
			}
		}
	}
}
