package main

import (
	"fmt"
	"math"
	"time"
)

var timesknown1 int = 0
var timesrec int = 0

var recTimes bool

var timesMap1 []int
var timesMap2 []int
var timesMap3 []int
var squrt5 float64

func init() {
	squrt5 = math.Pow(5, 0.5)
}

func main() {

	testCorrect()
}

// 簡單驗證函數是否正確
func testCorrect() {
	for i := int64(0); i < 67; i++ {
		fmt.Println(Fibonacci_Fast_Doubling(i))
		fmt.Println(Fibonacci_BinetFormula(i + 1))
	}
}

// recTimes設置為true，想要看看n呼叫多少次
func testCallTimeForFibonacci_Recursion() {
	recTimes = true
	testvalue := int64(40)

	timesMap1 = make([]int, testvalue+1)
	timesMap2 = make([]int, testvalue+1)
	timesMap3 = make([]int, testvalue+1)

	fmt.Println(Fibonacci_RecursionV2(testvalue))
	fmt.Println(Fibonacci_RecursionV3(testvalue))
	fmt.Println(Fibonacci_Recursion(testvalue))

	fmt.Println(timesMap1)
	fmt.Println(timesMap2)
	fmt.Println(timesMap3)
}

// 驗證迴圈版本效能
func FibonacciEfftion_Fast() {
	// meth1()
	// meth2()
	times := 100
	testvalue := int64(10000000)
	var totaltime time.Duration
	var avgt time.Duration

	totaltime = 0
	for i := 0; i < times; i++ {
		begin := time.Now()
		Fibonacci_Loop2(testvalue)
		totaltime += time.Now().Sub(begin)
	}
	avgt = totaltime / time.Duration(times)
	fmt.Println(avgt.Microseconds())

	// totaltime = 0
	// for i := 0; i < times; i++ {
	// 	begin := time.Now()
	// 	Fibonacci_Loop(testvalue)
	// 	totaltime += time.Now().Sub(begin)
	// }
	// avgt = totaltime / time.Duration(times)
	// fmt.Println(avgt)

	// //遞迴第三版夠快，所以也來這邊測試
	// totaltime = 0
	// for i := 0; i < times; i++ {
	// 	begin := time.Now()
	// 	Fibonacci_RecursionV3(testvalue)
	// 	totaltime += time.Now().Sub(begin)
	// }
	// avgt = totaltime / time.Duration(times)
	// fmt.Println(avgt)

	totaltime = 0
	for i := 0; i < times; i++ {
		begin := time.Now()
		Fibonacci_Fast_Doubling(testvalue)
		totaltime += time.Now().Sub(begin)
	}
	avgt = totaltime / time.Duration(times)
	fmt.Println(avgt.Microseconds())

	totaltime = 0
	for i := 0; i < times; i++ {
		begin := time.Now()
		Fibonacci_BinetFormula(testvalue)
		totaltime += time.Now().Sub(begin)
	}
	avgt = totaltime / time.Duration(times)
	fmt.Println(avgt.Microseconds())

}

func FibonacciEfftionRecursionSlow() {
	times := 20
	testvalue := int64(100)
	var totaltime time.Duration
	var avgt time.Duration

	totaltime = 0
	for i := 0; i < times; i++ {
		begin := time.Now()
		Fibonacci_RecursionV3(testvalue)
		totaltime += time.Now().Sub(begin)
	}
	avgt = totaltime / time.Duration(times)
	fmt.Println(avgt)

	totaltime = 0
	for i := 0; i < times; i++ {
		begin := time.Now()
		Fibonacci_RecursionV2(testvalue)
		totaltime += time.Now().Sub(begin)
	}
	avgt = totaltime / time.Duration(times)
	fmt.Println(avgt)

	// totaltime = 0
	// for i := 0; i < times; i++ {
	// 	begin := time.Now()
	// 	Fibonacci_Recursion(testvalue)
	// 	totaltime += time.Now().Sub(begin)
	// }
	// avgt = totaltime / time.Duration(times)
	// fmt.Println(avgt)

}

// 費式數列迴圈解法1，效能比2差，推測是因為要額外計算奇數偶數
func Fibonacci_Loop(n int64) int64 {
	n0 := int64(1)
	n1 := int64(1)
	for i := int64(0); i < n; i++ {
		if i%2 == 0 {
			n0 += n1
		} else {
			n1 += n0
		}
	}
	if n%2 == 0 {
		return n0
	} else {
		return n1
	}
}

