package main

import (
	"fmt"
	"math"
	"time"
)

func TopK(arr []int, k int) int {
	QSort(arr, len(arr)-k, 0, len(arr)-1)
	return arr[len(arr)-k]
}

func partition(arr []int, left, right int) int {
	pivot := left
	index := pivot + 1
	for i := left + 1; i <= right; i++ {
		if arr[i] < arr[pivot] {
			arr[index], arr[i] = arr[i], arr[index]
			index++
		}
	}
	arr[pivot], arr[index-1] = arr[index-1], arr[pivot]
	return index - 1
}

func QSort(arr []int, k, left, right int) {
	if left < right {
		index := partition(arr, left, right)
		if index == k {
			return
		} else if index < k {
			QSort(arr, k, index+1, right)
		} else {
			QSort(arr, k, left, index-1)
		}
		//QSort(arr, left, index-1)
		//QSort(arr, index+1, right)
	}
}

func quickSort(arr []int) {
	if len(arr) <= 1 {
		return
	}
	pivot := arr[0]
	left, right := 0, len(arr)-1
	for i := 1; i <= right; {
		if arr[i] < pivot {
			arr[left], arr[i] = arr[i], arr[left]
			left++
			i++
		} else {
			arr[right], arr[i] = arr[i], arr[right]
			right--
		}
	}
	quickSort(arr[0:left])
	quickSort(arr[left+1:])
}

func selectionSort(arr []int) {
	for i := 0; i < len(arr); i++ {
		index := i
		for j := i + 1; j < len(arr); j++ {
			if arr[index] > arr[j] {
				index = j
			}
		}
		if index != i {
			arr[i], arr[index] = arr[index], arr[i]
		}
	}
}

func bubbleSort(arr []int) {
	for i := len(arr) - 1; i > 0; i-- {
		var flag = true
		for j := 0; j < i; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
				flag = false
			}
		}
		if flag {
			return
		}
	}
}

