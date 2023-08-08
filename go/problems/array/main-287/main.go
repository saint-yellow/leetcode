// LeetCode 主站 Problem Nr. 287: 寻找重复数

package main

import (
	"fmt"
)

func main() {
	numbers := []int{1, 4, 4, 2, 2}
	fmt.Println(findDuplicate(numbers))
}

func findDuplicate(numbers []int) int {
	return method1(numbers)
}

// 双指针解法
// 将问题等价为在环形链表中寻找环的入口
func method1(numbers []int) int {
	slow, fast := 0, 0
	slow = numbers[slow]
	fast = numbers[numbers[fast]]
	for slow != fast {
		slow = numbers[slow]
		fast = numbers[numbers[fast]]
	}

	pointer1, pointer2 := 0, slow
	for pointer1 != pointer2 {
		pointer1 = numbers[pointer1]
		pointer2 = numbers[pointer2]
	}

	return pointer1
}
