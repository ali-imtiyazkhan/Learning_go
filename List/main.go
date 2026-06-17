package main

import "fmt"

type DynamicArray struct {
	data   []int
	length int
}

func NewDynamicArray() *DynamicArray {
	return &DynamicArray{data: make([]int, 0), length: 0}
}

func (d *DynamicArray) Append(val int) {
	d.data = append(d.data, val)
	d.length++
}

func (d *DynamicArray) Get(index int) int {
	return d.data[index]
}

func (d *DynamicArray) Set(index, val int) {
	d.data[index] = val
}

func (d *DynamicArray) Remove(index int) {
	d.data = append(d.data[:index], d.data[index+1:]...)
	d.length--
}

func (d *DynamicArray) Len() int {
	return d.length
}

func (d *DynamicArray) Print() {
	fmt.Println(d.data)
}

func main() {
	arr := NewDynamicArray()
	arr.Append(10)
	arr.Append(20)
	arr.Append(30)
	arr.Print()

	arr.Set(1, 25)
	arr.Print()

	arr.Remove(0)
	arr.Print()

	fmt.Println("Length:", arr.Len())
	fmt.Println("Element at 0:", arr.Get(0))
}