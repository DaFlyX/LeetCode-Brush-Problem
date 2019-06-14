package main

/*
给出两个 非空 的链表用来表示两个非负的整数。
其中，它们各自的位数是按照 逆序 的方式存储的，并且它们的每个节点只能存储 一位 数字。

如果，我们将这两个数相加起来，则会返回一个新的链表来表示它们的和。

您可以假设除了数字 0 之外，这两个数都不会以 0 开头。
示例：

输入：(2 -> 4 -> 3) + (5 -> 6 -> 4)
输出：7 -> 0 -> 8
原因：342 + 465 = 807

来源：力扣（LeetCode）
*/
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	//1.初始化节点
	var node=&ListNode{0,nil}
	//2.定义 进位
	var temp = 0
	//3.定义返回值
	result := node
	//3.for循环加和，判断l1!=nil || l2!=nil||temp!=0
	for l1!=nil || l2!=nil|| temp!=0{
		var v1,v2 int
		if l1!=nil{
			v1 = l1.Val
			l1 = l1.Next
		}
		if l2!=nil{
			v2 = l2.Val
			l2 = l2.Next
		}
		sum := v1 + v2 + temp //从个位计算，个位如果满十则进位（再取个位数也就是temp=0）
		temp = sum/10 //取整（不是1就是0）1代表进位 0代表加和不超过十不用进位
		result.Next = &ListNode{sum%10,nil} //
		result = result.Next
	}
	return node.Next
}
