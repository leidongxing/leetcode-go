package linked_list

//set存储
func hasCycle1(head *ListNode) bool {
	hash := make(map[*ListNode]int)
	for head != nil {
		if _, ok := hash[head]; ok {
			return true
		}
		hash[head] = head.Val
		head = head.Next
	}
	return false
}

//特殊值修改链表
func hasCycle2(head *ListNode) bool {
	for head != nil {
		if head.Val == 18464187222 {
			return true
		}
		head.Val = 18464187222
		head = head.Next
	}
	return false
}

//快慢指针
func hasCycle(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return false
	}
	var slow = head
	var fast = head.Next
	for fast != nil && fast.Next != nil && slow != nil {
		if fast == slow {
			return true
		}
		slow = slow.Next
		fast = fast.Next.Next
	}
	return false
}
