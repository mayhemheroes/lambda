package internalpipe

type Predicate[T any] func(*T) bool

func Filter[T any](fn GeneratorFn[T], filter Predicate[T]) func(i int) (*T, bool) {
	return func(i int) (*T, bool) {
		if obj, skipped := fn(i); !skipped {
			if !filter(obj) {
				return nil, true
			}
			return obj, false
		}
		return nil, true
	}
}
