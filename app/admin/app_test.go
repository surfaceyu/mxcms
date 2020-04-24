package main

import (
	"fmt"
	"testing"
)

func TestStart(t *testing.T) {
	a := []int{1,2,3,4,5,6,7,8}
	//b := []int{}
	//b = append(b, append(a[:3], a[4:]...)...)

	for i := 0; i<len(a);i++  {
		fmt.Println(i,a[i])
		k:=i
		//v := a[i]
		if k >= 2 {
			a = append(a[:k], a[k+1:]...)
			i--
		}
	}
	fmt.Println(a)
}
