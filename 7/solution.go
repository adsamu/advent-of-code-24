package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
	"sync"
	"runtime"
	"time"
)

type job struct {
	target   int
	sequence []int
}

func concat(a, b int) int {
	digits := int(math.Log10(float64(b))) + 1
	return a*int(math.Pow10(digits)) + b
}

func solve2(target int, sequence []int) bool {
	rows := len(sequence)

	dp := make([]map[int]bool, rows)
	for i := range dp {
		dp[i] = make(map[int]bool)
	}

	dp[0][sequence[0]] = true
	for i := 1; i < rows; i++ {
		for prevValue := range dp[i-1] {
			if prevValue > target {
				continue
			}
			dp[i][prevValue+sequence[i]] = true
			dp[i][prevValue*sequence[i]] = true
			dp[i][concat(prevValue, sequence[i])] = true

		}
	}

	if dp[rows-1][target] {
		return true
	}
	return false
}

func solve1(target int, sequence []int) bool {
	rows := len(sequence)

	dp := make([]map[int]bool, rows)
	for i := range dp {
		dp[i] = make(map[int]bool)
	}

	dp[0][sequence[0]] = true

	for i := 1; i < rows; i++ {
		for prevValue := range dp[i-1] {
			if prevValue > target {
				continue
			}
			dp[i][prevValue+sequence[i]] = true
			dp[i][prevValue*sequence[i]] = true
		}
	}

	if dp[rows-1][target] {
		return true
	}
	return false
}

func worker(jobs <-chan job, results chan<- int, wg *sync.WaitGroup, solve func(int, []int) bool) {
	defer wg.Done()
	for j := range jobs {
		if solve(j.target, j.sequence) {
			results <- j.target
		}
	}
}

func main() {
	start := time.Now()
	scanner := bufio.NewScanner(os.Stdin)
	var values []int
	var sequences [][]int

	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if len(line) > 0 {
			temp := strings.Split(line, ":")
			value, _ := strconv.Atoi(temp[0])
			values = append(values, value)

			sequenceStrings := strings.Fields(temp[1])
			sequence := make([]int, len(sequenceStrings))
			for i, s := range sequenceStrings {
				sequence[i], _ = strconv.Atoi(s)
			}
			sequences = append(sequences, sequence)
		}
	}

	result := make(chan int)
	jobs := make(chan job, len(values))
	var wg sync.WaitGroup

	numWorkers := runtime.NumCPU()
	for i := 0; i < numWorkers; i++ {
		wg.Add(1)
		go worker(jobs, result, &wg, solve2)
	}

	for i := 0; i < len(values); i++ {
		jobs <- job{values[i], sequences[i]}
	}
	close(jobs)

	go func() {
		wg.Wait()
		close(result)
	}()

	sum := 0
	for r := range result {
		sum += r
	}
	fmt.Println(time.Since(start), sum)
}

