package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello, playground")

	wftm := 3260989 // amt of token A
	geist := 2362145 // amt of token B
	
	pwftm := 2.08 // price of token A

	fmt.Printf("liquidity of wftm: %v\n", wftm)
	fmt.Printf("liquidity of geist: %v\n", geist)

	pgeist := float64(wftm) * pwftm
	pgeist = pgeist / float64(geist) // price of token B

	k := wftm * geist // constant product of the pool

	fmt.Printf("k: %v\n", k)

	fmt.Printf("price of geist: %f\n", pgeist)
	
	liqRemoved := 1900000

	wftm = wftm - liqRemoved // reduced token A liq

	geist2 := k / wftm

	fmt.Printf("liquidity of new wftm: %v\n", wftm)
	fmt.Printf("liquidity of new geist: %v\n", geist2)

	pgeist2 := float64(wftm) * pwftm 
	pgeist2 = pgeist2 / float64(geist2) // new price of token B

	fmt.Printf("price of new geist: %v\n", pgeist2)

}