package main

import "fmt"

//约瑟夫问题(单向环形链表应用实例)
type Boy struct {
	No   int  //编号
	Next *Boy //指向下一个小孩的指针

}

//编写一个函数构成单向环形链表
//num小孩的个数
//*Boy返回该环形链表的第一个小孩的指针
func AddBoy(num int) *Boy {
	first := &Boy{}  //空结点
	curBoy := &Boy{} //空结点
	//判断
	if num < 1 {
		fmt.Println("num的值错误")
		return first
	}
	//循环的构建这个环形链表
	for i := 1; i < num; i++ {
		boy := &Boy{
			No:   i,
			Next: nil,
		}
		//构成循环链表训要一个辅助指针（帮忙的）
		//1.第一个小孩比较特殊,如果只有一个小孩的情况
		if i == 1 { //第一个小孩
			first = boy //指向后就不能动了，因为是头指针
			curBoy = boy
			curBoy.Next = first //构成环状

		} else {
			curBoy.Next = boy
			curBoy = boy        //这句没懂
			curBoy.Next = first //构成环形

		}

	}
	return first

}

//显示单向环形链表[遍历]
func ShowBoy(first *Boy) {
	//处理空环形链表
	if first.Next == nil {
		fmt.Println("链表为空")
		return
	}
	//至少有一个小孩，创建一个指针帮助遍历
	temp := first

	for {
		fmt.Printf("小孩编号=%d", temp.No)
		//遍历完成，退出的条件是：环形的最后结点指向头结点first
		if temp.Next == first {
			break
		}
		//移动下一个
		temp = temp.Next

	}

}

//约瑟夫问题处理
/*
设编号为1,2,.....n的n个人围坐一圈，约定编号为k(1<=k<=n)的
人从1开始报数，数到m的那个人出列，他的下一位又从1开始报数,数到
m的那个人又出列，依次类推，直到所有人出列为止，由此产生一个出队编号的序列
*/
//思路分析
//1.编写一个函数，数小孩,PlayGame(first *Boy,startNo int,countNum int)
//2.最后我们使用一个算法，按照要求，在环形链表中留下最后一个人
func PlayGame(first *Boy, startNo int, countNum int) {
	//1.空链表我们单独处理一下
	if first.Next == nil {
		fmt.Println("空链表，没有小孩")
		return
	}
	//判断startNo<=小孩的总数
	//2.定义辅助指针，帮助我们删除小孩
	tail := first
	//3.让tail指向环形链表的最后一个小孩,因为tail在删除小孩的时候使用的到
	for {
		if tail.Next == first {
			//说明到最后一个小孩了
			break

		}
		tail = tail.Next
	}
	//4.让first 移动到startNo[后面删除小孩就以first为准]
	for i := 0; i <= startNo-1; i++ {
		first = first.Next
		tail = tail.Next
	}
	//5.开始数countNum下，然后就删除first指向的小孩
	for {
		//开始数countNum-1次
		for i := 1; i <= countNum; i++ {
			first = first.Next
			tail = tail.Next
		}
		fmt.Printf("小孩编号为%d的出圈\n", first.No)
		//删除first指向的结点
		first = first.Next
		tail.Next = first
		//判断如果tail==first，则圈中只有一个小孩
		if tail == first {
			break
		}
	}
	fmt.Printf("最后的小孩编号为：%d\n", first.No)

}
func main() {
	first := AddBoy(5)
	//ShowBoy(first)
	PlayGame(first,2,3)

}
