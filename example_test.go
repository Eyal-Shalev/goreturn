package goreturn_test

import (
	"fmt"
	"os"
	"time"

	"github.com/Eyal-Shalev/goreturn"
	"github.com/dottedmag/must"
)

func Example() {
	file1 := must.OK1(os.CreateTemp("", "file1-*.txt"))
	defer os.Remove(file1.Name())
	_ = must.OK1(file1.WriteString("Hello, world!"))

	c1 := goreturn.Return2(func() (string, error) {
		data, err := os.ReadFile(file1.Name())
		return string(data), err
	})
	c2 := goreturn.Return2(func() (string, error) {
		data, err := os.ReadFile("nonexistent.txt")
		return string(data), err
	})

	file1Values := <-c1
	if file1Values.V2 != nil {
		fmt.Printf("file1 read error: %v", file1Values.V2)
		return
	}
	fmt.Println("file1:", file1Values.V1)

	file2Values := <-c2
	if file2Values.V2 != nil {
		fmt.Printf("file2 read error: %v", file2Values.V2)
		return
	}
	fmt.Println("file2:", file2Values.V1)

	// Output:
	// file1: Hello, world!
	// file2 read error: open nonexistent.txt: no such file or directory
}

func ExampleReturn0() {
	c := goreturn.Return0(func() {
		fmt.Println("Hello, world!")
	})
	select {
	case <-c:
		fmt.Println("Return0() called the function")
	case <-time.After(time.Millisecond):
		fmt.Println("Return0() took too long to return")
	}
	// Output:
	// Hello, world!
	// Return0() called the function
}

func ExampleReturn1() {
	c := goreturn.Return1(func() int {
		return 42
	})
	select {
	case v := <-c:
		fmt.Println("Return1() returned", v)
	case <-time.After(time.Millisecond):
		fmt.Println("Return1() took too long to return")
	}
	// Output:
	// Return1() returned 42
}

func ExampleReturn2() {
	c := goreturn.Return2(func() (int, string) {
		return 42, "hello"
	})
	select {
	case v := <-c:
		fmt.Println("Return2() returned", v.V1, v.V2)
	case <-time.After(time.Millisecond):
		fmt.Println("Return2() took too long to return")
	}
	// Output:
	// Return2() returned 42 hello
}

func ExampleReturn3() {
	c := goreturn.Return3(func() (int, string, bool) {
		return 42, "hello", true
	})
	select {
	case v := <-c:
		fmt.Println("Return3() returned", v.V1, v.V2, v.V3)
	case <-time.After(time.Millisecond):
		fmt.Println("Return3() took too long to return")
	}
	// Output:
	// Return3() returned 42 hello true
}
