package internalpipe

func Map[T any](fn GeneratorFn[T], mapFn func(T) T) func(int) (*T, bool) {
	return func(i int) (*T, bool) {
		if obj, skipped := fn(i); !skipped {
			res := mapFn(*obj)
			return &res, false
		}
		return nil, true
	}
}
