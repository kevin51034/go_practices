func rotate(matrix [][]int)  {
    if len(matrix) <= 1 {
        return
	}
	// rotate by diagonal
    for i:=0; i<len(matrix); i++ {
        for j:=i+1; j<len(matrix[i]); j++ {
            matrix[i][j] , matrix[j][i] = matrix[j][i], matrix[i][j]
        }
	}
	// rotate by vertical centerline
    for i:=0; i<len(matrix); i++ {
        for j:=0; j<len(matrix[i])/2; j++ {
            matrix[i][j], matrix[i][len(matrix[i])-j-1] = matrix[i][len(matrix[i])-j-1], matrix[i][j]
        }
    }
    return
}