package linkedlist

import "testing"

func TestInsert(t *testing.T) {
	nums := []int{10, 20, 30}
	var ll List
	for _, i := range nums {
		ll.Insert(i)
	}
	got := ll.len
	want := 3
	if got != want {
		t.Errorf("Got %d, want %d", got, want)
	}
}

func TestGet_Tail(t *testing.T) {
	nums := []int{10, 20, 30, 40}
	var ll List
	for _, i := range nums {
		ll.Insert(i)
	}
	got := 10
	want, _ := ll.Get_Tail()
	if got != want {
		t.Errorf("Got %d, want %d", got, want)
	}
}

func TestAppend(t *testing.T) {
	nums := []int{10, 20, 30, 40}
	var ll List
	for _, i := range nums {
		ll.Append(i)
	}
	got := 40
	want, _ := ll.Get_Tail()
	if got != want {
		t.Errorf("Got %d, want %d", got, want)
	}
}
