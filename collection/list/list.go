package list

import "fmt"

type Collection struct {
	first, last *Node
	length      int
}

type Node struct {
	next, prev *Node
	value      int
}

func (c *Collection) firstAdd(element int) {
	c.first = &Node{nil, nil, element}
	c.last = c.First()
}

func (c *Collection) otherAdd(element int) {
	newNode := &Node{nil, nil, element}
	c.last.next = newNode
	newNode.prev = c.Last()
	c.last = newNode
}

func (c *Collection) Add(element int) {
	if c.length == 0 {
		c.firstAdd(element)
	} else {
		c.otherAdd(element)
	}
	c.length++
}

func (c *Collection) verification(element int) bool {
	if element > c.length || element < 0 {
		return true
	}
	return false
}

func (c *Collection) returnLink(element int) *Node {
	pos := 0
	node := c.first
	for {
		if pos == element {
			return node
		}
		node = node.next
		pos++
	}
}

func (c *Collection) Get(element int) *Node {
	if c.verification(element) {
		return nil
	} else {
		return c.returnLink(element)
	}
}

func (с *Collection) removeVerification(index int) bool {
	if index >= 0 && index <= с.length-1 {
		return true
	} else {
		return false
	}
}

func (с *Collection) Remove(index int) {
	if с.removeVerification(index) {
		var node *Node
		if с.length-index < index {
			node = с.Last()
			for i := с.length - 1; i != index; i-- {
				node = node.Prev()
			}
		} else {
			node = с.First()
			for i := 0; i != index; i++ {
				node = node.Next()
			}
		}
		switch node {
		case с.First():
			с.first = node.Next()
		case с.Last():
			с.last = node.Prev()
		}
		if node.prev != nil {
			node.prev.next = node.Next()
		}
		if node.next != nil {
			node.next.prev = node.Prev()
		}
		node = nil
		с.length--
	}
}

func (с *Collection) First() *Node {
	return с.first
}

func (с *Collection) Last() *Node {
	return с.last
}

func (с *Collection) Length() int {
	return с.length
}

func (c *Collection) Print() {
	node := c.First()
	for node != nil {
		fmt.Printf("%v ", node.Value())
		node = node.Next()
	}
	fmt.Println()
}

func (n *Node) Next() *Node {
	return n.next
}

func (n *Node) Prev() *Node {
	return n.prev
}

func (n *Node) Value() int {
	return n.value
}
