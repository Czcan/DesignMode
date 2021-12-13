package main

//组合模式
func main() {
	file1 := &file{name: "File1"}
	file2 := &file{name: "File2"}
	file3 := &file{name: "File3"}

	folder1 := &folder{
		name: "Folder1",
	}

	folder1.add(file1)

	folder2 := &folder{
		name: "Folder2",
	}
	folder2.add(file2)
	folder2.add(file3)
	folder2.add(folder1)

	folder2.search("rose")
}

/*
组合模式：1. 需要实现树状对象结构时
		2. 实现客户端使用相同模式处理简单和复杂元素

		优点：1. 利用多态和递归机制更方便地使用复杂树结构
			 2. 符合开闭原则，不用修改原有代码，就可以在应用中添加新元素， 使其成为树的一部分
		缺点： 对于功能差异较大的类，提供公共接口较困难。特定情况可能需要过度一般化公共接口，会变得难以理解
*/
