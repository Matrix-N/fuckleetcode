package fuckleetcode

import "sort"

func twoSum(nums []int, target int) []int {
	m := make(map[int]int)
	for i, v := range nums {
		if j, ok := m[target-v]; ok {
			return []int{j, i}
		}
		m[v] = i
	}
	return nil
}

func addTwoNumbers(l1, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}
	var head, tail *ListNode
	carry := 0
	for l1 != nil || l2 != nil || carry != 0 {
		var v1, v2 int
		if l1 != nil {
			v1 = l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			v2 = l2.Val
			l2 = l2.Next
		}
		sum := v1 + v2 + carry
		carry = sum / 10
		if head == nil {
			head = &ListNode{Val: sum % 10}
			tail = head
		} else {
			tail.Next = &ListNode{Val: sum % 10}
			tail = tail.Next
		}
	}

	return head
}

func lengthOfLongestSubstring(s string) int {
	if s == "" {
		return 0
	}

	left, right := 0, -1
	result := 0

	freq := [256]byte{}

	for left < len(s) {
		if right+1 < len(s) && freq[s[right+1]-'a'] == 0 {
			freq[s[right+1]-'a'] = 1
			right++
		} else {
			freq[s[left]-'a'] = 0
			left++
		}
		result = max(result, right-left+1)
	}

	return result
}

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	nums := append(nums1, nums2...)
	sort.Ints(nums)
	if len(nums)%2 == 0 {
		return float64(nums[len(nums)/2]+nums[len(nums)/2-1]) / 2
	}
	return float64(nums[len(nums)/2])
}

func longestPalindrome(s string) string {
	if s == "" {
		return ""
	}
	if len(s) == 1 {
		return s
	}
	start, end := 0, 0
	for i := 0; i < len(s); i++ {
		start1, end1 := expand(s, i, i)
		start2, end2 := expand(s, i, i+1)
		if end1-start1 > end-start {
			start, end = start1, end1
		}
		if end2-start2 > end-start {
			start, end = start2, end2
		}
	}

	return s[start : end+1]
}

func expand(s string, left, right int) (int, int) {
	for left >= 0 && right < len(s) && s[left] == s[right] {
		left--
		right++
	}

	return left + 1, right - 1
}

func maxArea(height []int) int {
	if height == nil || len(height) == 0 {
		return 0
	}

	left, right := 0, len(height)-1
	result := 0
	wid := 0
	high := 0

	for left <= right {
		wid = right - left
		if height[left] < height[right] {
			high = height[left]
			left++
		} else {
			high = height[right]
			right--
		}
		result = max(result, wid*high)
	}

	return result
}

func threeSum(nums []int) [][]int {
	if len(nums) == 0 {
		return [][]int{}
	}

	m := make(map[int]int)

	for _, v := range nums {
		m[v]++
	}

	u := []int{}

	for key := range m {
		u = append(u, key)
	}
	sort.Ints(u)

	var res [][]int

	for i := 0; i < len(u); i++ {
		if u[i]*3 == 0 && m[u[i]] >= 3 {
			res = append(res, []int{u[i], u[i], u[i]})
		}
		for j := i + 1; j < len(u); j++ {
			if u[i]*2+u[j] == 0 && m[u[i]] >= 2 {
				res = append(res, []int{u[i], u[i], u[j]})
			}
			if u[i]+u[j]*2 == 0 && m[u[j]] >= 2 {
				res = append(res, []int{u[i], u[j], u[j]})
			}
			c := 0 - u[i] - u[j]
			if c < u[i] && m[c] >= 1 {
				res = append(res, []int{u[i], u[j], c})
			}
		}
	}

	return res
}

var letterMap = []string{
	"",
	" ",
	"abc",
	"def",
	"ghi",
	"jkl",
	"mno",
	"pqrs",
	"tuv",
	"wxyz",
}

func helper(digits string, index int, c string, res *[]string) {
	if index == len(digits) {
		*res = append(*res, c)
		return
	}
	n := digits[index]
	nums := letterMap[n-'0']
	for i := range nums {
		helper(digits, index+1, c+string(nums[i]), res)
	}
}

func letterCombinations(digits string) []string {
	if digits == "" {
		return []string{}
	}

	var result []string
	var c string

	helper(digits, 0, c, &result)

	return result
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	if head == nil {
		return nil
	}
	if n == 0 {
		return head
	}

	fast, slow := head, head
	for i := 0; i < n; i++ {
		fast = fast.Next
	}
	if fast == nil {
		return head.Next
	}
	for fast.Next != nil {
		fast = fast.Next
		slow = slow.Next
	}
	slow.Next = slow.Next.Next

	return head
}

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}

	head := &ListNode{}
	cur := head

	for l1 != nil && l2 != nil {
		if l1.Val < l2.Val {
			cur.Next = l1
			l1 = l1.Next
		} else {
			cur.Next = l2
			l2 = l2.Next
		}
		cur = cur.Next
	}

	if l1 != nil {
		cur.Next = l1
	}
	if l2 != nil {
		cur.Next = l2
	}

	return head.Next
}

func mergeKLists(lists []*ListNode) *ListNode {
	if len(lists) == 0 {
		return nil
	}
	if len(lists) == 1 {
		return lists[0]
	}

	left := mergeKLists(lists[:len(lists)/2])
	right := mergeKLists(lists[len(lists)/2:])

	return mergeTwoLists(left, right)
}
