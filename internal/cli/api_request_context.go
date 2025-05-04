package cli

type ApiRequestContext struct {
	previous string
	next     string
}

func (ctx *ApiRequestContext) update(previous, next string) {
	ctx.next = next
	ctx.previous = previous
}
