package main

import (
	"fmt"

	"time"
)

func main() {

	i := 2
	fmt.Println("Write", i, "as")
	switch i {
	case 1:
		fmt.Println("one")

	case 2:
		fmt.Println("two")

	case 3:
		fmt.Println("three")
	}

	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("it's a weeken bitch")
	default:
		fmt.Println("it's a weekday workers")
	}

	// you need to define the time first
	t := time.Now()

	switch {
	case t.Hour() < 12:
		fmt.Println("it's still morning back to bed")
	default:
		fmt.Println("it's the day")
	}

	myType := func(i interface{}) {
		switch t := i.(type) {
		case bool:
			fmt.Println("my type is bool")
		case int:
			fmt.Println("my type is int")
		case string:
			fmt.Println("my type is a string")
		default:
			fmt.Println("i dont know your type", t)
		}
	}

	myType(1)
	myType(true)
	myType("seyi")

	var a [5]int
	fmt.Println("amp:", a)

	a[4] = 100
	fmt.Println("beta:", a)
	fmt.Println("alpha:", a[4])
	fmt.Println("len:", len(a))

	b := [5]int{1, 2, 3, 4, 5}
	fmt.Println("DOPA:", b)

	b = [...]int{50, 3: 400, 500}
	fmt.Println("antidopa:", b)

	var TWOd [2][3]int
	for i := 0; i < 2; i++ {
		for j := 0; j < 3; j++ {
			TWOd[i][j] = i + j
		}
	}
	fmt.Println("cd array:", TWOd)

	var s []string
	fmt.Println("uninit:", s, s == nil, len(s) == 0)

	s = make([]string, 3)
	fmt.Println("emp", s, "len", len(s), "cap", cap(s))

	s[0] = "a"
	s[1] = "b"
	s[2] = "c"
	fmt.Println("set:", s)
	fmt.Println("get:", s[2])

	fmt.Println("len:", len(s))
	s = append(s, "d")
	s = append(s, "e", "f")
	fmt.Println("apd:", s)

	nums := []int{2, 3, 4}
	sum := 0
	for _, num := range nums {
		sum += num
	}
	fmt.Println("sum:", sum)

	for i, num := range nums {
		if num == 3 {
			fmt.Println("index", i)
		}

	}
	for i, c := range "go" {
		fmt.Println(i, c)
	}

}

//bvn 22535871047
