package main

import "unsafe"

// 假设此函数不会被内联（inline）。
func createInt() *int {
	return new(int)
}

func foo() {
	p0, y, z := createInt(), createInt(), createInt()
	var p1 = unsafe.Pointer(y) // 和y一样引用着同一个值
	var p2 = uintptr(unsafe.Pointer(z))

	// 此时，即使z指针值所引用的int值的地址仍旧存储
	// 在p2值中，但是此int值已经不再被使用了，所以垃圾
	// 回收器认为可以回收它所占据的内存块了。另一方面，
	// p0和p1各自所引用的int值仍旧将在下面被使用。

	// uintptr值可以参与算术运算。
	p2 += 2; p2--; p2--

	*p0 = 1                         // okay
	*(*int)(p1) = 2                 // okay
	*(*int)(unsafe.Pointer(p2)) = 3 // 危险操作！
}
