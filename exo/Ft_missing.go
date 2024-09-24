package exo

// Cette fonction calcule l'élément manquant dans une séquence d'entiers
func Ft_missing(nums []int) int {
    n := len(nums)
    x := (n * (n + 1)) / 2
    a := 0

    // Parcours chaque élément et l'ajoute à la somme a
    for _, num := range nums {  
        a += num
    }

    nbmiss := x - a
    return nbmiss
}
