package main

//原型接口
type inode interface {
	print(string)
	clone() inode
}
