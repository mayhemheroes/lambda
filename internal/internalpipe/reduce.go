package internalpipe

type AccumFn[T any] func(*T, *T) T

func Reduce[T any](dataFn func() []T, accum AccumFn[T]) *T {
	data := p.Do()
	switch len(data) {
	case 0:
		return nil
	case 1:
		return &data[0]
	default:
		res := data[0]
		for _, val := range data[1:] {
			res = accum(&res, &val)
		}
		return &res
	}
}
