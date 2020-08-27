func twoSum(numbers []int, target int) []int {
    another := make(map[int]int)
    for i,v := range numbers {
        if _,ok := another[target-v]; ok {
            return []int{another[target-v]+1, i+1}
        }
        another[v] = i
    }
    return nil
}

// optimize

func twoSum(numbers []int, target int) []int {
    i, j := 0, len(numbers)-1
    for i < j {
        if numbers[i] + numbers[j] == target {
            return []int{i+1, j+1}
        } else if numbers[i] + numbers[j] < target {
            i++
        } else {
            j--
        }
    }
    return nil
}