// solution1
// iteration
func fib(N int) int {
	if N == 0 {
		return 0
	}
	if N == 1 {
		return 1
	}
	first := 0
	second := 1
	for i:=2; i<N; i++ {
		tmp := first + second
		first = second
		second = tmp
	}
	return first + second
}
// solution2
// recursion

func fib(N int) int {
    if N == 0 {
        return 0
    }
    if N == 1 {
        return 1
    }
    return fib(N-1) + fib(N-2)
}

// solution3
// use array to optimize

var m map[int]int = make(map[int]int, 0)
func fib(N int) int {
    if N == 0 {
        return 0
    }
    if N == 1 {
        return 1
    }
    if m[N]!=0 {
        return m[N]
    }
    m[N] = fib(N-1) + fib(N-2)
    return m[N]
}