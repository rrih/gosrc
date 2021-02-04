package main

import (
	"fmt"
	"time"
)

func main() {
	i := 2
	switch i {
	case 1:
		fmt.Println("onw")
	case 2:
		fmt.Println("two")
	case 3:
		fmt.Println("thee")
	}

	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("休みです")
	default:
		fmt.Println(time.Now().Weekday(), "は平日です")
	}

	t := time.Now()
	fmt.Println(t)
	switch {
	case t.Hour() < 12:
		fmt.Println("ごぜんだよー(´･_･`)")
	default:
		fmt.Println("ごごだよー(´･_･`)")
	}

	whatAmI := func(i interface{}) {
		switch typeee := i.(type) {
		case bool:
			fmt.Printf("私は %T です\n", typeee)
		case int:
			fmt.Printf("私は %T です\n", typeee)
		default:
			fmt.Printf("type -> %T\n", typeee)
		}
	}
	whatAmI(true)
	whatAmI(1)
	whatAmI("hoge")
}
