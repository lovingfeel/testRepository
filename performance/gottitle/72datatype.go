package main

func main() {
	var (
		a int         = 0
		b int64       = 0
		c interface{} = int(0)
		d interface{} = int64(0)
		e int32         = 0
	)

	println(c == int(0))
	println(c == a)
	println(c == b)
	println(d == b)
	println(d == int64(0))
	println(e == rune(0))
}
