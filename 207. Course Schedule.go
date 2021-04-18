// Time complexity: O(n)

func canFinish(numCourses int, prerequisites [][]int) bool {
	graph := make([][]int, numCourses)
	for _, p := range prerequisites {
		graph[p[0]] = append(graph[p[0]], p[1])
	}

	// states: 0 = unkonwn, 1 == visiting, 2 = visited
	v := make([]int, numCourses)

	for i := 0; i < numCourses; i++ {
		// check if there is cycle in courses
		if dfs(i, v, graph) {
			return false
		}
	}
	return true
}

func dfs(cur int, v []int, graph [][]int) bool {
	// if this course is visiting means there is a cycle, v[cur] == 1
	if v[cur] == 1 {
		return true
	} else if v[cur] == 2 {
		return false
	}

	// set this course visiting
	v[cur] = 1
	for _, c := range graph[cur] {
		if dfs(c, v, graph) {
			return true
		}
	}
	// set this course visited
	v[cur] = 2
	return false
}