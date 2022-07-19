package main

import (
	"fmt"
)

func main() {

	//第一题
	arr := []int64{1, 2, 19, 8}
	maxdiff := maxDiff(arr)
	fmt.Println(maxdiff)

	//第二题
	shuArr := shuffle()
	fmt.Println(shuArr)

	//第三题
	cManager := cacheManager{m: make(map[string]cache)}
	cManager.add("tom", "123", 123)
	val, _ := cManager.get("tom")
	fmt.Println(val)
	fmt.Println(cManager)

	//第四题
	x, _ := game(3, 2)
	fmt.Println(x)
}
