package internalpipe

func Filter(i int) (*T, bool) {
	return func(i int) (*T, bool) {
		if obj, skipped := p.Fn(i); !skipped {
			if !fn(obj) {
				return nil, true
			}
			return obj, false
		}
		return nil, true
	}
}
