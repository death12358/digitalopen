package note

import "fmt"

func increase(n *int) {
	*n++
	fmt.Printf("increase後=:%v n的地址:%v\n", n, &n)
}

func Pointer() {
	var src = 2024
	var ptr = &src
	ptr = new(int)
	ptr2 := &ptr
	increase(ptr)
	fmt.Printf("ptr=%v ptr的地址:%v ptr的值:%v\n", ptr, &ptr, *ptr)
	fmt.Printf("ptr2=%v ptr2的地址:%v ptr2的值:%v \n", ptr2, &ptr2, *ptr2)
}

func Fmt() {
	a := 1000000
	fmt.Printf("%%x:%x %%d:%d %%b:%b", a, a, a)
}

func Label() {

	for i := 0; i < 10; i++ {

		for j := 0; j < 10; j++ {
			fmt.Print("+")

			if j == i {
				fmt.Println(1)

			}
			continue
		}
		fmt.Println()
	}
}
