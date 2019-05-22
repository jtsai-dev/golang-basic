package main

import "fmt"

type treeNode struct {
	value int
	left, right *treeNode
}
// "类"构造函数
func createTreeNode(value int) *treeNode {
	return &treeNode{value: value} // 创建一个变量返回对应的指针
}
// 为结构定义方法
func (node treeNode) print() {
	fmt.Println(node.value)
}

func main(){
	/* 面向对象：golang仅支持封装，不支持继承、多态
	 * 值接收者 & 指针接收者
	 *  需要改变内容则使用指针接收者
	 *  结构过大使用指针接收者(对应go只有值传递)
	 *  如有指针接收者，则尽可能统一使用指针接收者
	 */
	node := create()
	// node.print()
	preSequence(&node)
}

func create() treeNode {
	root := treeNode{ value: 3}
	root.left = &treeNode{}
	root.left.right = createTreeNode(2)
	root.right = &treeNode{ 5, nil, nil}
	root.right.left = new(treeNode)
	// fmt.Println(root)
	return root

	// nodes := []treeNode{
	// 	{ value: 3 },
	// 	{},
	// 	{ 4, nil, nil },
	// }
	// fmt.Println(nodes)
}

func preSequence(root *treeNode) {
	if root == nil {
		return
	}
	root.print()
	preSequence(root.left)
	preSequence(root.right)
}