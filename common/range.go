package common

// Only usable after Go1.23
// func Range[T Integer](l, r, s T) iter.Seq[T] {
// 	return func(yield func(T) bool) {
// 		for i := l; i < r; i += s {
// 			if !yield(i) {
// 				return
// 			}
// 		}
// 	}
// }
