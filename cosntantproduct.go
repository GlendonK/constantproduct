package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello, playground")

	wftm := 3260989
	geist := 2362145

	fmt.Printf("liquidity of wftm: %v\n", wftm)
	fmt.Printf("liquidity of geist: %v\n", geist)

	pgeist := float64(wftm) * 2.08
	pgeist = pgeist / float64(geist)

	k := wftm * geist

	fmt.Printf("k: %v\n", k)

	fmt.Printf("price of geist: %f\n", pgeist)

	wftm = wftm - 1900000

	geist2 := k / wftm

	fmt.Printf("liquidity of new wftm: %v\n", wftm)
	fmt.Printf("liquidity of new geist: %v\n", geist2)

	pgeist2 := float64(wftm) * 2.08
	pgeist2 = pgeist2 / float64(geist2)

	fmt.Printf("price of new geist: %v\n", pgeist2)

}
