package code

import (
	"fmt"
	"testing"
)

func TestLeastBricks(t *testing.T) {
	wall := [][]int{{1, 2, 2, 1}, {3, 1, 2}, {1, 3, 2}, {2, 4}}
	//wall := [][]int{[]int{1,2,2,1},[]int{3,1,2},[]int{2,4}}
	//wall := [][]int{[]int{1},[]int{1},[]int{1}}
	//wall := [][]int{[]int{100000000},[]int{100000000},[]int{100000000}}
	res := LeastBricks(wall)
	fmt.Println(res)
}
