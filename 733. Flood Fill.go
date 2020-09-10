// image doesn't have to pass in the pointer cuz it's slice

func floodFill(image [][]int, sr int, sc int, newColor int) [][]int {
    result := make([][]int, 0)
    if len(image) == 0 {
        return result
    }
    sColor := image[sr][sc]
    dfs(image, sr, sc, sColor, newColor)
    return image
}

func dfs(image [][]int, i, j, sColor, newColor int) {
    if i<0 || j<0 || i>=len(image) || j>=len(image[i]) || image[i][j]==newColor || image[i][j]!=sColor {
        return
    }
    image[i][j] = newColor
    dfs(image, i-1, j, sColor, newColor)
    dfs(image, i+1, j, sColor, newColor)
    dfs(image, i, j-1, sColor, newColor)
    dfs(image, i, j+1, sColor, newColor)
    return

}

// with pointer, in case image is not slice but array

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