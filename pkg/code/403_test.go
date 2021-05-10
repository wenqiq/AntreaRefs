package code

import (
	"fmt"
	"testing"
)

func TestCanCross(t *testing.T) {
	//[0,1,3,5,6,8,12,17]
	stones := []int{0,1,3,5,6,8,12,17}
	res := CanCross(stones)
	fmt.Println(res)
}
