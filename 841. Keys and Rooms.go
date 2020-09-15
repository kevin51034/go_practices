// DFS

func canVisitAllRooms(rooms [][]int) bool {
    roomMap := make(map[int]bool, len(rooms))
    roomMap[0] = true
    dfs(rooms, roomMap, 0)
    return len(rooms) == len(roomMap)
}

func dfs(rooms [][]int, roomMap map[int]bool, from int) {
    for _,to := range rooms[from] {
        if roomMap[to] == true {
            continue
        }
        roomMap[to] = true
        dfs(rooms, roomMap, to)
    }
}

// BFS

func canVisitAllRooms(rooms [][]int) bool {
    roomMap := make(map[int]bool, len(rooms))
    roomMap[0] = true
    queue := make([]int, 0)
    for i:=0; i<len(rooms[0]); i++ {
        queue = append(queue, rooms[0][i])
    }
    
    for len(queue) > 0 {
        for l:=len(queue)-1; l>=0; l-- {
            node := queue[0]
            queue = queue[1:]
            if roomMap[node] == true {
                continue
            }
            roomMap[node] = true
            for i:=0; i<len(rooms[node]); i++ {
                queue = append(queue, rooms[node][i])
            }
        }
    }
    fmt.Println(roomMap)
    for i:=0; i<len(rooms); i++ {
        if roomMap[i] == false {
            return false
        }
    }
    return true
}