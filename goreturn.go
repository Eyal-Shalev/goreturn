package goreturn

func Return0(fn func()) <-chan any {
	c := make(chan any)
	go func() {
		fn()
		c <- nil
	}()
	return c
}

func Return1[T1 any](fn func() T1) <-chan T1 {
	c := make(chan T1)
	go func() {
		c <- fn()
	}()
	return c
}

type Tuple[T1 any, T2 any] struct {
	V1 T1
	V2 T2
}

func Return2[T1 any, T2 any](fn func() (T1, T2)) <-chan *Tuple[T1, T2] {
	c := make(chan *Tuple[T1, T2])
	go func() {
		v1, v2 := fn()
		c <- &Tuple[T1, T2]{v1, v2}
	}()
	return c
}

type Tuple3[T1 any, T2 any, T3 any] struct {
	V1 T1
	V2 T2
	V3 T3
}

func Return3[T1 any, T2 any, T3 any](fn func() (T1, T2, T3)) <-chan *Tuple3[T1, T2, T3] {
	c := make(chan *Tuple3[T1, T2, T3])
	go func() {
		v1, v2, v3 := fn()
		c <- &Tuple3[T1, T2, T3]{v1, v2, v3}
	}()
	return c
}
