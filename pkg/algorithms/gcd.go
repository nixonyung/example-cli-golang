package algorithms

func GCD(x, y uint) uint {
	if x < y {
		x, y = y, x
	}
	for y > 0 {
		x, y = y, x%y
	}
	return x
}
