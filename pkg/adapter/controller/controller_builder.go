package controller

type ControllerBuilder[O any] struct {
	controllers []Controller[O]
}

func NewControllerBuilder[O any]() *ControllerBuilder[O] {
	return &ControllerBuilder[O]{
		controllers: []Controller[O]{},
	}
}

func (b *ControllerBuilder[O]) Add(
	controller Controller[O],
) *ControllerBuilder[O] {
	b.controllers = append(b.controllers, controller)

	return b
}

func (b *ControllerBuilder[O]) Build() Controller[O] {
	cur := b.controllers[0]

	i := 1

	for i < len(b.controllers) {
		next := b.controllers[i]
		cur.SetNext(next)
		cur = next
		i += 1
	}

	return b.controllers[0]
}
