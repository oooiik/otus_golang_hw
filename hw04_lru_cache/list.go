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
	ele *ListItem
}

func (l *list) Len() (i int) {
	i = 0
	if l.ele == nil {
		return
	}
	l.ele = l.Front()
	for {
		i++
		if l.ele.Next == nil {
			return
		}
		l.ele = l.ele.Next
	}
}

func (l *list) Front() *ListItem {
	if l.ele == nil {
		return nil
	}
	for {
		if l.ele.Prev == nil {
			return l.ele
		}
		l.ele = l.ele.Prev
	}
}

func (l *list) Back() (ele *ListItem) {
	if l.ele == nil {
		return nil
	}
	for {
		if l.ele.Next == nil {
			return l.ele
		}
		l.ele = l.ele.Next
	}
}

func (l *list) PushFront(v interface{}) (li *ListItem) {
	li = &ListItem{
		Value: v,
		Next:  nil,
		Prev:  nil,
	}
	if l.Len() == 0 {
		l.ele = li
		return
	}
	li.Next = l.Front()
	l.Front().Prev = li
	return
}

func (l *list) PushBack(v interface{}) (li *ListItem) {
	li = &ListItem{
		Value: v,
		Next:  nil,
		Prev:  nil,
	}
	if l.Len() == 0 {
		l.ele = li
		return
	}
	li.Prev = l.Back()
	l.Back().Next = li
	return
}

func (l *list) Remove(i *ListItem) {
	defer func() {
		i.Next, i.Prev = nil, nil
	}()

	if i.Next == nil && i.Prev == nil {
		l.ele = nil
		return
	}

	if i.Next == nil {
		i.Prev.Next = nil
		l.ele = i.Prev
		return
	}

	if i.Prev == nil {
		i.Next.Prev = nil
		l.ele = i.Next
		return
	}

	i.Next.Prev = i.Prev
	i.Prev.Next = i.Next
	l.ele = i.Prev
}

func (l *list) MoveToFront(i *ListItem) {
	l.Remove(i)
	i.Next = nil
	i.Prev = nil

	if l.Len() == 0 {
		l.ele = i
		return
	}

	i.Next = l.Front()
	l.Front().Prev = i
}

func NewList() List {
	return new(list)
}
