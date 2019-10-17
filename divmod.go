package piscine

func DivMod(a int, b int, div *int, mod *int) {
	c := a / b
	d := a % b
	*div = c
	*mod = d
}
