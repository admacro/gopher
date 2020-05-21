package dtst

import (
	"testing"
)

func TestNew(t *testing.T) {
	list := New()
	if list.Length() != 0 || list.header != nil {
		t.Errorf("LinkedList not properly initialized: %v", list)
	}
}

func TestAppend(t *testing.T) {
	list := New()
	list.Append("hello")
	list.Append("world")
	if list.Length() != 2 || list.header.data != "hello" {
		t.Errorf("LinkedList Append failed: %v", list)
	}
}

func TestAppendAll(t *testing.T) {
	list := New()
	list.Append("Go")
	list.AppendAll("hello", 123)
	if list.Length() != 3 {
		t.Errorf("LinkedList AppendAll failed: %v", list)
	}
	list.AppendAll(true, 3.14)
	if list.Length() != 5 {
		t.Errorf("LinkedList AppendAll failed: %v", list)
	}
}

func TestPrepend(t *testing.T) {
	list := New()
	list.Prepend("hello")
	list.Prepend(123)
	if list.Length() != 2 || list.header.data != 123 {
		t.Errorf("LinkedList Prepend failed: %v", list)
	}
}

func TestClear(t *testing.T) {
	list := New()
	list.Append("hello")
	list.Append(123)
	list.Clear()
	if list.Length() != 0 || list.header != nil {
		t.Errorf("LinkedList Clear failed: %v", list)
	}
}

func TestDelete(t *testing.T) {
	list := New()
	list.Delete(123)
	if list.Length() != 0 {
		t.Errorf("LinkedList Delete failed: %v", list)
	}

	list.Append("hello")
	list.Delete("hello")
	if list.Length() != 0 {
		t.Errorf("LinkedList Delete failed: %v", list)
	}

	list.Append(123)
	list.Append(true)
	list.Delete(123)
	if list.Length() != 1 || list.header.data != true {
		t.Errorf("LinkedList Delete failed: %v", list)
	}

	list.Clear()
	list.Append("hello")
	list.Append(123)
	list.Append(true)
	list.Delete(123)
	if list.Length() != 2 || list.header.data != "hello" {
		t.Errorf("LinkedList Delete failed: %v", list)
	}
}

func TestGet(t *testing.T) {
	list := New()
	list.Append("hello")
	list.Append(123)
	list.Append(true)
	if list.Length() != 3 || list.Get(1) != 123 {
		t.Errorf("LinkedList Get failed: %v", list)
	}
}

func TestContains(t *testing.T) {
	list := New()
	list.Append("hello")
	list.Append(123)
	if list.Contains("hello") == false || list.Contains(123) == false {
		t.Errorf("LinkedList Contains failed: %v", list)
	}
}
