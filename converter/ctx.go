package converter

import (
	"context"

	"golang.org/x/net/html"
)

// func GetValue[K string, V any](ctx context.Context, key K) V {
// 	val, _ := ctx.Value(key).(V)
// 	return val
// }
// func SetValue[K string, V any](ctx context.Context, key K, val V) context.Context {
// 	return context.WithValue(ctx, key, val)
// }

type ctxKey string

const (
	ctxKeyAssembleAbsoluteURL ctxKey = "AssembleAbsoluteURL"
	ctxKeyDomain              ctxKey = "Domain"

	ctxKeySetState    ctxKey = "SetState"
	ctxKeyUpdateState ctxKey = "UpdateState"
	ctxKeyGetState    ctxKey = "GetState"
)

func provideDomain(ctx context.Context, domain string) context.Context {
	_ = "STUB: not implemented"
	return *new(context.Context)
}

func GetDomain(ctx context.Context) string { _ = "STUB: not implemented"; return "" }

// - - - - - - - - - - - - - - - - - - - - - //

type AssembleAbsoluteURLFunc func(tagName string, rawURL string, domain string) string

func assembleAbsoluteURL(ctx context.Context, tagName string, rawURL string) string {
	_ = "STUB: not implemented"
	return ""

	// TODO: since this gets passed down from the converter, it doesn't have to provided from the ctx anymore
}

func provideAssembleAbsoluteURL(ctx context.Context, fn AssembleAbsoluteURLFunc) context.Context {
	_ = "STUB: not implemented"
	return *new(context.Context)
}

// - - - - - - - - - - - - - - - - - - - - - //

type SetStateFunc func(key string, val any)
type UpdateStateFunc func(key string, fn func(any) any)
type GetStateFunc func(key string) any

type globalState struct {
	data map[string]any
}

func newGlobalState() *globalState { _ = "STUB: not implemented"; return nil }

func (s *globalState) setState(key string, val any) { _ = "STUB: not implemented"; return }

func (s *globalState) updateState(key string, fn func(any) any) { _ = "STUB: not implemented"; return }

func (s *globalState) getState(key string) any { _ = "STUB: not implemented"; return *new(any) }

func (s *globalState) provideGlobalState(ctx context.Context) context.Context {
	_ = "STUB: not implemented"
	return *new(context.Context)
}

func GetState[V any](ctx context.Context, key string) V { _ = "STUB: not implemented"; return *new(V) }

func SetState[V any](ctx context.Context, key string, val V) { _ = "STUB: not implemented"; return }

func UpdateState[V any](ctx context.Context, key string, fn func(V) V) {
	_ = "STUB: not implemented"
	return
}

// TODO: slog?

// - - - - - - //

// Context extends the normal context.Context with some additional
// methods useful for the process of converting.
type Context interface {
	context.Context

	AssembleAbsoluteURL(ctx Context, tagName string, rawURL string) string

	GetTagType(tagName string) (tagType, bool)

	RenderNodes(ctx Context, w Writer, nodes ...*html.Node)
	RenderChildNodes(ctx Context, w Writer, n *html.Node)

	EscapeContent(content []byte) []byte
	UnEscapeContent(content []byte) []byte

	WithValue(key any, val any) Context
}

type converterContext struct {
	context.Context
	conv *Converter
}

func newConverterContext(ctx context.Context, conv *Converter) Context {
	_ = "STUB: not implemented"
	return *new(Context)
}

func (c *converterContext) AssembleAbsoluteURL(ctx Context, tagName string, rawURL string) string {
	_ = "STUB: not implemented"
	return ""
}

func (c *converterContext) RenderNodes(ctx Context, w Writer, nodes ...*html.Node) {
	_ = "STUB: not implemented"
	return
}

func (c *converterContext) RenderChildNodes(ctx Context, w Writer, n *html.Node) {
	_ = "STUB: not implemented"
	return
}

func (c *converterContext) GetTagType(tagName string) (tagType, bool) {
	_ = "STUB: not implemented"
	return *new(tagType), false
}

func (c *converterContext) EscapeContent(content []byte) []byte {
	_ = "STUB: not implemented"
	return nil
}

func (c *converterContext) UnEscapeContent(content []byte) []byte {
	_ = "STUB: not implemented"
	return nil
}

func (c *converterContext) WithValue(key any, val any) Context {
	_ = "STUB: not implemented"
	return *new(Context)
}
