package dtst

import "fmt"

type LinkedList struct {
	header *node
	length int
}

type node struct {
	data interface{}
	next *node
}

func New() *LinkedList {
	return &LinkedList{}
}

func (list *LinkedList) Length() int {
	return list.length
}

func (list *LinkedList) Append(data interface{}) {
	if list.header == nil {
		list.header = &node{data, nil}
	} else {
		curr := list.header
		for curr.next != nil {
			curr = curr.next
		}
		curr.next = &node{data, nil}
	}
	list.length++
}

func (list *LinkedList) AppendAll(dataList ...interface{}) {
	tempHeader := &node{}
	tempCurr := tempHeader
	for i, data := range dataList {
		if i == 0 {
			tempHeader.data = data
		} else {
			tempCurr.next = &node{data, nil}
			tempCurr = tempCurr.next
		}
	}
	if list.header == nil {
		list.header = tempHeader
	} else {
		curr := list.header
		for curr.next != nil {
			curr = curr.next
		}
		curr.next = tempHeader
	}
	list.length += len(dataList)
}

func (list *LinkedList) Prepend(data interface{}) {
	if list.header == nil {
		list.header = &node{data, nil}
	} else {
		list.header = &node{data, list.header}
	}
	list.length++
}

// delete the first occurrence of the specified data from the list, if it is present
func (list *LinkedList) Delete(data interface{}) {
	if list.header != nil {
		if list.header.data == data {
			// delete first node
			match := list.header
			list.header = match.next
			match.next = nil
			list.length--
		} else {
			// delete middle or last node
			curr := list.header
			for curr.next != nil {
				if curr.next.data == data {
					match := curr.next
					curr.next = match.next
					match.next = nil
					list.length--
					break
				}
				curr = curr.next
			}
		}
	}

}

func (list *LinkedList) Get(index int) interface{} {
	i := 0
	for curr := list.header; curr != nil; curr = curr.next {
		if i == index {
			return curr.data
		}
		i++
	}
	return nil
}

func (list *LinkedList) Contains(data interface{}) bool {
	for curr := list.header; curr != nil; curr = curr.next {
		if curr.data == data {
			return true
		}
	}
	return false
}

func (list *LinkedList) Clear() {
	curr := list.header
	for curr != nil {
		list.header = list.header.next
		curr.next = nil
		curr = list.header
	}
	list.length = 0
}

func (list *LinkedList) String() (s string) {
	if list.header == nil {
		s = "&LinkedList[]"
	} else {
		s = fmt.Sprintf("&LinkedList[(%v)", list.header.data)
		for curr := list.header.next; curr != nil; curr = curr.next {
			s += fmt.Sprintf("->(%v)", curr.data)
		}
		s += "]"
	}
	return
}
