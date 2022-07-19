package main

import (
	"errors"
	"math"
)

//第四题
type node struct {
	val  int
	pre  *node
	next *node
}

func game(n int, m int) (int, error) {
	if !(n >= 1 && n <= int(math.Pow10(5)) && m >= 1 && m <= int(math.Pow10(6))) {
		return 0, errors.New("gama over")
	}

	//创建新游戏数据
	header := newGame(n)

	number := n
	i := 0
	curr := header
	for {
		//只有一个就是胜利者
		if number == 1 {
			return curr.val, nil
		}

		//开始计数
		i++
		if i == m {
			//找到前驱和后驱
			pre := curr.pre
			next := curr.next

			//在链上去掉当前节点
			pre.next = next
			next.pre = pre

			//数量减一
			number--
			//计数器归零
			i = 0
		}

		//修改当前节点到下一个节点。
		curr = curr.next
	}

}

func newGame(n int) *node {

	header := &node{0, nil, nil}
	curr := header

	for i := 1; i < n; i++ {
		newNode := node{i, nil, nil}
		newNode.pre = curr
		curr.next = &newNode
		curr = &newNode
	}
	// 让它首尾相连。
	curr.next = header
	header.pre = curr

	return header
}
