package main

import "fmt"

var arr = [...]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
var slice0 []int = arr[2:8]
var slice1 []int = arr[0:6]        //可以简写为 var slice []int = arr[:end]
var slice2 []int = arr[5:10]       //可以简写为 var slice[]int = arr[start:]
var slice3 []int = arr[0:len(arr)] //var slice []int = arr[:]
var slice4 = arr[:len(arr)-1]      //去掉切片的最后一个元素

func main() {
	////前包后不包（下标）
	//fmt.Printf("全局变量：arr %v\n", arr)
	//fmt.Printf("全局变量：slice0 %v\n", slice0)
	//fmt.Printf("全局变量：slice1 %v\n", slice1)
	//fmt.Printf("全局变量：slice2 %v\n", slice2)
	//fmt.Printf("全局变量：slice3 %v\n", slice3)
	//fmt.Printf("全局变量：slice4 %v\n", slice4)
	//fmt.Printf("-----------------------------------\n")
	//arr2 := [...]int{9, 8, 7, 6, 5, 4, 3, 2, 1, 0}
	//slice5 := arr[2:8]
	//slice6 := arr[0:6]         //可以简写为 slice := arr[:end]
	//slice7 := arr[5:10]        //可以简写为 slice := arr[start:]
	//slice8 := arr[0:len(arr)]  //slice := arr[:]
	//slice9 := arr[:len(arr)-1] //去掉切片的最后一个元素
	//fmt.Printf("局部变量： arr2 %v\n", arr2)
	//fmt.Printf("局部变量： slice5 %v\n", slice5)
	//fmt.Printf("局部变量： slice6 %v\n", slice6)
	//fmt.Printf("局部变量： slice7 %v\n", slice7)
	//fmt.Printf("局部变量： slice8 %v\n", slice8)
	//fmt.Printf("局部变量： slice9 %v\n", slice9)
	//
	////未初始化的slice为nil
	//var s1 []int
	//if s1 == nil {
	//	fmt.Println("是空")
	//} else {
	//	fmt.Println("不是空")
	//}
	//s2 := []int{}
	//// 3.make()
	//var s3 []int = make([]int, 0)
	//fmt.Println(s1 == nil, s2 == nil, s3 == nil)
	//fmt.Println(s1, s2, s3)
	//
	//// 4.初始化赋值
	//var s4 []int = make([]int, 0, 0)
	//fmt.Println(s4)
	//s5 := []int{1, 2, 3}
	//fmt.Println(s5)
	//
	//// 读写操作实际目标是底层数组，只需注意索引号的差别。
	//data := [...]int{0, 1, 2, 3, 4, 5}
	//s := data[2:4]
	//s[0] += 100
	//s[1] += 200
	//fmt.Println(s)
	//fmt.Println(data)
	//
	//// 使用 make 动态创建slice，避免了数组必须用常量做长度的麻烦。还可用指针直接访问底层数组，退化成普通数组操作。
	//ss := []int{0, 1, 2, 3}
	//p := &ss[2] // *int, 获取底层数组元素指针。
	//*p += 100
	//fmt.Println(ss)


	//向 slice 尾部添加数据，返回新的 slice 对象。
	//s1 := make([]int, 0, 5)
	//fmt.Printf("%p\n", &s1)
	//
	//s2 := append(s1, 1)
	//fmt.Printf("%p\n", &s2)
	//fmt.Println(s1, s2)

	//切片拷贝，copy ：函数 copy 在两个 slice 间复制数据，复制长度以 len 小的为准。两个 slice 可指向同一底层数组，允许元素区间重叠。
	s1 := []int{1, 2, 3, 4, 5}
	fmt.Printf("slice s1 : %v\n", s1)
	s2 := make([]int, 10)
	fmt.Printf("slice s2 : %v\n", s2)
	copy(s2, s1)
	fmt.Printf("copied slice s1 : %v\n", s1)
	fmt.Printf("copied slice s2 : %v\n", s2)
	s3 := []int{1, 2, 3}
	fmt.Printf("slice s3 : %v\n", s3)
	s3 = append(s3, s2...)
	fmt.Printf("appended slice s3 : %v\n", s3)
	s3 = append(s3, 4, 5, 6)
	fmt.Printf("last slice s3 : %v\n", s3)

	//切片resize
	var a = []int{1, 3, 4, 5}
	fmt.Printf("slice a : %v , len(a) : %v\n, cap(a) : %v\n", a, len(a), cap(a))
	b := a[1:2]
	fmt.Printf("slice b : %v , len(b) : %v\n, cap(b) : %v\n", b, len(b), cap(b))
	c := b[0:3] // 截取
	fmt.Printf("slice c : %v , len(c) : %v\n, cap(c) : %v\n", c, len(c), cap(c))
	fmt.Println(b[1])
}


