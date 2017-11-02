package main

import "fmt"
import "os"
import "log"
import "runtime/pprof"

import pbms "./problems"

func main() {
	f, err := os.Create("cpu.prof")

	if err != nil {
		log.Printf("Failed to create cpu.prof\n")
	} else {
		pprof.StartCPUProfile(f)
		defer pprof.StopCPUProfile()
	}

	fmt.Println("Hello, Leetcode!")

	pbms.LetterCombination()
	pbms.GenerateParenthesis()
	pbms.CombinationSum()
	pbms.CombinationSum2()
	pbms.Permute()
	pbms.PermuteUnique()
	pbms.GetPermutation()
	pbms.Combine()
	pbms.Subsets()
	pbms.Exist()
	pbms.GrayCode()
	pbms.TwoSum()
	pbms.AddTwoNumbers()
	pbms.SubsetsWithDup()
	pbms.LengthOfLongestSubstring()
	pbms.FindMedianSortedArrays()
	pbms.RestoreIpAddresses()
	pbms.IsValidBST()
	pbms.RecoverTree()
	pbms.IsSameTree()
	pbms.IsSymmetric()
	pbms.LevelOrder()
	pbms.ZigzagLevelOrder()
	pbms.MaxDepth()
	pbms.BuildTree()
	pbms.BuildTree2()
	pbms.LevelOrderBottom()
	pbms.SortedArrayToBST()
	pbms.SortedListToBST()
	pbms.IsBalanced()
	pbms.MinDepth()
	pbms.HasPathSum()
	pbms.PathSum()
	pbms.LongestPalindrome()
	pbms.Convert()
	pbms.Reverse()
	pbms.MyAtoi()
	pbms.Flatten()
	pbms.MaxPathSum()
}

