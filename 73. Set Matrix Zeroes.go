// Time complexity: O(m*n)
// Space complexity: O(1)

func setZeroes(matrix [][]int)  {
    firstcolzero, firstrowzero := false, false
    // check if there is zero in first column
	// because we use first column/row to store if there is a zero in that column/zero
	// so we need to remember if there has leading zero at the begining
    for i:=0; i<len(matrix); i++ {
        if matrix[i][0] == 0 {
            firstcolzero = true
        }
    }
    // check if there is zero in first row
    for i:=0; i<len(matrix[0]); i++ {
        if matrix[0][i] == 0 {
            firstrowzero = true
        }
    }
    // using first column/row to store if there is a zero in that column/zero
    for i:=1; i<len(matrix); i++ {
        for j:=1; j<len(matrix[i]); j++ {
            if matrix[i][j] == 0 {
                matrix[i][0] = 0
                matrix[0][j] = 0
            }
        }
    }
    // change the value if first column is zero
    for i:=1; i<len(matrix); i++ {
        if matrix[i][0] == 0 {
            for j:=0; j<len(matrix[i]); j++ {
                matrix[i][j] = 0
            }
        }
    }
    // change the value if first row is zero
    for j:=1; j<len(matrix[0]); j++ {
        if matrix[0][j] == 0 {
            for i:=0; i<len(matrix); i++ {
                matrix[i][j] = 0
            }
        }
    }
    // because there might have leading zero at the begining, so check again
    if firstcolzero {
        for i:=0; i<len(matrix); i++ {
            matrix[i][0] = 0
        }
    }
    if firstrowzero {
        for j:=0; j<len(matrix[0]); j++ {
            matrix[0][j] = 0
        }
    }
}