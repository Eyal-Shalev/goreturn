package goreturn

// Return0 runs a function in a goroutine and returns a channel that will receive a single `nil` value when the function completes.
func Return0(fn func()) <-chan any {
	c := make(chan any)
	go func() {
		fn()
		c <- nil
		close(c)
	}()
	return c
}

// Return1 runs a function with one return value in a goroutine and returns a channel that will receive the return value.
func Return1[T1 any](fn func() T1) <-chan T1 {
	c := make(chan T1)
	go func() {
		c <- fn()
		close(c)
	}()
	return c
}

// Tuple is a struct that holds two values of any type.
type Tuple[T1 any, T2 any] struct {
	V1 T1
	V2 T2
}

// Return2 runs a function with two return values in a goroutine and returns a channel that will receive a tuple containing the return values.
func Return2[T1 any, T2 any](fn func() (T1, T2)) <-chan *Tuple[T1, T2] {
	c := make(chan *Tuple[T1, T2])
	go func() {
		v1, v2 := fn()
		c <- &Tuple[T1, T2]{v1, v2}
		close(c)
	}()
	return c
}

// Tuple3 is a struct that holds three values of any type.
type Tuple3[T1 any, T2 any, T3 any] struct {
	V1 T1
	V2 T2
	V3 T3
}

// Return3 runs a function with three return values in a goroutine and returns a channel that will receive a tuple containing the return values.
func Return3[T1 any, T2 any, T3 any](fn func() (T1, T2, T3)) <-chan *Tuple3[T1, T2, T3] {
	c := make(chan *Tuple3[T1, T2, T3])
	go func() {
		v1, v2, v3 := fn()
		c <- &Tuple3[T1, T2, T3]{v1, v2, v3}
		close(c)
	}()
	return c
}
