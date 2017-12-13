package main

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
	pbms.LongestCommonPrefix()
	pbms.ThreeSum()
	pbms.FindLadders()
	pbms.ThreeSumClosest()
	pbms.FourSum()
	pbms.RemoveNthFromEnd()
	pbms.IsValid()
	pbms.MergeTwoLists()
	pbms.SumNumbers()
	pbms.Solve()
	pbms.Partition()
	pbms.MergeKLists()
	pbms.SwapPairs()
	pbms.ReverseKGroup()
	pbms.RemoveDuplicates()
	pbms.RemoveElement()
	pbms.StrStr()
	pbms.Divide()
	pbms.NextPermutation()
	pbms.LongestValidParentheses()
	pbms.Search()
	pbms.SearchRange()
	pbms.SearchInsert()
	pbms.IsValidSudoku()
	pbms.CountAndSay()
	pbms.FirstMissingPositive()
	pbms.Trap()
	pbms.Multiply()
	pbms.Rotate()
	pbms.Jump()
	pbms.GroupAnagrams()
	pbms.MyPow()
	pbms.MaxSubArray()
	pbms.SpiralOrder()
	pbms.LengthOfLastWord()
	pbms.CanJump()
	pbms.Merge()
	pbms.Insert()
	pbms.GenerateMatrix()
	pbms.RotateRight()
	pbms.UniquePaths()
	pbms.UniquePathsWithObstacles()
	pbms.MinPathSum()
	pbms.IsNumber()
	pbms.PlusOne()
	pbms.AddBinary()
	pbms.MySqrt()
	pbms.ClimbStairs()
	pbms.SimplifyPath()
	pbms.SearchMatrix()
	pbms.SortColors()
	pbms.MinWindow()
	pbms.SetZeroes()
	pbms.MinDistance()
	pbms.RemoveDuplicates2()
	pbms.Search2()
	pbms.DeleteDuplicates()
	pbms.DeleteDuplicates2()
	pbms.LargestRectangleArea()
	pbms.MergeSortedArray()
	pbms.MaximalRectangle()
	pbms.PartitionList()
	pbms.NumDecodings()
	pbms.ReverseBetween()
	pbms.InorderTraversal()
	pbms.NumTrees()
	pbms.IsInterleave()
	pbms.GenerateTrees()
	pbms.NumDistinct()
	pbms.Generate()
	pbms.GetRow()
	pbms.MinimumTotal()
	pbms.MaxProfit()
	pbms.MaxProfit2()
	pbms.IsPalindrome()
	pbms.MaxProfit3()
	pbms.LadderLength()
	pbms.LongestConsecutive()
	pbms.MinCut()
	pbms.CanCompleteCircuit()
	pbms.Candy()
	pbms.SingleNumber()
	pbms.SingleNumber2()
	pbms.WordBreak()
	pbms.WordBreak2()
	pbms.ReorderList()
	pbms.PreorderTraversal()
	pbms.PostorderTraversal()
	pbms.LRUCacheSample()
	pbms.InsertionSortList()
	pbms.SortList()
	pbms.MaxPoints()
	pbms.EvalRPN()
	pbms.MaxProduct()
	pbms.FindMin()
	pbms.FindMinDuplicated()
	pbms.MinStackSample()
	pbms.FindPeakElement()
	pbms.CompareVersion()
	pbms.MaximumGap()
	pbms.FractionToDecimal()
	pbms.TwoSumSorted()
	pbms.ConvertToTitle()
	pbms.MajorityElement()
	pbms.TitleToNumber()
	pbms.TrailingZeroes()
	pbms.CalculateMinimumHP()
	pbms.LargestNumber()
}