// 費式數列迴圈解法2，直接記錄三個數值最後回傳
func Fibonacci_Loop2(n int64) int64 {
	if n < 3 {
		if n == 2 {
			return 2
		} else if n == 1 {
			return 1
		} else if n == 0 {
			return 1
		}
	}
	prevPrev := int64(1)
	prev := int64(1)
	current := int64(2)
	for i := int64(3); i < n; i++ {
		current = prev + prevPrev
		prevPrev = prev
		prev = current
	}
	return current
}

// 最基本款的遞迴費式數列，非常慢
func Fibonacci_Recursion(n int64) int64 {
	if recTimes {
		timesMap1[n]++
	}
	if n < 2 {
		return 1
	}
	return Fibonacci_Recursion(n-1) + Fibonacci_Recursion(n-2)
}

// 遞迴費式數列，但先計算n-2，再把n-2塞給n-1，一定程度減少重複計算
func Fibonacci_RecursionV2(n int64) int64 {
	r := _Fibonacci_RecursionV2_1(n)
	return r
}
func _Fibonacci_RecursionV2_1(n int64) int64 {
	if recTimes {
		timesMap2[n]++
	}
	if n < 2 {
		return 1
	}
	v2 := _Fibonacci_RecursionV2_1(n - 2)
	v1 := _Fibonacci_RecursionV2_2(n-1, v2)
	return v2 + v1
}
func _Fibonacci_RecursionV2_2(n int64, v1 int64) int64 {
	if n < 2 {
		return 1
	}
	v2 := _Fibonacci_RecursionV2_1(n - 2)
	return v1 + v2
}

// 遞迴費式數列，子函數會回傳 n 與 n-1的答案
// 先計算n-2，同時回傳n-3，這樣可以算出n-1，這樣就沒有重複計算了，效能跟迴圈接近(但疑似因為函數不斷呼叫，效能還是比普通LOOP慢6倍左右，不過不太受最大數值影響)
func Fibonacci_RecursionV3(n int64) int64 {
	r, _ := _Fibonacci_RecursionV3_1(n)
	return r
}
func _Fibonacci_RecursionV3_1(n int64) (int64, int64) {
	if recTimes {
		timesMap3[n]++
	}
	if n < 3 {
		if n == 2 {
			return 2, 1
		} else {
			return 1, 1
		}
	}
	v2, v3 := _Fibonacci_RecursionV3_1(n - 2)
	v1 := v2 + v3
	return v2 + v1, v1
}

// 快速解法
func Fibonacci_Fast_Doubling(n int64) int64 {
	a, _ := _Fibonacci_Fast_Doubling(n)
	return a
}
func _Fibonacci_Fast_Doubling(n int64) (int64, int64) {
	//公式:
	//F(2n+1)=F(n+1)^2+F(n)^2
	//F(2n)=F(n)(2*F(n+1)-F(n))

	//實現:
	//假設起始數字為n0
	//n1為 n0/2 的商
	//n1假設為單數 就要回傳 n1+1 ,n1
	//n1假如為雙數 就回傳n1 ,n1-1
	//
	//表現就是希望上一階層回傳 n+1 和 n，並且n+1為奇數往下才不會產生更多需要計算的項目
	//如果n+1並非奇數，可以轉換成在計算 n n-1，並回傳 2n+n-1當作n+1

	//假設我要計算
	//777 776
	//可以拆解成 388 *2 +1 和 388*2
	//需要計算 388 和 388+1的數值
	//389 388 則需要計算 194 數值
	//195 194 計算 97 98
	//97 98->97 96的數值
	//97 96 計算 48
	//49 48 計算 24
	//25 24 計算 12
	//13 12 計算 6
	//7 6 計算3
	//3 4 計算 3 2
	//3 2計算1
	//1 2計算1 0
	//0 回傳 1 ,0

	if n == 0 {
		return 1, 0
	}

	q := n / 2
	r := n % 2

	f2, f1 := _Fibonacci_Fast_Doubling(q)

	f22 := f1*f1 + f2*f2
	f11 := f1 * (2*f2 - f1)
	if r == 0 {
		return f22, f11
	} else {
		return f22 + f11, f22
	}

}

// 函數算法(使用浮點數)，最高到74(網路說71)
func Fibonacci_BinetFormula(n int64) int64 {
	base := float64(n)
	p1 := math.Pow((1+squrt5)/2, base)
	p2 := math.Pow((1-squrt5)/2, base)
	rtn := math.Round((p1 - p2) / squrt5)
	return int64(rtn)
}
