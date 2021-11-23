package main

import "fmt"

type HeroNode struct {
	no       int
	name     string
	nickname string
	pre      *HeroNode //指向前一个结点
	next     *HeroNode //指向下一个结点
}

//给双向链表插入一个结点
func InsertHeroNode(head *HeroNode, newHeroNode *HeroNode) {
	temp := head
	for {
		if temp.next == nil {
			break
		}
		temp = temp.next
	}
	temp.next = newHeroNode
	newHeroNode.pre = temp

}

//给双向链表中插入一个结点，按照no的序号从小到大
func InsertHeroNode2(head *HeroNode, newHeroNode *HeroNode) {
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
		newHeroNode.pre = temp
		//temp.next.pre = newHeroNode //(存在漏洞。刚好是最后一个结点则没有前驱结点，因此要判断一下结点为不为空)
		if temp.next != nil {
			temp.next.pre = newHeroNode
			temp.next = newHeroNode

		}

	}

}

//删除双向链表的一个结点
func DelHeroNode(head *HeroNode, id int) {
	temp := head
	var flag = false
	for {
		if temp.next == nil {
			break
		} else if temp.next.no == id {
			flag = true
			break

		}
		temp = temp.next
	}
	if flag {
		//找到删除
		temp.next = temp.next.next //如果是最后一个结点，则存在风险
		if temp.next != nil {
			temp.next.pre = temp
		}
	} else {
		fmt.Println("要删除的id不存在")
	}

}
func ListHeroNode(head *HeroNode) {
	temp := head
	//先判断链表是否为空
	if temp.next == nil {
		fmt.Println("链表空的一批")
		return
	}
	for {
		fmt.Printf("%d,%s,%s", temp.next.no, temp.next.name, temp.next.nickname)
		temp = temp.next
		if temp.next == nil {
			break
		}
	}
}

//从后往前遍历双向链表
func ListHeroNode2(head *HeroNode) {
	temp := head
	//1。先判断链表是否为空
	if temp.next == nil {
		fmt.Println("双向链表空的一批")
		return
	}
	//2.让temp定位到双向链表的最后结点,也就是next为nil
	for {
		if temp.next == nil {
			break
		}
		temp = temp.next

	}
	//3.遍历这个链表
	for {
		fmt.Printf("[%d,%s,%s]", temp.no, temp.name, temp.nickname)
		temp = temp.pre
		//判断链表头是否为空
		if temp.pre == nil {
			break
		}
	}

}

func main() {
	head := &HeroNode{}
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
	InsertHeroNode(head, hero1)
	InsertHeroNode(head, hero2)
	InsertHeroNode(head, hero3)
	//InsertHeroNode(head, hero4)
	ListHeroNode(head)
	ListHeroNode2(head) //逆序打印
	InsertHeroNode(head,hero4)
}
