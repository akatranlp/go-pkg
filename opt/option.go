package opt

type Optional[T any] struct {
	ok    bool
	value T
}

func None[T any]() Optional[T] {
	return Optional[T]{}
}

func Some[T any](v T) Optional[T] {
	return Optional[T]{
		ok:    true,
		value: v,
	}
}

func Of[T any](v T, ok bool) Optional[T] {
	return Optional[T]{
		ok:    ok,
		value: v,
	}
}

func OfNillable[T any](v *T) Optional[T] {
	if v == nil {
		return Optional[T]{}
	}
	return Optional[T]{
		ok:    true,
		value: *v,
	}
}

func OfNillablePtr[T any](v *T) Optional[*T] {
	if v == nil {
		return Optional[*T]{}
	}
	return Optional[*T]{
		ok:    true,
		value: v,
	}
}
