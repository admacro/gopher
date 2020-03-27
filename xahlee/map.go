package main

import "fmt"

// map is an unordered collection of key value pairs
// similar to hashtable, dictionary, associative array in other languages
func main() {
	// map[key_type_name]value_type_name
	var m map[string]int  // declare a map with key of string and value of int, value of m is nil
	fmt.Printf("%T\n", m) // map[string]int
	fmt.Println(m)        // map[]
	fmt.Println(len(m))   // 0

	// the zero value of a map is nil
	// a nil map has no keys, nor can keys be added
	m["a"] = 1 // panic: assignment to entry in nil map

	// literal map
	var p = map[string]int{"dad": 50, "mom": 48}
	fmt.Printf("%#v\n", p) // map[string]int{"dad":50, "mom":48}
	fmt.Println(p)         // map[dad:50 mom:48]
	fmt.Println(len(p))    // 2

	// get value
	fmt.Println(p["mom"]) // 48

	// In fact, map[k] returns two values, but the second can be omitted if you don't need it
	// 1st is the associated value
	// 2nd is true if k exist, else false
	var val, exist = p["dad"]
	fmt.Printf("p[dad]:%v, %v\n", val, exist) // p[dad]:50, true

	// when a key doesn't exist, the first value is
	// the zero value of the type of value (0, "", nil, false)
	val, exist = p["grandpa"]
	fmt.Printf("p[grandpa]:%v, %v\n", val, exist) // p[grandpa]:0, false

	// make a map
	var r = make(map[rune]string)
	r['a'] = "AAAAA"       // add key value pair
	fmt.Printf("%#v\n", r) // map[int32]string{97:"AAAAA"}
	fmt.Println(r)         // map[97:AAAAA]
	r['a'] = "aaaaa"       // set value
	for k, v := range r {
		fmt.Printf("[%c: %v]\n", k, v) // [a: aaaaa]
	}

	var vvv, vvv_exist = r['b']
	fmt.Printf("r[b]:%#v, %#v\n", vvv, vvv_exist) // r[b]:"", false

	// delete key
	delete(p, "dad")
	fmt.Println(p) // map[mom:48]
}
