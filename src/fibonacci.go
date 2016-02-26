package main

import "flag"
import "fmt"
import "time"

func main() {
	var fibNum int
	flag.IntVar(&fibNum, "fib", 14, "fibonacci number to calculate")

	flag.Parse()

	start := time.Now()
	fibCache := make(map[int]int64)
	fibOut := fibCached(fibNum, fibCache)
	duration := time.Since(start).Nanoseconds()
	fmt.Printf("fibCached(%d)=%d, duration=%d(ns)\n", fibNum, fibOut, duration)

	start = time.Now()
	fibOut = fib(fibNum)
	duration = time.Since(start).Nanoseconds()
	fmt.Printf("fib(%d)=%d, duration=%d(ns)\n", fibNum, fibOut, duration)
}

// Uses a map to cache the results of prior calculations to reduce repeating the same calculations over and over
func fibCached(x int, m map[int]int64) int64 {
	i, ok := m[x]
	if ok {
		return i
	}

	if x == 0 {
		return 0
	}
	if x == 1 {
		return 1
	}
	curFib := fibCached(x - 1, m) + fibCached(x - 2, m)
	m[x] = curFib
	return curFib
}

func fib(x int) int64 {
	if x == 0 {
		return 0
	}
	if x == 1 {
		return 1
	}
	return fib(x - 1) + fib(x - 2)
}
