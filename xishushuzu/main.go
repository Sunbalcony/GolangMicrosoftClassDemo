package main

import (
	"fmt"
)

type ValNode struct {
	row int //行
	col int //列
	val int //数据

}

func main() {
	//创建一个原始数组
	var chessMap [11][11]int
	chessMap[1][2] = 1 //黑子
	chessMap[2][3] = 2 //蓝子
	for _, v := range chessMap {
		for _, v2 := range v {
			fmt.Printf("%d\t", v2)

		}
		fmt.Println()
	}
//转换为稀疏数组 想-->算法
//思路
//遍历chessMap，如果有一个值不为0，那么就创建一个node结构体，将其放入slice中
	var sparseArr []ValNode
	//标准的稀疏数组应该还有一个记录原始数组规模（行和列 的默认值）
	valNode := ValNode{
		row: 11,
		col: 11,
		val: 0,
	}
	sparseArr = append(sparseArr, valNode)

	for i, v := range chessMap {
		for j, v2 := range v {
			if v2 != 0 {
				//记录值
				//创建一个ValNode类型的值结点
				valNode := ValNode{
					row: i,
					col: j,
					val: v2,
				}
				sparseArr = append(sparseArr, valNode)
			}
		}
	}
	//输出稀疏数组
	for i, valNode := range sparseArr {
		fmt.Println(i, valNode)
	}
	//从稀疏数组恢复为原始数组

	//创建一个原始数组(go中创建切片)
	var chessMap2 [11][11]int
	//遍历稀疏数组
	for i, vv := range sparseArr {
		//chessMap2[vv.row][vv.col] = vv.val 提示数组越界,因为读取了原始数组规模那一行，因此我们将稀疏数组下标为0的元素忽略掉
		if i != 0 {
			chessMap2[vv.row][vv.col] = vv.val
		}
	}
	//确认原始数组是否恢复
	for _, v := range chessMap2 {
		for _, v2 := range v {
			fmt.Printf("%d\t", v2)

		}
		fmt.Println()
	}

}
