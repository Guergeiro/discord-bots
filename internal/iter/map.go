package iter

import "iter"

// From https://github.com/DeedleFake/xiter/blob/master/transform.go

func Map[T1, T2 any](seq iter.Seq[T1], f func(T1) T2) iter.Seq[T2] {
	return func(yield func(T2) bool) {
		seq(func(v T1) bool {
			return yield(f(v))
		})
	}
}
