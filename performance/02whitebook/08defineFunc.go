package main

import "log"

func HalfAndNegative(n int) (int, int) {
	return n/2, -n
}

func AddSub(a, b int) (int, int) {
	return a+b, a-b
}

func Dummy(values ...int) {}

/**
  1,一个有且只有一个返回值的函数的每个调用总可以被当成一个单值表达式使用。
   比如，它可以被内嵌在其它函数调用中当作实参使用，或者可以被当作其它表达式中的操作数使用。
 */
func main() {
	// 这几行编译没问题。
	AddSub(HalfAndNegative(6))
	AddSub(AddSub(AddSub(7, 5)))
	AddSub(AddSub(HalfAndNegative(6)))
	Dummy(HalfAndNegative(6))
	_, _ = AddSub(7, 5)

	Dummy(AddSub(7, 5), 9)
	// 下面这几行编译不通过。
	/*
		_, _, _ = 6, AddSub(7, 5)
		Dummy(AddSub(7, 5), 9)
		Dummy(AddSub(7, 5), HalfAndNegative(6))
	*/
}
func tang(y int )(x int,z int)  {
	log.Println("输出结果")
	return 3,2
}


