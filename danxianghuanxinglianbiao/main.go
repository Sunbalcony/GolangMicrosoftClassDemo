package main

import "fmt"

type CatNode struct {
	no   int //猫猫的编号
	name string
	next *CatNode
}

func InsertCatNode(head *CatNode, newCatNode *CatNode) {
	//判断是不是添加第一只猫
	if head.next == nil {
		//如果头猫为空，那么就将添加的新猫的no和name给head，当做头猫用
		head.no = newCatNode.no
		head.name = newCatNode.name
		head.next = head //构成环状
		fmt.Println(newCatNode, "加入到环形链表中")
		return
	}
	//定义一个临时变量，帮忙:找到环形的最后结点
	temp := head
	for {
		if temp.next == head {
			break
		}
		temp = temp.next
	}
	//加入到链表中
	temp.next = newCatNode
	newCatNode.next = head

}

//删除一个结点
func DelCatNode(head *CatNode, id int) *CatNode {
	//删除一个环形单向链表结点的思路如下：
	//1.先让temp指向head
	//2.让helper指向环形链表最后
	//3.temp和要删除的id进行比较，如果相同，则同helper完成删除(这里必须考虑如果删除的结点是头结点怎么办？)

	temp := head
	helper := head
	//空链表
	if temp.next == nil {
		fmt.Println("这是一个空环形链表，无法删除")
		return head
	}
	//如果只有一个结点
	if temp.next == head { //只有一个结点
		temp.next = nil
		return head
	}
	//将helper定位到环形链表最后
	for {
		if helper.next == head {
			break
		}
		helper = helper.next
	}
	//如果有两个或两个以上结点
	var flag = true
	for {
		if temp.next == head { //如果到这里,说明比较到最后一个【最后一个还没有比较】
			break

		}
		if temp.no == id {
			if temp == head { //说明删除的是头结点
				head = head.next

			}
			//找到了要删除的id
			helper.next = temp.next
			fmt.Printf("删除的结点为%d", id)
			flag = false
			break

		}
		//temp,helper在不断移动
		temp = temp.next     //移动【比较】
		helper = helper.next //移动【一旦找到要删除的结点 helper】
	}
	//这里还有一次比较
	if flag { //如果flag为真，则我们上面没有删除
		if temp.no == id {
			//找到了要删除的id
			helper.next = temp.next
			fmt.Printf("删除的结点为%d", id)

		} else {
			fmt.Printf("这个结点没有%d", id)
		}

	}
	return head

}

//输出环形链表
func ListCircle(head *CatNode) {
	fmt.Println("环形链表的具体情况是:")
	temp := head
	if temp.next == nil {
		fmt.Println("空环形链表，不玩了结束")
		return
	}

	for {
		fmt.Printf("猫的信息为:id=%d,name=%s ->", temp.no, temp.name)
		if temp.next == head {
			//如果temp的下一个结点是head头结点，那么已经到了环形链表的最后了
			break

		}
		temp = temp.next
	}

}
func main() {
	//初始化环形链表头结点
	head := &CatNode{}
	//创建一只猫
	cat1 := &CatNode{
		no:   1,
		name: "tom",
	}
	cat2 := &CatNode{
		no:   2,
		name: "tom2",
	}
	cat3 := &CatNode{
		no:   3,
		name: "tom3",
	}
	InsertCatNode(head, cat1)
	InsertCatNode(head, cat2)
	InsertCatNode(head, cat3)

	ListCircle(head)
	head = DelCatNode(head, 3)
	ListCircle(head)

}
