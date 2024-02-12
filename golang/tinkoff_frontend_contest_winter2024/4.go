package tinkoff_frontend_contest_winter2024

import (
	"fmt"
)

type Node struct {
	parent   int
	price    int
	company  string
	children []*Node
}

func depthAirstSearch(node *Node, companies map[string]bool, coveredCompanies map[string]bool) (int, bool) {
	price := 0

	if companies[node.company] && !coveredCompanies[node.company] {
		price = node.price
		delete(companies, node.company)
		coveredCompanies[node.company] = true
	}

	for _, child := range node.children {
		childPrice, _ := depthAirstSearch(child, companies, coveredCompanies)
		price += childPrice
	}

	return price, len(companies) == 0
}

func Task4() {
	var n, k int
	fmt.Scan(&n, &k)
	companies := make(map[string]bool)
	for i := 0; i < k; i++ {
		var company string
		fmt.Scan(&company)
		companies[company] = true
	}
	nodes := make([]*Node, n+1)
	for i := 1; i <= n; i++ {
		node := &Node{}
		fmt.Scan(&node.parent, &node.price, &node.company)

		nodes[i] = node
		if node.parent != 0 {
			nodes[node.parent].children = append(nodes[node.parent].children, node)
		}
	}

	root := nodes[1]
	covered := make(map[string]bool)
	result, _ := depthAirstSearch(root, companies, covered)

	if len(companies) == 0 {
		fmt.Println(result)
	} else {
		fmt.Println(-1)
	}
}
