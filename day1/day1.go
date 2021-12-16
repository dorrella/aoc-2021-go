package main

/*
You're minding your own business on a ship at sea when the overboard alarm goes off!
You rush to see if you can help. Apparently, one of the Elves tripped and accidentally
sent the sleigh keys flying into the ocean!

Before you know it, you're inside a submarine the Elves keep ready for situations like
this. It's covered in Christmas lights (because of course it is), and it even has an
experimental antenna that should be able to track the keys if you can boost its signal
strength high enough; there's a little meter that indicates the antenna's signal strength
by displaying 0-50 stars.

Your instincts tell you that in order to save Christmas, you'll need to get all fifty stars
by December 25th.

Collect stars by solving puzzles. Two puzzles will be made available on each day in the Advent
calendar; the second puzzle is unlocked when you complete the first. Each puzzle grants one
star. Good luck!

As the submarine drops below the surface of the ocean, it automatically performs a sonar sweep
of the nearby sea floor. On a small screen, the sonar sweep report (your puzzle input) appears:
each line is a measurement of the sea floor depth as the sweep looks further and further away
from the submarine.

For example, suppose you had the following report:

199
200
208
210
200
207
240
269
260
263

This report indicates that, scanning outward from the submarine, the sonar sweep found depths
of 199, 200, 208, 210, and so on.

The first order of business is to figure out how quickly the depth increases, just so you know
what you're dealing with - you never know if the keys will get carried into deeper water by an
ocean current or a fish or something.

To do this, count the number of times a depth measurement increases from the previous measurement.
(There is no measurement before the first measurement.) In the example above, the changes are as follows:

199 (N/A - no previous measurement)
200 (increased)
208 (increased)
210 (increased)
200 (decreased)
207 (increased)
240 (increased)
269 (increased)
260 (decreased)
263 (increased)

In this example, there are 7 measurements that are larger than the previous measurement.

How many measurements are larger than the previous measurement?
*/

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
)

func ReadInput(c chan<- int, r io.Reader) {
	scanner := bufio.NewScanner(r)
	scanner.Split(bufio.ScanWords)
	for scanner.Scan() {
		x, err := strconv.Atoi(scanner.Text())
		if err != nil {
			panic(err)
		}
		c <- x
	}
	close(c)
}

func ProcessInput(c <-chan int) int {
	var last int = -1
	var count int = 0

	for depth := range c {
		if last < 0 {
			last = depth
			fmt.Println(depth)
			continue
		}

		if depth > last {
			count++
			fmt.Printf("%d (increased)\n", depth)
		} else {
			fmt.Printf("%d (decreased)\n", depth)
		}

		last = depth
	}

	return count

}

// using last 3 as sliding window
func ProcessInput2(c <-chan int) int {
	var count int = 0
	var last_sum int = 0
	var sum int = 0
	last := [3]int{0, 0, 0}

	for depth := range c {

		last[0] = last[1]
		last[1] = last[2]
		last[2] = depth

		//should only happen on first few
		if last[0] == 0 {
			continue
		}

		last_sum = sum
		sum = last[0] + last[1] + last[2]

		//skip first rolling widown
		if last_sum == 0 {
			fmt.Printf("%d (NA)\n", sum)
			continue
		}

		if sum > last_sum {

			count++
			fmt.Printf("%d (increased)\n", sum)
		} else {
			//counts no change as decrease, but should be fine for this
			fmt.Printf("%d (decreased)\n", sum)
		}

	}

	return count

}

// Scans file and counts increases
// uses rolling window when rw is true
func ScanFile(file string, rw bool) {
	f, err := os.Open(file)
	if err != nil {
		panic(err)
	}

	c := make(chan int)
	go ReadInput(c, f)

	var count int
	if rw {
		count = ProcessInput2(c)
	} else {
		count = ProcessInput(c)
	}

	fmt.Printf("file \"%s\" increased: %d times\n", file, count)
}

func main() {
	//no rolling window
	ScanFile("sample.txt", false)
	fmt.Println()
	ScanFile("input.txt", false)
	fmt.Println()

	// rolling window
	ScanFile("sample.txt", true)
	fmt.Println()
	ScanFile("input.txt", true)
}
