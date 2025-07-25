package types

type OneOf[T1, T2 any] struct {
	t1   T1
	t2   T2
	isT1 bool
	isT2 bool
}

func NewOneOf[T1, T2 any]() OneOf[T1, T2] {
	return OneOf[T1, T2]{}
}

func (o *OneOf[T1, T2]) Present() (bool, bool) {
	return o.isT1, o.isT2
}

func (o *OneOf[T1, T2]) GetT1() (T1, bool) {
	return o.t1, o.isT1
}

func (o *OneOf[T1, T2]) GetT2() (T2, bool) {
	return o.t2, o.isT2
}

func (o *OneOf[T1, T2]) SetT1(t1 T1) {
	o.t1 = t1
	var zero T2
	o.t2 = zero
	o.isT1 = true
	o.isT2 = false
}

func (o *OneOf[T1, T2]) SetT2(t2 T2) {
	o.t2 = t2
	var zero T1
	o.t1 = zero
	o.isT2 = true
	o.isT1 = false
}
