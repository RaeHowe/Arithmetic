package main

func main()  {
	var array = []int{4,3,2,7,8,2,3,1}


}

func removeElement(nums []int, val int) int {
	var number = 0

	for _, item := range nums{
		if item != val{
			nums[number] = item

			number++
		}
	}

	return number
}