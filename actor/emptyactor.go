package actor

type EmptyActor struct {
	receive Receive
}

func (state *EmptyActor) Receive(context Context) {
	switch context.Message().(type) {
	case Started:
		context.Become(state.receive)
	}
}
