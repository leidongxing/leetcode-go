package linked_list

//二进制链表转整数
func getDecimalValue(head *ListNode) int {
	result := head.Val
	for head.Next != nil {
		result = result*2 + head.Next.Val
		head = head.Next
	}
	return result
}
