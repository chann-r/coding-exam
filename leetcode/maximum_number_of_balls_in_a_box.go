package leetcode

/*
	You are working in a ball factory where you have n balls numbered from lowLimit up to highLimit inclusive (i.e., n == highLimit - lowLimit + 1), and an infinite number of boxes numbered from 1 to infinity.

	Your job at this factory is to put each ball in the box with a number equal to the sum of digits of the ball's number. For example, the ball number 321 will be put in the box number 3 + 2 + 1 = 6 and the ball number 10 will be put in the box number 1 + 0 = 1.

	Given two integers lowLimit and highLimit, return the number of balls in the box with the most balls.
 */

func countBalls(lowLimit int, highLimit int) int {
	box := make(map[int]int, highLimit-lowLimit+1)
	var max int
	for limit:=lowLimit; limit<=highLimit; limit++ {
		var boxNum int

		num := limit
		for num > 0 {
			boxNum += num % 10
			num = num / 10
		}

		//for limit > 0 {
		//	boxNum += limit % 10
		//	limit = limit / 10
		//}

		box[boxNum]++

		if max < box[boxNum] {
			max = box[boxNum]
		}
	}
	return max
}