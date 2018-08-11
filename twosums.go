package main

import ("log")

func main() {
	log.Println(twoSum([]int{1,2,3,4,8}, 7))
}

func twoSum(nums []int, target int) []int {
    for i:= 0; i < len(nums); i++ {
        for j := len(nums) - 1; j > i; j-- {
            if nums[i] + nums[j] == target {
                return []int{i, j}
            }
        }
    }
    
    return []int{0,0} 
}