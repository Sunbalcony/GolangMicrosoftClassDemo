package main

import (
	"fmt"
)

// HeroNode 单链表：为了比较好的对单链表进行增删改查的操作，我们会给单链表设置一个头结点，头结点的主要作用是用来标识链表头。
//头结点本身不存放数据
type HeroNode struct {
	no       int
	name     string
	nickname string
	next     *HeroNode
}

//给链表插入一个结点
//第一种方法，在单链表的最后加入

func InsertHeroNode(head *HeroNode, newHeroNode *HeroNode) {
	//思路
	//1.先找到该链表的最后一个结点
	//2.创建一个辅助结点(跑龙套帮忙)
	temp := head
	for {
		if temp.next == nil { //标识找到最后的结点
			break
		}
		temp = temp.next //让temp不断地指向下一个结点
	}
	//3.将newHeroNode加入到链表的最后
	temp.next = newHeroNode

}

func InsertHeroNode2(head *HeroNode, newHeroNode *HeroNode) {
	//最常用的链表方法
	//第二种插入方法。根据no的编号从小到大插入编写
	//1.找到适当的节点位置
	//2.创建辅助结点
	temp := head
	var ff = true
	//让插入的结点的no和temp的下一个结点的no作比较
	for {
		if temp.next == nil {
			break
		} else if temp.next.no > newHeroNode.no {
			//说明newHeroNode就应该插入到temp后面
			break

		} else if temp.next.no == newHeroNode.no {
			//说明我们的链表中已经有这个no，不需要插入
			ff = false
			break

		}
		temp = temp.next

	}
	if !ff {
		fmt.Println("对不起，已经存在这个no", newHeroNode.no)
		return
	} else {
		newHeroNode.next = temp.next
		temp.next = newHeroNode
	}

}

//显示链表的所有信息
func ListHeroNode(head *HeroNode) {
	//1.创建一个辅助结点(跑龙套帮忙)
	temp := head
	//先判断链表是不是空
	if temp.next == nil {
		fmt.Println("该链表为空链表")
		return
	}
	//2.遍历这个链表
	for {
		fmt.Printf("%d,%s,%s", temp.next.no, temp.next.name, temp.next.nickname)
		//判断是否到链表最后
		temp = temp.next
		if temp.next == nil {
			break
		}

	}

}
func DelHeroNode(head *HeroNode, id int) {
	//删除一个结点
	temp := head
	var ff = false
	//找到要删除的结点的no，和temp的下一个结点比较
	for {
		if temp.next == nil {
			//说明找到最后一个结点都没有
			break
		} else if temp.next.no == id {
			//说明找到了
			ff = true
			break

		}
		temp = temp.next

	}
	if ff {
		//找到删除
		temp.next = temp.next.next
	} else {
		fmt.Println("对不起，删除的id不存在")
	}

}

func main() {
	//使用带head头的单向链表

	//1 先创建一个头结点
	head := &HeroNode{}
	//2 创建一个新的HeroNode
	hero1 := &HeroNode{
		no:       1,
		name:     "宋江",
		nickname: "及时雨",
	}
	hero2 := &HeroNode{
		no:       2,
		name:     "卢俊义",
		nickname: "玉麒麟",
	}
	hero3 := &HeroNode{
		no:       3,
		name:     "林冲",
		nickname: "豹子头",
	}
	hero4 := &HeroNode{
		no:       5,
		name:     "吴用",
		nickname: "智多星",
	}
	////3.加入
	//InsertHeroNode(head, hero1)
	//InsertHeroNode(head, hero2)
	//InsertHeroNode(head, hero3)
	////打印：1,宋江,及时雨2,卢俊义,玉麒麟3,林冲,豹子头
	////链表将要求顺序,如果给定了随机的数要按次序排列
	////InsertHeroNode(head, hero3)
	////InsertHeroNode(head, hero1)
	////InsertHeroNode(head, hero2)
	////打印：3,林冲,豹子头1,宋江,及时雨2,卢俊义,玉麒麟,这样不符合要求，需要改进插入方法

	InsertHeroNode2(head, hero3)
	InsertHeroNode2(head, hero2)
	InsertHeroNode2(head, hero1)
	InsertHeroNode2(head, hero4)
	//4.显示
	//ListHeroNode(head)
	//5.删除
	DelHeroNode(head, 3)
	ListHeroNode(head)



}
