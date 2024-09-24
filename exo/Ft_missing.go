package exo

func Ft_missing(nums []int) int {
	n := len(nums)
	x := (n * (n + 1)) / 2
	a := 0
	for _, range nums {
		a += nums
	}
	nbmiss := x - a
	return nbmiss
}