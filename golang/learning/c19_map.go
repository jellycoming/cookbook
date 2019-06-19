package main

import "fmt"

type Program struct {
	Name    string
	Version string
}

// map是一组键和值的组合。
// 无序。
// 零值为 nil 。
// make 函数会返回给定类型的映射，并将其初始化备用。
func map1() {
	// 通过 make 创建
	m := make(map[string]Program)
	m["go"] = Program{Name: "golang", Version: "1.11.4"}
	m["py"] = Program{Name: "python", Version: "2.7.9"}
	fmt.Println(m)

	// 通过映射文法创建。若顶级类型只是一个类型名，你可以在文法的元素中省略它，此处可以省略 Program
	mm := map[string]Program{
		"go": {Name: "golang", Version: "1.11.4"},
		"py": {Name: "python", Version: "2.7.9"},
	}
	fmt.Println(mm)

	program := map[string]string{
		"golang": "1.11.4",
		"python": "2.7.9",
	}

	// 遍历映射
	for k, v := range program {
		if k == "golang" {
			fmt.Printf("we are using %s%s\n", k, v)
		}
	}

	// 插入或修改元素
	program["python"] = "3.5.6"
	program["py"] = "3.5.6"
	fmt.Println(program)

	// 获取元素
	goversion := program["golang"]
	fmt.Println(goversion)
	// 如果获取的 key 值不存在，会获取到映射的元素类型零值，不会报错（python会报 KeyError）
	nothing := program["go"]
	fmt.Println(nothing)

	// 删除元素
	delete(program, "py")
	// 如果删除的 key 值不存在，会执行一次空操作，不会报错（python会报 KeyError）
	delete(program, "pyc")
	fmt.Println(program)

	// 通过双赋值检测某个键是否存在
	// 若 key 在 m 中，ok 为 true ；否则，ok 为 false。若 key 不在映射中，那么 v 是该映射元素类型的零值。
	v, ok := program["java"]
	if ok == false {
		fmt.Println("The value:", v, "Present?", ok)
	}
	v, ok = program["golang"]
	if ok == true {
		fmt.Println("The value:", v, "Present?", ok)
	}
}

/**
可以在 make map 时指定map的大小，这样在 map 中存储的数据小于该值时，map不会进行频繁的扩容，达到优化性能的目的
*/
func map2() {
	m := make(map[string]string, 2)
	m["go"] = "golang"
	m["py"] = "python"
	m["ja"] = "java"
	fmt.Println(m)
}

// 映射可变
func testMutable() {
	m := make(map[string]string)
	m["go"] = "golang"
	fmt.Printf("before filled: %v\n", m) // map[go:golang]
	fillMap(m)
	fmt.Printf("after filled: %v\n", m) // map[go:golang py:python ja:java]
}

func fillMap(m map[string]string) {
	m["py"] = "python"
	m["ja"] = "java"
}

func main() {
	map1()
	map2()
	testMutable()
}
