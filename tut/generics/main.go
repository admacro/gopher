// https://go.dev/doc/tutorial/generics
package main

import "fmt"

func SumInts(m map[string]int64) int64 {
	var s int64
	for _, v := range m {
		s += v
	}
	return s
}

func SumFloats(m map[string]float64) float64 {
	var s float64
	for _, v := range m {
		s += v
	}
	return s
}

func Sum[K string, V int64 | float64](m map[K]V) V {
	var s V
	for _, val := range m {
		s += val
	}
	return s
}

// K and V are "type parameters"
// "comparable" and "int64 | float64" are "type contraint"
//
// the comparable constraint
// Intended specifically for cases like these, the comparable constraint is
// predeclared in Go. It allows any type whose values may be used as an
// operand of the comparison operators == and !=. Go requires that map
// keys be comparable. So declaring K as comparable is necessary so you
// can use K as the key in the map variable. It also ensures that calling
// code uses an allowable type for map keys.
//
// union contraint
// Contraint for V is a union of two types: int64 and float64. Using | specifies
// a union of the two types, meaning that this constraint allows either type.
// Either type will be permitted by the compiler as an argument in the calling code.
//
// map[K]V
// Note that we know map[K]V is a valid map type because K is a comparable type.
// If we hadn’t declared K comparable, the compiler would reject the reference to map[K]V.
// You can use any particular comparable type for the constraint of keys of a map, e.g. string.
func SumIntsOrFloats[K comparable, V int64 | float64](m map[K]V) V {
	var s V
	for _, val := range m {
		s += val
	}
	return s
}

type Number interface {
	int64 | float64
}

func SumNumbers[K comparable, V Number](m map[K]V) V {
	var s V
	for _, val := range m {
		s += val
	}
	return s
}

func main() {
	fmt.Println("Non-generic sum")
	intMap := map[string]int64{"a": 5, "b": 8, "c": 2, "d": 9}
	fmt.Println("SumInts(intMap):", SumInts(intMap))
	floatMap := map[string]float64{"a": 3.44, "b": 4.75, "c": 6.22, "d": 7.28}
	fmt.Println("SumFloats(floatMap):", SumFloats(floatMap))

	fmt.Println("\nGeneric sum")
	fmt.Println("Generic func with type params: [K string, V int64|float64]")
	fmt.Println("Call with type args")
	// In cases when type arguments are unnecessary, compiler will tell you
	fmt.Println("Sum[string, int64](intMap):", Sum[string, int64](intMap))
	fmt.Println("Sum[string, float64](floatMap):", Sum[string, float64](floatMap))

	// You can OFTEN omit the type arguments in the function call, when the
	// Go compiler can infer the types from the types of function arguments.
	// Note that this isn’t always possible. For example, if you needed to call
	// a generic function that had no arguments, you would need to include the
	// type arguments in the function call.
	fmt.Println("Call without type args (type parameters inferred from func args)")
	fmt.Println("Sum(intMap):", Sum(intMap))
	fmt.Println("Sum(floatMap):", Sum(floatMap))

	// While a type parameter’s constraint typically represents a set of types,
	// at compile time the type parameter stands for a single type – the type
	// provided as a type argument by the calling code. If the type argument’s
	// type isn’t allowed by the type parameter’s constraint, the code won’t compile.
	//
	// float32Map := map[string]float32{"a": 3.44, "b": 4.75, "c": 6.22, "d": 7.28}
	// fmt.Println("Sum(float32Map):", Sum(float32Map)) // compile error: float32 does not implement int64|float64

	// Note here type arguments are concrete types (string), not the comparable interface
	// You can not call SumIntsOrFloats[comparable, int64](intMap)
	fmt.Println("\nGeneric func with type params: [K comparable, V int64|float64]")
	fmt.Println("SumIntsOrFloats(intMap):", SumIntsOrFloats(intMap))
	fmt.Println("SumIntsOrFloats(floatMap):", SumIntsOrFloats(floatMap))

	fmt.Println("\nGeneric func with type params: [K comparable, V Number]")
	fmt.Println("SumNumbers(intMap):", SumNumbers(intMap))
	fmt.Println("SumNumbers(floatMap):", SumNumbers(floatMap))
}
