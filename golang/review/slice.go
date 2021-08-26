// 切片的内部实现结构通过指针引用底层数组，是对数组一段连续片段的引用
package main

import (
	"fmt"
	"reflect"
	"sync"
	"unsafe"
)

// 切片的数据结构
type slice struct {
	array unsafe.Pointer
	len   int
	cap   int
}

// 切片的运行时表示
type SliceHeader struct {
	Data uintptr // 指向数组的指针
	Len  int     // 当前切片的长度
	Cap  int     // 当前切片的容量，即Data指向的数组的大小
}

type _type struct {
}

// 切片扩容
// runtime.growslice
func growslice(et *_type, old slice, cap int) slice {
	// ...
	newcap := old.cap
	doublecap := newcap + newcap
	// 如果新申请的容量（cap）大于2倍的旧容量（old.cap），最终容量（newcap）就会使用新申请的容量（cap）
	if cap > doublecap {
		newcap = cap
	} else {
		// 如果当前切片的容量（old.cap）小于1024，最终容量（newcap）就是将当前容量翻倍
		if old.cap < 1024 {
			newcap = doublecap
		} else {
			// 如果当前容量大于1024，最终容量将从原来的旧容量（old.cap）开始循环增加1/4，直到最终容量（newcap）大于等于新申请的容量（cap）
			// Check 0 < newcap to detect overflow
			// and prevent an infinite loop.
			for 0 < newcap && newcap < cap {
				newcap += newcap / 4
			}
			// 如果最终容量（newcap）计算值溢出，则最终容量（newcap）就是新申请的容量（cap）
			// Set newcap to the requested cap when
			// the newcap calculation overflowed.
			if newcap <= 0 {
				newcap = cap
			}
		}
	}
	var p unsafe.Pointer
	// ...
	// 旧数组的拷贝，最终返回长度为就切片的长度，容量为最终新容量的新的切片
	// memmove(p, old.array, lenmem)
	return slice{p, old.len, newcap}
}

// 切片拷贝
// 可以在不同长度的切片间进行拷贝，内存拷贝的长度取决于较短的那个切片
// runtime.slicecopy
func slicecopy(toPtr unsafe.Pointer, toLen int, fromPtr unsafe.Pointer, fromLen int, width uintptr) int {
	// 源切片或目标切片有一个长度为0，则直接返回
	if fromLen == 0 || toLen == 0 {
		return 0
	}
	// n记录源切片和目标切片较短的那一个的长度
	n := fromLen
	if toLen < n {
		n = toLen
	}

	if width == 0 {
		return n
	}

	size := uintptr(n) * width
	//if raceenabled {
	//	callerpc := getcallerpc()
	//	pc := funcPC(slicecopy)
	//	racereadrangepc(fromPtr, size, callerpc, pc)
	//	racewriterangepc(toPtr, size, callerpc, pc)
	//}
	//if msanenabled {
	//	msanread(fromPtr, size)
	//	msanwrite(toPtr, size)
	//}

	if size == 1 { // common case worth about 2x to do here
		// TODO: is this still worth it with new memmove impl?
		*(*byte)(toPtr) = *(*byte)(fromPtr) // known to be a byte pointer
	} else {
		// 整块内存拷贝
		//memmove(toPtr, fromPtr, size)
	}
	return n
}

// nil切片和空切片
func empty() {
	// nil切片底层数组的指针地址为0，空切片引用的数组的指针地址是存在的，但没有分配任何内存空间呢。
	// 不同空切片引用的数组的指针地址为同一个固定值
	var a []int            // nil切片
	var b = make([]int, 0) // 空切片
	var c = []int{}        // 空切片，字面量创建
	aa := *(*reflect.SliceHeader)(unsafe.Pointer(&a))
	bb := *(*reflect.SliceHeader)(unsafe.Pointer(&b))
	cc := *(*reflect.SliceHeader)(unsafe.Pointer(&c))
	// a={0 0 0} b={824634522680 0 0} c={824634522680 0 0}
	fmt.Printf("a=%v b=%v c=%v\n", aa, bb, cc)
	// 可以向nil切片或空切片append元素
	a = append(a, 1)
	b = append(b, 2)
	// a=[1] b=[2]
	fmt.Printf("a=%v b=%v\n", a, b)
}

