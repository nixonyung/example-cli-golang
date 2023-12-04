package algorithms

func LCM(x, y uint) uint {
	return x / GCD(x, y) * y
}
