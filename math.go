package go_utils

func Add(nums... int) int  {
	sum := 0
	for i := 0; i < len(nums); i++ {
		sum += nums[i]
	}
	return sum
}

func Times(nums... int) int  {
	sum := 1
	for i := 0; i < len(nums); i++ {
		sum *= nums[i]
	}
	return sum
}
