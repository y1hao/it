package it

func NaturalsCh() <-chan int {
	ch := make(chan int)
	go func() {
		var i int
		for {
			ch <- i
			i++
		}
	}()
	return ch
}

func TakeCh[T any](n int, ch <-chan T) <-chan T {
	out := make(chan T)
	if n <= 0 {
		return out
	}
	go func() {
		var i int
		for x := range ch {
			if i == n {
				close(out)
				return
			}
			out <- x
			i++
		}
	}()
	return out
}

func MapCh[T1, T2 any](f func(T1) T2, ch <-chan T1) <-chan T2 {
	out := make(chan T2)
	go func() {
		for x := range ch {
			out <- f(x)
		}
		close(out)
	}()
	return out
}

func FilterCh[T any](predicate func(T) bool, ch <-chan T) chan T {
	out := make(chan T)
	go func() {
		for x := range ch {
			if predicate(x) {
				out <- x
			}
		}
		close(out)
	}()
	return out
}

func ReduceCh[T1, T2 any](base T2, f func(T1, T2) T2, ch <-chan T1) T2 {
	for x := range ch {
		base = f(x, base)
	}
	return base
}
