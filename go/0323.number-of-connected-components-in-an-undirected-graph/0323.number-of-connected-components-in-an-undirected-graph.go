package countComponents

// Union Find

func countComponents(n int, edges [][]int) int {
	parents := make([]int, n)
	rank := make([]int, n)
	for i := 0; i < n; i++ {
		parents[i] = i
		rank[i] = 1
	}

	count := n

	for _, edge := range edges {
		p1 := find(parents, edge[0])
		p2 := find(parents, edge[1])

		count -= union(parents, p1, p2, rank)

	}

	return count
}

func find(par []int, node int) int {
	res := node

	for res != par[res] {
		par[res] = par[par[res]]
		res = par[res]
	}

	return res
}

func union(par []int, n1, n2 int, rank []int) int {
	if n1 == n2 {
		return 0
	}

	if rank[n1] > rank[n2] {
		rank[n1] += rank[n2]
		par[n2] = par[n1]
	} else {
		rank[n2] += rank[n1]
		par[n1] = par[n2]
	}
	return 1
}

// DFS
func countComponents2(n int, edges [][]int) int {
	if n == 0 {
		return 0
	}
	adjMap := make(map[int][]int)
	visited := make(map[int]bool, n)

	// create graph
	for _, edge := range edges {
		start, end := edge[0], edge[1]
		adjMap[start] = append(adjMap[start], end)
		adjMap[end] = append(adjMap[end], start)
	}

	var count int
	for i := 0; i < n; i++ {
		if !visited[i] {
			count++
			search(adjMap, &visited, i)
		}
	}

	return count
}

func search(adjMap map[int][]int, visited *map[int]bool, pos int) {
	(*visited)[pos] = true

	nodes := adjMap[pos]

	for _, node := range nodes {
		if !(*visited)[node] {
			search(adjMap, visited, node)
		}
	}
}
