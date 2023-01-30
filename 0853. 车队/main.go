package main

import (
	"fmt"
	"sort"
)

// https://leetcode.cn/problems/car-fleet/

/*
分析：
	既然后面的不能超过前面的，到终点赶上都算一个车队，那么就看第一个车到终点的时候，后面的车是否能跟上了，跟上了就是一个车队
*/

type Car struct {
	position []int
	speed    []int
}

func (c Car) Len() int {
	return len(c.position)
}

func (c Car) Less(i, j int) bool {
	return c.position[i] < c.position[j]
}

func (c Car) Swap(i, j int) {
	c.position[i], c.position[j] = c.position[j], c.position[i]
	c.speed[i], c.speed[j] = c.speed[j], c.speed[i]
}

func carFleet(target int, position []int, speed []int) int {
	// 首先要根据position给车都排个序，当然，他的speed也得跟着动
	car := Car{
		position: position,
		speed:    speed,
	}
	sort.Sort(car)

	fmt.Println(car.position)
	fmt.Println(car.speed)

	count, spend := 0, -1.0
	for i := car.Len() - 1; i >= 0; i-- {
		spendTmp := float64(target-car.position[i]) / float64(car.speed[i])
		if spendTmp > spend {
			count++
			spend = spendTmp
		}
	}

	return count
}

func main() {
	fmt.Println(carFleet(12, []int{10, 8, 0, 5, 3}, []int{2, 4, 1, 1, 3}))
	fmt.Println(carFleet(100, []int{0, 2, 4}, []int{4, 2, 1}))
}
