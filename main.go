package main

import (
	"fmt"
    "github.com/Michh78/TP.git/exo" 
)

func main() {
	fmt.Println(exo.Ft_non_overlap([][]int{{1, 2}, {2, 3}, {3, 4}, {1, 3}})) // resultat : 1
	// pour que les intervalles soient tous des intervalles qui ne se superpose pas,
	// il suffit de d'enlever qu'un seul intervalle, [1,3]
	fmt.Println(exo.Ft_non_overlap([][]int{{1, 2}, {2, 3}})) // resultat : 0
	fmt.Println(exo.Ft_non_overlap([][]int{{1, 2}, {1, 2}, {1, 2}})) // resultat : 2
}

// exo Ft_coin
//fmt.Println(exo.Ft_coin([]int{1, 2, 5}, 11)) // resultat : 3 car (11 == 5 + 5 + 1)
	//fmt.Println(exo.Ft_coin([]int{2}, 3)) // resultat : -1
	//fmt.Println(exo.Ft_coin([]int{1}, 0)) // resultat : 0


// exo Ft_missing
//fmt.Println(exo.Ft_missing([]int{3, 1, 2}))        // Affiche : 0      
	//fmt.Println(exo.Ft_missing([]int{0, 1}))           // Affiche : 2
    //fmt.Println(exo.Ft_missing([]int{9, 6, 4, 2, 3, 5, 7, 0, 1})) // Affiche : 8

