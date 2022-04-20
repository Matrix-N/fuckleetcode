package fuckleetcode

import (
	"math"
	"sort"
)

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

func isValid(s string) bool {
	if s == "" {
		return true
	}

	stack := []byte{}

	for i := 0; i < len(s); i++ {
		if s[i] == '(' || s[i] == '[' || s[i] == '{' {
			stack = append(stack, s[i])
		} else {
			if len(stack) == 0 {
				return false
			}
			if s[i] == ')' && stack[len(stack)-1] != '(' {
				return false
			}
			if s[i] == ']' && stack[len(stack)-1] != '[' {
				return false
			}
			if s[i] == '}' && stack[len(stack)-1] != '{' {
				return false
			}
			stack = stack[:len(stack)-1]
		}
	}

	return len(stack) == 0
}

func nextPermutation(nums []int) {
	if len(nums) == 0 {
		return
	}

	for i := len(nums) - 1; i > 0; i-- {
		if nums[i] > nums[i-1] {
			for j := len(nums) - 1; j > i; j-- {
				if nums[j] > nums[i-1] {
					nums[i-1], nums[j] = nums[j], nums[i-1]
					break
				}
			}
			for j := i; j < len(nums)/2; j++ {
				nums[j], nums[len(nums)-j+i-1] = nums[len(nums)-j+i-1], nums[j]
			}
			return
		}
	}

	for i := 0; i < len(nums)/2; i++ {
		nums[i], nums[len(nums)-i-1] = nums[len(nums)-i-1], nums[i]
	}
}

func hasCycle(head *ListNode) bool {
	if head == nil {
		return false
	}
	if head.Next == nil {
		return false
	}

	slow, fast := head, head.Next
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
		if slow == fast {
			return true
		}
	}

	return false
}

func inorderTraversal(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}

	res := []int{}
	stack := []*TreeNode{}
	cur := root
	for cur != nil || len(stack) > 0 {
		for cur != nil {
			stack = append(stack, cur)
			cur = cur.Left
		}
		cur = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		res = append(res, cur.Val)
		cur = cur.Right
	}

	return res
}

func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}

	return isSymmetricHelper(root.Left, root.Right)
}

func isSymmetricHelper(left, right *TreeNode) bool {
	if left == nil && right == nil {
		return true
	}
	if left == nil || right == nil {
		return false
	}
	if left.Val != right.Val {
		return false
	}

	return isSymmetricHelper(left.Left, right.Right) && isSymmetricHelper(left.Right, right.Left)
}

func levelOrder(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}

	res := [][]int{}
	queue := []*TreeNode{root}
	for len(queue) > 0 {
		level := []int{}
		n := len(queue)
		for i := 0; i < n; i++ {
			level = append(level, queue[i].Val)
			if queue[i].Left != nil {
				queue = append(queue, queue[i].Left)
			}
			if queue[i].Right != nil {
				queue = append(queue, queue[i].Right)
			}
		}
		queue = queue[n:]
		res = append(res, level)
	}

	return res
}

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	return max(maxDepth(root.Left), maxDepth(root.Right)) + 1
}

func buildTree(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 {
		return nil
	}
	if len(inorder) == 0 {
		return nil
	}

	root := &TreeNode{preorder[0], nil, nil}
	k := 0
	for i := 0; i < len(inorder); i++ {
		if inorder[i] == preorder[0] {
			k = i
			break
		}
	}

	root.Left = buildTree(preorder[1:k+1], inorder[:k])
	root.Right = buildTree(preorder[k+1:], inorder[k+1:])

	return root
}

func flatten(root *TreeNode) {
	if root == nil {
		return
	}

	var res []int
	preorder(root, &res)
	cur := root
	for i := 1; i < len(res); i++ {
		cur.Left = nil
		cur.Right = &TreeNode{res[i], nil, nil}
		cur = cur.Right
	}
}

func preorder(root *TreeNode, res *[]int) {
	if root == nil {
		return
	}

	*res = append(*res, root.Val)
	preorder(root.Left, res)
	preorder(root.Right, res)
}

func maxProfit(prices []int) int {
	if len(prices) == 0 {
		return 0
	}

	min, max := prices[0], 0
	for i := 1; i < len(prices); i++ {
		if prices[i]-min > max {
			max = prices[i] - min
		}
		if prices[i] < min {
			min = prices[i]
		}
	}

	return max
}

func maxPathSum(root *TreeNode) int {
	if root == nil {
		return 0
	}

	max := math.MinInt32
	maxPathSumHelper(root, &max)

	return max
}

func maxPathSumHelper(root *TreeNode, maxRes *int) int {
	if root == nil {
		return 0
	}

	left := maxPathSumHelper(root.Left, maxRes)
	right := maxPathSumHelper(root.Right, maxRes)

	curMax := max(root.Val, max(root.Val+left, root.Val+right))
	*maxRes = max(*maxRes, max(curMax, left+right+root.Val))

	return curMax
}

func singleNumber(nums []int) int {
	res := 0
	for _, num := range nums {
		res ^= num
	}

	return res
}

func climbStairs(n int) int {
	if n <= 2 {
		return n
	}
	a, b := 1, 2
	for i := 3; i <= n; i++ {
		a, b = b, a+b
	}

	return b
}

func rotate(matrix [][]int) {
	n := len(matrix)
	for i := 0; i < n/2; i++ {
		for j := i; j < n-i-1; j++ {
			matrix[i][j], matrix[j][n-i-1], matrix[n-i-1][n-j-1], matrix[n-j-1][i] = matrix[n-j-1][i], matrix[i][j], matrix[j][n-i-1], matrix[n-i-1][n-j-1]
		}
	}
}

func trap(height []int) int {
	if len(height) == 0 {
		return 0
	}

	left, right := 0, len(height)-1
	leftMax, rightMax := height[left], height[right]
	res := 0

	for left <= right {
		if leftMax < rightMax {
			if height[left] > leftMax {
				leftMax = height[left]
			} else {
				res += leftMax - height[left]
			}
			left++
		} else {
			if height[right] > rightMax {
				rightMax = height[right]
			} else {
				res += rightMax - height[right]
			}
			right--
		}
	}

	return res
}

func permute(nums []int) [][]int {
	if len(nums) == 0 {
		return nil
	}

	res := [][]int{}

	permuteHelper(nums, []int{}, &res)

	return res
}

func permuteHelper(nums []int, cur []int, res *[][]int) {
	if len(nums) == 0 {
		*res = append(*res, cur)
		return
	}

	for i := 0; i < len(nums); i++ {
		newCur := make([]int, len(cur))
		copy(newCur, cur)
		newCur = append(newCur, nums[i])
		newNums := make([]int, len(nums))
		copy(newNums, nums)
		newNums = append(newNums[:i], newNums[i+1:]...)
		permuteHelper(newNums, newCur, res)
	}
}
