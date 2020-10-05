package main

type MovingAverage struct {
	nums   []int
	count  int
	curSum int
	size   int
}

/** Initialize your data structure here. */
func Constructor(size int) MovingAverage {
	nums := make([]int, 0)
	return MovingAverage{nums: nums, size: size}
}

func (this *MovingAverage) Next(val int) float64 {
	var res float64

	if this.size == 0 {
		return 0
	}

	this.count += 1
	this.nums = append(this.nums, val)
	if this.count <= this.size {
		this.curSum += val
		res = float64(this.curSum) / float64(this.count)

	} else {
		sum := 0
		for i := (this.count - this.size); i < this.count; i++ {
			sum += this.nums[i]

		}

		return float64(sum) / float64(this.size)
	}

	return res
}
