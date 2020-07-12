package linked_list

//set存储
func detectCycle1(head *ListNode) *ListNode {
	hash := make(map[*ListNode]int)
	for head != nil {
		if _, ok := hash[head]; ok {
			return head
		}
		hash[head] = head.Val
		head = head.Next
	}
	return nil
}

//快慢指针  先找到是否有环
//从头节点和相遇点 两个点出发 找环的入口
func detectCycle(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return nil
	}
	var slow = head
	var fast = head
	for fast.Next != nil && fast.Next.Next != nil {
		if fast.Next == nil {
			return nil
		}
		fast = fast.Next.Next
		slow = slow.Next
		if fast == slow {
			for head != fast {
				head = head.Next
				fast = fast.Next
			}
			return head
		}
	}
	return nil
}
