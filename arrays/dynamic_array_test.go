package arrays

import (
	"testing"
)

func TestDynamicArray(t *testing.T) {
	da := NewSlice(0, 16)
	t.Log(da.Size())
	t.Log(da.Capacity())
	t.Log(da.IsEmpty())
	// t.Log(da.At(2))
	da.Push(2)
	t.Log(da.Size())
	t.Log(da.Capacity())
	t.Log(da.IsEmpty())
	t.Log(da.At(0))
	t.Log(da.GetArray())
	for i := 1; i < 16; i++ {
		da.Push(i)
	}
	da.Insert(10, 101)
	t.Log(da.GetArray())
	t.Log(da.Pop())
	da.Delete(10)
	t.Log(da.GetArray())
	for i := 1; i < 4; i++ {
		da.Push(1)
	}
	t.Log(da.Remove(1))
	t.Log(da.GetArray())
	t.Log(da.Find(100000))
}