func insertionSort(arr []int) {
	for i := 1; i < len(arr); i++ {
		for j := i - 1; j >= 0; j-- {
			if arr[j+1] < arr[j] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
}

func merge(left, right []int) []int {
	lenL, lenR := len(left), len(right)
	indexL, indexR := 0, 0
	res := make([]int, 0)
	for indexL < lenL && indexR < lenR {
		if left[indexL] < right[indexR] {
			res = append(res, left[indexL])
			indexL++
		} else {
			res = append(res, right[indexR])
			indexR++
		}
	}
	if indexL < lenL {
		res = append(res, left[indexL:]...)
	}
	if indexR < lenR {
		res = append(res, right[indexR:]...)
	}
	return res
}

func mergeSort(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}
	m := len(arr) / 2
	left := mergeSort(arr[:m])
	right := mergeSort(arr[m:])
	return merge(left, right)
}

func reverse(s string) {
	var arr = []rune(s)
	for i, j := 0, len(arr)-1; i < j; i, j = i+1, j-1 {
		arr[i], arr[j] = arr[j], arr[i]
	}
	fmt.Println(string(arr))
}

func maxSubseqSum(arr []int) {
	maxSum := 0
	curSum := 0
	var sub []int
	tmp := make([]int, 0)
	for i := 0; i < len(arr); i++ {
		curSum += arr[i]
		tmp = append(tmp, arr[i])
		if curSum > maxSum {
			maxSum = curSum
			sub = make([]int, len(tmp))
			copy(sub, tmp)
		} else if curSum < 0 {
			curSum = 0
			tmp = make([]int, 0)
		}
	}
	fmt.Println(maxSum)
	fmt.Println(sub)
}

func minSteps(arr [][]int) int {
	dp := make([][]int, len(arr))
	for i := 0; i < len(arr); i++ {
		dp[i] = make([]int, len(arr[0]))
	}
	dp[0][0] = arr[0][0]
	for i := 1; i < len(arr); i++ {
		dp[0][i] = dp[0][i-1] + arr[0][i]
	}
	for i := 1; i < len(arr[0]); i++ {
		dp[i][0] = dp[i-1][0] + arr[i][0]
	}
	for i := 1; i < len(arr); i++ {
		for j := 1; j < len(arr[0]); j++ {
			dp[i][j] = int(math.Min(float64(dp[i-1][j]), float64(dp[i][j-1]))) + arr[i][j]
		}
	}
	return dp[len(arr)-1][len(arr[0])-1]
}

var res [][]string

func solution(arr []string) [][]string {
	item := make([]string, 0)
	index := make([]bool, len(arr))
	back(arr, item, index)
	return res
}

func back(arr []string, item []string, index []bool) {
	//fmt.Println(res)
	//fmt.Println(item)
	//fmt.Println("-----")
	if len(item) == len(arr) {
		tmp := make([]string, len(item))
		copy(tmp, item)
		res = append(res, tmp)
		return
	} else {
		for i := 0; i < len(arr); i++ {
			if index[i] {
				continue
			}
			item = append(item, arr[i])
			index[i] = true
			back(arr, item, index)
			index[i] = false
			item = item[0 : len(item)-1]
		}
	}
}

func maxSub(arr []int) int {
	dp := make([]int, len(arr))
	for i := 0; i < len(arr); i++ {
		dp[i] = 1
		for j := 0; j < i; j++ {
			if arr[j] < arr[i] {
				dp[i] = int(math.Max(float64(dp[i]), float64(dp[j]+1)))
			}
		}
	}
	return dp[len(arr)-1]
}

func maxWarter(arr []int) int {
	if len(arr) <= 1 {
		return 0
	}
	left, right := 0, len(arr)-1
	res := 0
	for left < right {
		height := min(arr[left], arr[right])
		res = max(res, height*(right-left))
		if arr[left] < arr[right] {
			left++
		} else {
			right++
		}
	}
	return res
}

func min(x, y int) int {
	if x-y > 0 {
		return y
	}
	return x
}

func max(x, y int) int {
	if x-y > 0 {
		return x
	}
	return y
}

type A struct {
	name  string
	score int
}

func (a A) nameit() string {
	return a.name
}

type B struct {
	name string
	A
}

func (b B) nameit() string {
	return b.name
}

type C struct {
	A
	B
}

func let(args ...string) {
	fmt.Println(len(args))
	fmt.Println(args)
}

func deepPrice() (uint32, uint32) {
	deepVal := float64(0) / 10000
	fmt.Printf("%.6f", deepVal)
	newPrice := uint32(math.Floor((0.0296 / deepVal) * float64(5000)))
	fmt.Printf("newPrice:%d\n", newPrice)
	if newPrice < 4000 {
		newPrice = 1
	}
	if newPrice > 6000 {
		newPrice = 6000
	}
	newEcpm := uint32(math.Floor(1000 * float64(newPrice) * 0.052))
	return newPrice, newEcpm
}

func div(x float64, y int) uint32 {
	z := float64(y) / 100
	return uint32(x / z)
}

type CC struct {
	running chan struct{}
}

func NewC() *CC {
	c := new(CC)
	c.running = make(chan struct{}, 1)
	return c
}

func (c *CC) Init() error {
	select {
	case c.running <- struct{}{}:
	default:
		fmt.Println("object is busy...")
		return nil
	}
	defer func() {
		<-c.running
	}()
	fmt.Println("do something in init...")
	time.Sleep(100 * time.Millisecond)
	return nil
}

func main() {
	//a, b := deepPrice()
	//fmt.Println(a, b)
	//var a float64 = 3.1213141516
	//aa := fmt.Sprintf("%.8f", a)
	//fmt.Println(aa)
	//var flags int
	//writing := 4
	//flags ^= writing
	//fmt.Println(flags)
	//flags ^= writing
	//fmt.Println(flags)
	//fmt.Println(flags & writing)
	var m map[string]string
	fmt.Println(m == nil)
	fmt.Println(len(m) / (len(m) + 1))
	fmt.Println(28 % 29)
	fmt.Println(2 & 1)
}
