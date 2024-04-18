package logger

import "context"

func With(ctx context.Context, key string, value any) context.Context {
	ctxAttrMu.Lock()
	defer ctxAttrMu.Unlock()

	if ctx.Value(CtxAttrKey{}) == nil {
		ctx = context.WithValue(ctx, CtxAttrKey{}, make([]any, 0))
	}

	kv, ok := ctx.Value(CtxAttrKey{}).([]any)
	if !ok {
		return ctx
	}

	return context.WithValue(ctx, CtxAttrKey{}, append(kv, key, value))
}

func AttrFromCtx(ctx context.Context) []any {
	ctxAttrMu.RLock()
	defer ctxAttrMu.RUnlock()

	kv, _ := ctx.Value(CtxAttrKey{}).([]any)

	return kv
}
