package textlist_test

import (
	"github.com/gonyyi/textlist"
	"testing"
)

func TestList(t *testing.T) {
	c := textlist.NewList()
	c.Add("test", "test123", "test", "test234")
	d := textlist.NewList("a", "a1", "test")

	out1, out2 := textlist.Compare(c,d)

	for _, v := range out1 {
		println("added:", v)
	}
	for _, v := range out2 {
		println("removed:", v)
	}
}

func TestList_Reset(t *testing.T) {
	c := textlist.NewList("a","b","c")
	println(c.Size(), "exp 3")
	c.Add("d")
	println(c.Size(), "exp 4")
	c.Add("b")
	println(c.Size(), "exp 4")
	c.Reset()
	println(c.Size(), "exp 0")
}
