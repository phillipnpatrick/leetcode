package Solutions

// https://leetcode.com/problems/moving-average-from-data-stream/description/
// Given a stream of integers and a window size, calculate the moving average of all integers in the sliding window.
// Implement the MovingAverage class:
// 	MovingAverage(int size) Initializes the object with the size of the window size.
// 	double next(int val) Returns the moving average of the last size values of the stream.

// Example 1:
// Input
// ["MovingAverage", "next", "next", "next", "next"]
// [[3], [1], [10], [3], [5]]
// Output
// [null, 1.0, 5.5, 4.66667, 6.0]

// Explanation
// MovingAverage movingAverage = new MovingAverage(3);
// movingAverage.next(1); // return 1.0 = 1 / 1
// movingAverage.next(10); // return 5.5 = (1 + 10) / 2
// movingAverage.next(3); // return 4.66667 = (1 + 10 + 3) / 3
// movingAverage.next(5); // return 6.0 = (10 + 3 + 5) / 3

type MovingAverage struct {
    queue []int
	size int
	sum int
}


func MovingAverageConstructor(size int) MovingAverage {
    return MovingAverage{
		queue: []int{},
		size: size,
		sum: 0,
	}
}


func (this *MovingAverage) Next(val int) float64 {
    this.queue = append(this.queue, val)
	this.sum += val

	if len(this.queue) > this.size {
		this.sum -= this.queue[0]
		this.queue = this.queue[1:]
	}

	return float64(this.sum) / float64(len(this.queue))
}


/**
 * Your MovingAverage object will be instantiated and called as such:
 * obj := Constructor(size);
 * param_1 := obj.Next(val);
 */