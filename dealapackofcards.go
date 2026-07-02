package piscine

import (
	"fmt"
)

func DealAPackOfCards(deck []int) {
	//each := len(deck)/4
	//for i:=0; i<4; i++ {
	//	for j:=i*3;
	//}
	for i := 0; i < 4; i++ {
		fmt.Printf("Player %d:", i+1)
		for j := i * 3; j < i*3+1; j++ {
			fmt.Printf(" %d", deck[j])
			fmt.Printf(", %d", deck[j+1])
			fmt.Printf(", %d", deck[j+2])
		}
		fmt.Printf("\n")
	}
}
