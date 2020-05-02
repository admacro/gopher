// https://golang.org/doc/effective_go.html#recover
package main

import "fmt"

// this func just panics if something wrong happens
// it's like functions that `throws exceptions` in Java
func panicFunc() {
	fmt.Println("in panicFunc")
	fmt.Println("panicFunc is panicking...")
	panic("panic in panicFunc")
}

// this func can recover itself if it panics
// it's like functions with a try-catch block in Java
func recoverFromPanicFunc() {
	// this resembles the catch block
	defer func() {
		if err := recover(); err != nil {
			fmt.Printf("recover from: [%v]\n", err)
		}
	}()

	// the rest is the try block
	fmt.Println("in recoverFromPanicFunc")
	panicFunc()
}

// this func tries to recover itself if it panics
// however, during the recovery, it may re-panic
// fortunately, it can recover from the re-panic
// it's like functions with nested try-catch blocks in Java
func recoverFromRepanicFunc() {
	defer func() {
		if err := recover(); err != nil {
			// this resembles a nested try-catch
			defer func() {
				if err := recover(); err != nil {
					fmt.Printf("recover from: [%v]\n", err)
				}
			}()
			fmt.Printf("tring to recover from: [%v]\n", err)
			fmt.Println("recoverFromRepanicFunc is panicking...")
			panic(fmt.Sprintf("[Fatal] Failure to cover from: [%v]", err))

			// the above can be put into a separate function
			// recoverFrom(err)
		}
	}()
	fmt.Println("in recoverFromRepanicFunc")
	panicFunc()
}

func recoverFrom(err interface{}) {
	defer func() {
		if err := recover(); err != nil {
			fmt.Printf("recover from: [%v]\n", err)
		}
	}()
	fmt.Printf("tring to recover from: [%v]\n", err)
	fmt.Println("the recovery is panicking...")
	panic(fmt.Sprintf("[Fatal] Failure to cover from: [%v]", err))
}

func main() {
	recoverFromPanicFunc()

	fmt.Println()

	recoverFromRepanicFunc()
}
