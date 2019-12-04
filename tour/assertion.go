package main

import "fmt"

type I interface {
	M()
}

type T struct {
	S string
}

func (t *T) M() {
	fmt.Println(t.S)
}

// type assertion provides access to an interface value's underlying concrete
// value: v := i.(T)
func main() {
	var i I
	var t = T{"hello"}
	i = &t

	var v = i.(*T) // will panic if i holds no value of type T
	fmt.Println(v) // &{hello}

	// will not panic if i holhs no value of type T
	// if ok is false, v will have zero value of T
	v, ok := i.(*T)
	fmt.Printf("(%v, %v)\n", v, ok) // (&{hello}, true)

	// useful when you are handling values of unknown types
	var s, f, it, b, aa interface{} = "hi", 123.45, 1212, true, [1]string{"haha"}

	// var vs = []interface{}{s, f, it, b}
	// PrintWithType(vs...)

	PrintWithType(s, f, it, b, aa)

	PrintWithTypeSwitch(s, f, it, b, aa)
}

func PrintWithType(ivs ...interface{}) {
	for _, i := range ivs {
		if s, ok := i.(string); ok {
			fmt.Printf("string: %v\n", s)
		} else if f, ok := i.(float64); ok {
			fmt.Printf("float64: %v\n", f)
		} else if it, ok := i.(int); ok {
			fmt.Printf("int: %v\n", it)
		} else if b, ok := i.(bool); ok {
			fmt.Printf("bool: %v\n", b)
		} else {
			fmt.Println("Unknown type")
		}
	}
}

func PrintWithTypeSwitch(ivs ...interface{}) {
	for _, i := range ivs {
		switch v := i.(type) { // type is keyword
		case string:
			fmt.Printf("string: %v\n", v)
		case float64:
			fmt.Printf("string: %v\n", v)
		case int:
			fmt.Printf("int: %v\n", v)
		case bool:
			fmt.Printf("bool: %v\n", v)
		default:
			fmt.Printf("Type %T is not supported", v)
		}
	}
}