// 并发安全性
// slice底层实现中并未涉及并发相关控制，不是并发安全的数据结构
func concurrency() {
	s := []int{1}
	wg := &sync.WaitGroup{}
	for i := 0; i < 4; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for n := 0; n < 10000; n++ {
				s[0]++
			}
		}()
	}
	wg.Wait()
	// 结果不确定
	fmt.Printf("%v", s)
}

type A struct {
	Name  string
	Score int
}

// 不同的切片底层引用相同的数组时，要注意对其中一个切片的修改可能会影响其他切片
func main() {
	// 切片的底层数组引用及扩容
	a := make([]int, 4)
	b := a[:2]
	// a与b底层引用相同的数组，a对数组的可读片段为[0:4]，b对数组的可读片段为[0:2]
	// output: a=[0 0 0 0] b=[0 0]
	fmt.Printf("a=%v b=%v\n", a, b)
	a[0] = 1
	b[1] = 2
	// 由于a与b底层引用相同的数组，上面两行代码对a与b的改变会影响到对方
	// output: a=[1 2 0 0] b=[1 2]
	fmt.Printf("a=%v b=%v\n", a, b)
	// 获取a与b所引用数组的内存地址，二者一样
	// output: a=0xc00001a080 b=0xc00001a080
	fmt.Printf("a=%p b=%p\n", unsafe.Pointer(&a[0]), unsafe.Pointer(&b[0]))
	// 向a追加元素，由于之前所引用数组的容量为4，再往里append元素时会引发扩容操作，扩容操作会分配一个更大的数组（这里容量为8），并将原来的数组元素的值拷贝到里面，再使a底层的数组指针指向这个新数组
	// 在切片长度小于切片容量时，append操作不会触发扩容，如a := make([]int, 0, 4)，前4次的append操作不会触发扩容
	// 切片扩容的实现代码为runtime.growslice
	// 由于a底层引用的数组发生了改变，而b依然引用原来的数组，所以对a与b的修改不再影响对方
	// 但如果切片内存储的元素是指针类型，虽然拷贝会生成不同的指针值，但不同的指针值指向的地址会是同一块，所以对这块相同内存的改变会对所有指向这里的指针可见
	a = append(a, 9)
	a[0] = 9
	b[1] = 3
	// output: a=[9 2 0 0 9] b=[1 3]
	fmt.Printf("a=%v b=%v\n", a, b)
	// a底层数组指向一个新的地址，b还是原来的地址
	// output: a=0xc00001c080 b=0xc00001a080
	fmt.Printf("a=%p b=%p\n", unsafe.Pointer(&a[0]), unsafe.Pointer(&b[0]))

	// 切片拷贝
	m := make([]int, 3)
	n := make([]int, 8)
	fmt.Println(copy(m, a)) // 3，m的长度
	fmt.Println(copy(n, a)) // 5，a的长度
	// a=[9 2 0 0 9] m=[9 2 0] n=[9 2 0 0 9 0 0 0]
	fmt.Printf("a=%v m=%v n=%v\n", a, m, n)
	a[0] = 10
	m[1] = 20
	n[2] = 30
	// 由于切片拷贝是整块内存的拷贝，所以拷贝完成后的3个切片内部指向了相互独立的3个内存块，修改互不影响
	// 但如果切片存储的是指针类型，拷贝操作复制的是指针值，修改其中一个指针指向的内存块同时会对其他指针可见
	// [10 2 0 0 9] m=[9 20 0] n=[9 2 30 0 9 0 0 0]
	fmt.Printf("a=%v m=%v n=%v\n", a, m, n)
	s1 := []*A{&A{Name: "a", Score: 1}, &A{Name: "b", Score: 2}, &A{Name: "c", Score: 3}}
	s2 := make([]*A, 2)
	s3 := make([]*A, 4)
	fmt.Printf("%d %d\n", copy(s2, s1), copy(s3, s1)) // 2 3
	s1[0].Score = 100
	// s1[0]的修改对s2，s3可见
	// &{a 100} &{a 100} &{a 100}
	fmt.Printf("%v %v %v\n", s1[0], s2[0], s3[0])

	// nil切片和空切片
	empty()

	// 切片间不能通过==比较，但切片可以和nil比较
	// fmt.Println(a == b) // invalid
	fmt.Println(a == nil) // false

	// 并发安全性
	concurrency()
}
