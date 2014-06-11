package main

import (
	"fmt"
)

/**
 * 计算平均值
 */
func CalculateAvg(list []float64) (avgValue float64) {
	//处理位数
	defer func() {
		avgStr := strconv.FormatFloat(avgValue, 'f', 6, 64)
		avgValue, _ = strconv.ParseFloat(avgStr, 64)
	}()

	if len(list) == 2 { //2个直接计算
		return (list[0] + list[1]) / 2
	}
	listSort(list)
	fmt.Println(list)
	n := len(list)

	//计算中间60%的区件值
	diff := list[n-1] - list[0]
	begin := list[0] + diff*0.2
	end := list[n-1] - diff*0.2
	// fmt.Println(diff * 0.2)
	// fmt.Println(begin)
	// fmt.Println(end)

	var pIndex, hIndex int
	for i := 1; i < n; i++ {
		if list[i] >= begin {
			pIndex = i
			break
		}
	}
	for i := n - 2; i > 0; i-- {
		if list[i] <= end {
			hIndex = i
			break
		}
	}
	if hIndex-pIndex >= 0 { //区间内有值
		if hIndex-pIndex == 0 { //直接返回
			avgValue = list[hIndex]
			return
		}

		temp := list[pIndex : hIndex+1]
		var sum float64
		for i := 0; i < len(temp); i++ {
			sum += temp[i]
		}
		avgValue = sum / float64(len(temp))
		return
	}

	//若该区间内无值，则去首位计算平均值
	temp := list[1 : n-1]
	var sum float64
	for i := 0; i < len(temp); i++ {
		sum += temp[i]
	}
	avgValue = sum / float64(len(temp))
	return
}

/**
 * 数组排序，暂用冒泡
 */
func listSort(a []float64) {
	for i := 0; i < len(a); i++ {
		for j := i + 1; j < len(a); j++ {
			if a[i] > a[j] {
				a[i], a[j] = a[j], a[i]
			}
		}
	}
}

func main() {
	a := []float64{116.480604, 116.481032, 116.480623, 116.480581, 116.481066, 116.481131, 116.480604}
	fmt.Println(CalculateAvg(a))
}
