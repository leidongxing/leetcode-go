package main

//set存储
func detectCycle(head *ListNode) *ListNode {
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
