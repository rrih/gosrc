package main

import (
	"errors"
	"fmt"
)

type argError struct {
	arg  int
	prob string
}

func (e *argError) Error() string {
	return fmt.Sprintf("%d - %s", e.arg, e.prob)
}

func f1(arg int) (int, error) {
	if arg == 42 {
		// errors.New() はエラーメッセージを引数に error 型の値を作成する
		return -1, errors.New("can't work with 42")
	}
	// error 型の nil はエラーがなかったことを示す
	return arg + 3, nil
}

func f2(arg int) (int, error) {
	if arg == 42 {
		// argError 型のオブジェクトを作成
		return -1, &argError{arg, "can't work with it"}
	}
	return arg + 3, nil
}

func main() {
	// 明示的にエラーを返す
	for _, i := range []int{7, 42} {
		if r, e := f1(i); e != nil {
			fmt.Println("f1 failed:", e)
		} else {
			fmt.Println("f1 worked:", r)
		}
	}
	for _, i := range []int{7, 42} {
		if r, e := f2(i); e != nil {
			fmt.Println("f2 failed:", e)
		} else {
			fmt.Println("f2 worked:", r)
		}
	}

	_, e := f2(42)
	if ae, ok := e.(*argError); ok {
		fmt.Println(ae.arg)
		fmt.Println(ae.prob)
	}
}
