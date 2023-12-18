package main

import (
	"fmt"
	"math/rand"
	"runtime"
	"time"
)

func main() {
	value()

	checkNil()
	//
	//arrAndSlice()
	//
	//copySlice()
	//
	//testLastChars(lastNumsBySlice)
	//testLastChars(lastNumsBySlice)
}

//1. 切片下标取值，下标前闭后开
func value() {
	var arr = [...]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	var slice0 = arr[2:8]
	var slice1 = arr[0:6]         //可以简写为 var g_slice []int = arr[:end]
	var slice2 = arr[5:10]        //可以简写为 var g_slice[]int = arr[start:]
	var slice3 = arr[0:]          //var g_slice []int = arr[:]
	var slice4 = arr[:len(arr)-1] //去掉切片的最后一个元素
	fmt.Printf("全局变量：arr %v\n", arr)
	fmt.Printf("全局变量：slice0 %v\n", slice0)
	fmt.Printf("全局变量：slice1 %v\n", slice1)
	fmt.Printf("全局变量：slice2 %v\n", slice2)
	fmt.Printf("全局变量：slice3 %v\n", slice3)
	fmt.Printf("全局变量：slice4 %v\n", slice4)
	fmt.Printf("-----------------------------------\n")
	return
}

//2. 未初始化的slice为nil
func checkNil() {

	var s1 []int
	s2 := []int{}
	//make第二个参数是长度，缺省第三个参数时，make（类型，长度）
	//容量是隐形的，只有扩容才会用到，如果长度大于0，则初始化为对应个数的0值
	s3 := make([]int, 0, 8)
	s4 := make([]int, 8, 8)
	s5 := make([]int, 8)
	s6 := make([][]int, 8)
	fmt.Println(s1 == nil, s2 == nil, s3 == nil)
	fmt.Println(s1, s2, s3)
	fmt.Println(cap(s1), cap(s2), cap(s3))
	s3 = append(s3, 1)
	s4 = append(s4, 1)
	fmt.Println("s3:make([]int, 0, 8)", s3)
	fmt.Println("s4:make([]int, 8, 8)", s4)
	fmt.Println("最常用：s5:make([]int, 8)", s5)
	fmt.Println("s6:make([][]int, 8)", s6)
	fmt.Printf("-----------------------------------\n")
	return
}

func arrAndSlice() {
	// 读写操作实际目标是底层数组，只需注意索引号的差别。
	data := [...]int{0, 1, 2, 3, 4, 5}
	s := data[2:4]
	s[0] += 100
	s[1] += 200
	fmt.Println(s)
	fmt.Println(data) //底层数组data也会改变

	c := data[:0]
	c = append(c, 0) //长度为0的切片添加元素时，也会影响底层数组
	fmt.Println(c[0], data)

	// 使用 make 动态创建slice，避免了数组必须用常量做长度的麻烦。还可用指针直接访问底层数组，退化成普通数组操作。
	ss := []int{0, 1, 2, 3}
	p := &ss[2] // *int, 获取底层数组元素指针。
	*p += 100
	fmt.Println(ss)

	//向 g_slice 尾部添加数据，返回新的 g_slice 对象。
	s1 := make([]int, 510, 510)
	fmt.Printf("%p\n", &s1)
	s2 := append(s1, 1)
	fmt.Printf("%p\n", &s2)
	fmt.Println(s1, s2)
	fmt.Println(cap(s1), len(s1)) //扩容时，小数组（len256以下）翻倍，大数组1.25倍
	fmt.Println(cap(s2), len(s2))

	//切片resize
	var a = []int{0, 1, 2, 3, 4, 5, 6}
	fmt.Printf("g_slice a : %v , len(a) : %v, cap(a) : %v\n", a, len(a), cap(a))
	b := a[1:4]
	fmt.Printf("g_slice b : %v , len(b) : %v, cap(b) : %v\n", b, len(b), cap(b))
	d := b[0:5] //截取时考虑底层数组是否越界，如果通过索引直接取值则考虑切片是否越界
	fmt.Printf("g_slice c : %v , len(c) : %v, cap(c) : %v\n", d, len(d), cap(d))
	return
}

func copySlice() {
	//切片拷贝，copy ：函数 copy 在两个 g_slice 间复制数据，复制长度以 len 小的为准。两个 g_slice 可指向同一底层数组，允许元素区间重叠。
	s1 := []int{1, 2, 3, 4, 5}
	fmt.Printf("g_slice s1 : %v\n", s1)
	s2 := make([]int, 10)
	fmt.Printf("g_slice s2 : %v\n", s2)
	copy(s2, s1)
	fmt.Printf("copied g_slice s1 : %v\n", s1)
	fmt.Printf("copied g_slice s2 : %v\n", s2)
	s3 := []int{1, 2, 3}
	fmt.Printf("g_slice s3 : %v\n", s3)
	s3 = append(s3, s2...)
	fmt.Printf("appended g_slice s3 : %v\n", s3)
	s3 = append(s3, 4, 5, 6)
	fmt.Printf("last g_slice s3 : %v\n", s3)
}

//原切片基础上进行切片
func lastNumsBySlice(origin []int) []int {
	return origin[len(origin)-2:]
}

//创建了一个新的切片，将 origin 的最后两个元素拷贝到新切片上，然后返回新切片
func lastNumsByCopy(origin []int) []int {
	result := make([]int, 2)
	copy(result, origin[len(origin)-2:])
	return result
}

//随机生成 n 个 int 整数，64位机器上，一个 int 占 8 Byte，128 * 1024 个整数恰好占据 1 MB 的空间
func generateWithCap(n int) []int {
	rand.Seed(time.Now().UnixNano())
	nums := make([]int, 0, n)
	for i := 0; i < n; i++ {
		nums = append(nums, rand.Int())
	}
	return nums
}

func testLastChars(f func([]int) []int) {
	ans := make([][]int, 0)
	for k := 0; k < 100; k++ {
		origin := generateWithCap(128 * 1024) // 1M
		ans = append(ans, f(origin))
	}
	_ = ans
	var rtm runtime.MemStats
	runtime.ReadMemStats(&rtm)
	fmt.Println("%.2f MB", float64(rtm.Alloc)/1024./1024.)
	return
}
