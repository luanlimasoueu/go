package main

import "fmt"

func main() {
	i, j := 42, 2701
	fmt.Println(i, j)
	fmt.Println(&i, &j)

	p := &i

	fmt.Println(*p)
	fmt.Println("%T\n", p)
	*p = 21
	fmt.Println(i)

	p = &j
	*p = *p / 37
	fmt.Println(j)

	a := 4
	squareVal(a)
	squareAdd(&a)
	fmt.Println(&a)

}

func squareVal(v int) {

	v *= v
	fmt.Println(&v, v)

}

func squareAdd(p *int) {

	*p *= *p
	fmt.Println(p, *p)

}
