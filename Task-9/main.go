package main

func main() {

}

//var arr [10]int
//for v := range arr {
//	fmt.Scan(&arr[v])
//}
//arr := make([]int, 8)
//for v := range arr {
//	arr = append(arr)
//}
//
//fmt.Println(len(arr), cap(arr))

//var slice = []int{5, 9, 8}
//s := make([]int, 3, 3)
//fmt.Println(s)
//n := copy(s, slice)
//
//fmt.Println(n)
//fmt.Println(s)

//}

// Delete slice with append
//a := []int{1, 2, 3, 4, 5, 6, 7}
//a = append(a[0:2], a[3:]...)
//fmt.Println(cap(a), len(a)) // [1 2 4 5 6 7]

//
//
//slice1 := []int{1, 2, 3, 4, 5}
//
//var MinIndex, MaxIndex int
//maximum := slice1[0]
//minimum := slice1[0]
//for i := 0; i < len(slice1); i++ {
////maximum
//if slice1[i] > maximum {
//maximum = slice1[i]
//MaxIndex = i
//}
////minimum
//if slice1[i] < minimum {
//minimum = slice1[i]
//MinIndex = i
//}
//}
//fmt.Println("maximum = ", maximum)
//fmt.Println("MaxIndex = ", MaxIndex)
//fmt.Println("minimum = ", minimum)
//fmt.Println("MinIndex = ", MinIndex)
//fmt.Println("MaxIndex =", len(slice1)-1, "Maximum", slice1[len(slice1)-1], "MinIndex =", 0, "Minimum =", slice1[0])
