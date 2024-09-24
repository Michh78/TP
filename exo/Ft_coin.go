package exo

func Ft_coin(coins []int, amount int) int {
	nb := make([]int, amount + 1)

	for i = 1; i <= amount; i++ {
		nb[i] = amount + 1
	}
	for i = 1; i <= amount + 1; i++ {
		for _, range coins {
			if coins <= i {
				nb[i] = min(nb[i] - nb[i - coins]+1)
			}
		}
	}
	if nb[amount] == amount+1 {
		return -1
	}
	return nb[amount]
}

func min (a, b int) int {
	if a < b {
		return a
	}
	return b
}
