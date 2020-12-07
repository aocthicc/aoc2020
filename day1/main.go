package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func main() {
	data, err := ioutil.ReadFile("input")
	if nil != err {
		panic(err)
	}
	// yes I know it's slow trash algorithms,
	// I don't know big O theory and that shit
	nums := convertToInt(strings.Split(string(data), "\r\n"))
	fmt.Println(findSum(nums))
	fmt.Println(findSumOfThree(nums))
}

func findSum(nums []int) (int, int, int) {
	for _, a := range nums {
		for _, b := range nums {
			if a+b == 2020 {
				return a,b, a*b
			}
		}
	}
	return 0,0,0
}

func findSumOfThree(nums []int) (int, int, int, int) {
	for _, a := range nums {
		for _, b := range nums {
			for _, c := range nums {
				if a+b+c == 2020 {
					return a,b,c, a*b*c
				}
			}
		}
	}
	return 0,0,0,0
}

func convertToInt(nums []string) []int {
	var numsInt []int
	for _, num := range nums {
		n, err := strconv.Atoi(num)
		if nil != err {
			continue
		}
		numsInt = append(numsInt, n)
	}
	return numsInt
}

