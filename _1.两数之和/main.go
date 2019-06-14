package main

import "fmt"

/*
给定一个整数数组 nums 和一个目标值 target，
请你在该数组中找出和为目标值的那 两个 整数，并返回他们的数组下标。
你可以假设每种输入只会对应一个答案。
但是，你不能重复利用这个数组中同样的元素。

示例:
给定 nums = [2, 7, 11, 15], target = 9
因为 nums[0] + nums[1] = 2 + 7 = 9
所以返回 [0, 1]
来源：力扣（LeetCode）
 */
//空间复杂度O(1) 空间换时间
func twoSum(nums []int, target int) []int {
	m := make(map[int]int)
	for index,value := range nums{
		//1.判断一下[target-v]在不在map的key中
		if v,ok := m[target-value];ok{
			return []int{v,index}
		}
		//2.不在，就将map[v] = i
		m[value] = index
	}
	return nil
}

func main(){
	nums := []int{2,7,11,15}
	target := 9
	result:=twoSum(nums,target)
	fmt.Println(result)
}