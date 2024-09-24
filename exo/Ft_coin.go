package exo

// Cette fonction calcule le nombre minimum de pièces nécessaire pour atteindre un montant donné
func Ft_coin(coins []int, amount int) int {
    nb := make([]int, amount+1)

    // Initialise tableau avec une valeur par défaut
    for i := 1; i <= amount; i++ {
        nb[i] = amount + 1
    }

    // Calculer le nombre minimum de pièce
    for i := 1; i <= amount; i++ {
        for _, coin := range coins {
            if coin <= i { // Vérifie la valeur de coin
                nb[i] = min(nb[i], nb[i-coin]+1) // Correction de l'appel à min
            }
        }
    }

    // Si le montant est inatteignable
    if nb[amount] == amount+1 {
        return -1
    }
    return nb[amount]
}

// min retourne le minimum de deux entiers
func min(a, b int) int {
    if a < b {
        return a
    }
    return b
}
