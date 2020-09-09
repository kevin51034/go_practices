func floodFill(image [][]int, sr int, sc int, newColor int) [][]int {
    if len(image) == 0 {
        return image
    }
    dfs(&image, sr, sc, image[sr][sc], newColor)
    return image
}

func dfs(image *[][]int, i, j, startColor, newColor int) {
    if i<0 || j<0 || i>=len((*image)) || j>=len((*image)[i]) || (*image)[i][j] == newColor || (*image)[i][j] != startColor {
        return
    }
    (*image)[i][j] = newColor
    dfs(image, i-1, j, startColor, newColor)
    dfs(image, i+1, j, startColor, newColor)
    dfs(image, i, j-1, startColor, newColor)
    dfs(image, i, j+1, startColor, newColor)
}