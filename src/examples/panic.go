package main

import "os"

// 予期せぬエラー発生時に panic を使う
func main() {

	panic("a problem")

	_, err := os.Create("/tmp/file")
	if err != nil {
		panic(err)
	}
}
