package main

import (
	"fmt"
)

func main() {
	var n, k, p, s int
	fmt.Scan(&n, &k, &p, &s)

	num_parties := (k + 1) / 2
	petya_party := (p + 1) / 2
	petya_variant := (p-1)%k + 1

	if petya_variant == 0 {
		petya_variant = k
	}

	vasya_party := petya_party + (petya_variant+s+1)/2
	if vasya_party > num_parties {
		vasya_party -= num_parties
	}

	if petya_party == vasya_party {
		fmt.Println("-01")
	} else {
		var vasya_side int
		if (petya_variant+s)%2 == 1 {
			vasya_side = 2
		} else {
			vasya_side = 1
		}
		fmt.Println(vasya_party*2-1, vasya_side)
	}
}
