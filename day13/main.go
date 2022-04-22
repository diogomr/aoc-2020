package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

func main() {
	partOne()
	partTwo()
}

func partOne() {
	lines := strings.Split(input, "\n")
	minDep, _ := strconv.Atoi(lines[0])
	buses := strings.Split(lines[1], ",")

	var ids []int
	for _, bus := range buses {
		if bus != "x" {
			id, _ := strconv.Atoi(bus)
			ids = append(ids, id)
		}
	}

	bus := -1
	min := math.MaxInt
	for _, id := range ids {
		res := 0
		for res < minDep {
			res += id
		}
		if res < min {
			min = res
			bus = id
		}
	}

	fmt.Println(bus * (min - minDep))
}

func partTwo() {
	lines := strings.Split(input, "\n")
	buses := strings.Split(lines[1], ",")

	var busDep [][2]int
	for i, b := range buses {
		if b != "x" {
			bInt, _ := strconv.Atoi(b)
			busDep = append(busDep, [2]int{bInt, i})
		}
	}

	step := 1
	var res int
	for _, b := range busDep {
		val, pos := b[0], b[1]
		for (res+pos)%val != 0 {
			res += step
		}
		step *= val
	}
	fmt.Println(res)
}

const input = `1000417
23,x,x,x,x,x,x,x,x,x,x,x,x,41,x,x,x,37,x,x,x,x,x,479,x,x,x,x,x,x,x,x,x,x,x,x,13,x,x,x,17,x,x,x,x,x,x,x,x,x,x,x,29,x,373,x,x,x,x,x,x,x,x,x,x,x,x,x,x,x,x,x,x,19`
