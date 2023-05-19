package tasks

import "fmt"

func CollatzSequence() {

	longest, pos := 0, 0
	for i := 1; i < 1_000_000; i++ {
		len := collatzLength(i)
		if len > longest {
			longest = len
			pos = i
		}
	}

	fmt.Println(pos)
}

func collatzLength(n int) int {

	count := 0
	for {
		count++
		if n == 1 {
			return count
		}
		if n%2 == 0 {
			n = n / 2
		} else {
			n = 3*n + 1
		}
	}
}
