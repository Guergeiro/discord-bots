package controller

type ControllerBuilder struct {
	controllers []Controller
}

func NewControllerBuilder() *ControllerBuilder {
	return &ControllerBuilder{
		controllers: []Controller{},
	}
}

func (b *ControllerBuilder) Add(
	controller Controller,
) *ControllerBuilder {
	b.controllers = append(b.controllers, controller)

	return b
}

func (b *ControllerBuilder) Build() Controller {
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
