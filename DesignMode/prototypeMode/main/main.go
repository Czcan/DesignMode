package main

import (
	"fmt"
)

func main() {
	file1 := &file{name: "File1"}
	file2 := &file{name: "File2"}
	file3 := &file{name: "File3"}

	folder1 := &folder{
		children: []inode{file1},
		name:     "Folder1",
	}
	folder2 := &folder{
		children: []inode{folder1, file2, file3},
		name:     "Folder2",
	}

	fmt.Println("\nPrinting hierarchy for Folder2")
	folder2.print("    ")

	cloneFolder := folder2.clone()
	fmt.Println("\nPrinting hierarchy for clone Folder")
	cloneFolder.print("    ")
}

/*
原型模式优点：1. 可以克隆对象，无需与其所属的类耦合
			2. 可克隆预生成原型，避免反复初始化代码（如ios 类中 public static shared: classname() -> 别处用到该类 classname.shared.xxx）
			3. 更方便生成复杂的对象（比如 初始化时参数很多的情况）
			4. 使用继承以外的方式处理 复杂对象 的 不同配置
		缺点： 克隆包含循环引用的对象时可能非常麻烦
*/
