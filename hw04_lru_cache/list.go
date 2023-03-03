package hw04lrucache

type List interface {
	Len() int
	Front() *ListItem
	Back() *ListItem
	PushFront(v interface{}) *ListItem
	PushBack(v interface{}) *ListItem
	Remove(i *ListItem)
	MoveToFront(i *ListItem)
}

type ListItem struct {
	Value interface{}
	Next  *ListItem
	Prev  *ListItem
}

type list struct {
	len   int
	front *ListItem
	back  *ListItem
}

func (l *list) Len() (i int) {
	return l.len
}

func (l *list) addLen(i int) {
	l.len += i
}

func (l *list) Front() *ListItem {
	return l.front
}

func (l *list) Back() (ele *ListItem) {
	return l.back
}

func (l *list) PushFront(v interface{}) (li *ListItem) {
	defer l.addLen(1)
	li = &ListItem{
		Value: v,
		Next:  nil,
		Prev:  nil,
	}
	if l.front == nil {
		l.front = li
		l.back = li
		return
	}
	li.Next = l.front
	l.front.Prev = li
	l.front = li
	return
}

func (l *list) PushBack(v interface{}) (li *ListItem) {
	defer l.addLen(1)
	li = &ListItem{
		Value: v,
		Next:  nil,
		Prev:  nil,
	}
	if l.back == nil {
		l.front = li
		l.back = li
		return
	}
	li.Prev = l.back
	l.back.Next = li
	l.back = li
	return
}

func (l *list) Remove(i *ListItem) {
	defer l.addLen(-1)
	defer func() {
		i.Next, i.Prev = nil, nil
	}()

	if i.Next == nil && i.Prev == nil {
		l.front = nil
		l.back = nil
		return
	}

	if l.front == i {
		l.front = i.Next
	}
	if l.back == i {
		l.back = i.Prev
	}

	if i.Next != nil {
		i.Next.Prev = i.Prev
	}
	if i.Prev != nil {
		i.Prev.Next = i.Next
	}
}

func (l *list) MoveToFront(i *ListItem) {
	l.Remove(i)
	i.Next = nil
	i.Prev = nil

	if l.front == nil {
		l.front = i
		return
	}
	i.Next = l.front
	l.front.Prev = i
	l.front = i
	return
}

func NewList() List {
	return new(list)
}
