// time complexity: O(n^2)
// space complexity: O(n^2)

func exist(board [][]byte, word string) bool {
    if len(board) == 0 {
        return false
    }
    visited := make([][]bool, len(board))
    for i:=0; i<len(visited); i++ {
        visited[i] = make([]bool, len(board[0]))
    }
    for i:=0; i<len(board); i++ {
        for j:=0; j<len(board[i]); j++ {
            if dfs(board, visited, word, 0, i, j) {
                return true
            }
        }
    }
    return false
}

func dfs(board [][]byte, visited [][]bool, word string, index, i, j int) bool {
    if i<0 || j<0 || i>=len(board) || j>=len(board[i]) || board[i][j]!=word[index] || visited[i][j] {
        return false
    }
    if index == len(word)-1 {
        return board[i][j] == word[index]
    }
    visited[i][j] = true
    for _,dir := range [][]int{{0, -1}, {0, 1}, {-1, 0}, {1, 0}} {
        if dfs(board, visited, word, index+1, i+dir[0], j+dir[1]) {
            return true
        }
    }
    visited[i][j] = false
    return false
}