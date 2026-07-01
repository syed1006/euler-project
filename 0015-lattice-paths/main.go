/*
<p>Starting in the top left corner of a $2 \times 2$ grid, and only being able to move to the right and down, there are exactly $6$ routes to the bottom right corner.</p>
<div class="center">
<img src="resources/images/0015.png?1678992052" class="dark_img" alt=""></div>
<p>How many such routes are there through a $20 \times 20$ grid?</p>
*/

package main

import "fmt"

func main() {

	fmt.Println(findNumberOfPaths(Point{
		X: 0,
		Y: 0,
	}, Point{
		X: 20,
		Y: 20,
	}))

	fmt.Println(numberOfPaths(20, 20))
	fmt.Println(choose(40, 20))
}

func NewBoolGrid(rows, cols int) [][]bool {
	grid := make([][]bool, rows)

	for i := range grid {
		grid[i] = make([]bool, cols)
	}

	return grid
}

type Point struct {
	X int
	Y int
}

var cache = make(map[string]int)

func findNumberOfPaths(start, target Point) int {

	current := findPath(start.X, start.Y, target.X, target.Y)

	return current

}
func findPath(startX, startY, targetX, targetY int) int {
	if startX == targetX && startY == targetY {
		return 1
	}
	if startX > targetX || startY > targetY {
		return 0
	}

	key := fmt.Sprintf("%v:%v", startX, startY)

	if value, ok := cache[key]; ok {
		return value
	}

	paths := findPath(startX+1, startY, targetX, targetY) + findPath(startX, startY+1, targetX, targetY)

	cache[key] = paths

	return paths
}

// 137846528820

// dp
func numberOfPaths(rows, cols int) uint64 {
	dp := make([][]uint64, rows+1)

	for i := range dp {
		dp[i] = make([]uint64, cols+1)
	}

	// First row
	for j := 0; j <= cols; j++ {
		dp[0][j] = 1
	}

	// First column
	for i := 0; i <= rows; i++ {
		dp[i][0] = 1
	}

	for i := 1; i <= rows; i++ {
		for j := 1; j <= cols; j++ {
			dp[i][j] = dp[i-1][j] + dp[i][j-1]
		}
	}

	return dp[rows][cols]
}

// math

func choose(n, k int) uint64 {
	if k > n-k {
		k = n - k
	}

	var result uint64 = 1

	for i := 1; i <= k; i++ {
		result = result * uint64(n-k+i) / uint64(i)
	}

	return result
}
