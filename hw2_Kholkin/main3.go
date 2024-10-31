package main

import (
    "fmt"
    "calculator"
)

func main(){
	var s []Squarer
	for i:=0; i<=10; i++ {
		s = append(s, circle{point{0, 0}, float64(i)})
	}
	fmt.Println(TotalArea(s))
	for i:=0; i<=10; i++ {
		s = append(s, rectangle{point{0,0}, point{float64(i), float64(i)}})
	}
	fmt.Println(TotalArea(s))
	c1 := circle{point{0,0}, 100}
	c2 := circle{point{0,0}, 10}
	r1 := rectangle{point{-10,10}, point{100,-15}}
	r2 := rectangle{point{10,10}, point{20,20}}
	fmt.Println(TotalArea([]Squarer{c1, c2, r1, r2}))
}
