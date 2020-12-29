// Code generated by gtrace. DO NOT EDIT.

// +build gtrace

package test

import (
	"context"
)

// Compose returns a new AnotherBuildTagTrace which has functional fields composed
// both from t and x.
func (t AnotherBuildTagTrace) Compose(x AnotherBuildTagTrace) (ret AnotherBuildTagTrace) {
	switch {
	case t.OnSomethingA == nil:
		ret.OnSomethingA = x.OnSomethingA
	case x.OnSomethingA == nil:
		ret.OnSomethingA = t.OnSomethingA
	default:
		h1 := t.OnSomethingA
		h2 := x.OnSomethingA
		ret.OnSomethingA = func() func() {
			r1 := h1()
			r2 := h2()
			switch {
			case r1 == nil:
				return r2
			case r2 == nil:
				return r1
			default:
				return func() {
					r1()
					r2()
				}
			}
		}
	}
	switch {
	case t.OnSomethingB == nil:
		ret.OnSomethingB = x.OnSomethingB
	case x.OnSomethingB == nil:
		ret.OnSomethingB = t.OnSomethingB
	default:
		h1 := t.OnSomethingB
		h2 := x.OnSomethingB
		ret.OnSomethingB = func(in0 int8, in1 int16) func(int32, int64) {
			r1 := h1(in0, in1)
			r2 := h2(in0, in1)
			switch {
			case r1 == nil:
				return r2
			case r2 == nil:
				return r1
			default:
				return func(in0 int32, in1 int64) {
					r1(in0, in1)
					r2(in0, in1)
				}
			}
		}
	}
	switch {
	case t.OnSomethingC == nil:
		ret.OnSomethingC = x.OnSomethingC
	case x.OnSomethingC == nil:
		ret.OnSomethingC = t.OnSomethingC
	default:
		h1 := t.OnSomethingC
		h2 := x.OnSomethingC
		ret.OnSomethingC = func(in0 Type) func(Type) {
			r1 := h1(in0)
			r2 := h2(in0)
			switch {
			case r1 == nil:
				return r2
			case r2 == nil:
				return r1
			default:
				return func(in0 Type) {
					r1(in0)
					r2(in0)
				}
			}
		}
	}
	return ret
}

type anotherBuildTagTraceContextKey struct{}

// WithAnotherBuildTagTrace returns context which has associated AnotherBuildTagTrace with it.
func WithAnotherBuildTagTrace(ctx context.Context, t AnotherBuildTagTrace) context.Context {
	return context.WithValue(ctx,
		anotherBuildTagTraceContextKey{},
		ContextAnotherBuildTagTrace(ctx).Compose(t),
	)
}

// ContextAnotherBuildTagTrace returns AnotherBuildTagTrace associated with ctx.
// If there is no AnotherBuildTagTrace associated with ctx then zero value 
// of AnotherBuildTagTrace is returned.
func ContextAnotherBuildTagTrace(ctx context.Context) AnotherBuildTagTrace {
	t, _ := ctx.Value(anotherBuildTagTraceContextKey{}).(AnotherBuildTagTrace)
	return t
}

func (t AnotherBuildTagTrace) onSomethingA(ctx context.Context) func() {
	c := ContextAnotherBuildTagTrace(ctx)
	var fn func() func() 
	switch {
	case t.OnSomethingA == nil:
		fn = c.OnSomethingA
	case c.OnSomethingA == nil:
		fn = t.OnSomethingA
	default:
		h1 := t.OnSomethingA
		h2 := c.OnSomethingA
		fn = func() func() {
			r1 := h1()
			r2 := h2()
			switch {
			case r1 == nil:
				return r2
			case r2 == nil:
				return r1
			default:
				return func() {
					r1()
					r2()
				}
			}
		}
	}
	if fn == nil {
		return func() {
			return
		}
	}
	res := fn()
	if res == nil {
		return func() {
			return
		}
	}
	return res
}
func (t AnotherBuildTagTrace) onSomethingB(ctx context.Context, in0 int8, in1 int16) func(int32, int64) {
	c := ContextAnotherBuildTagTrace(ctx)
	var fn func(int8, int16) func(int32, int64) 
	switch {
	case t.OnSomethingB == nil:
		fn = c.OnSomethingB
	case c.OnSomethingB == nil:
		fn = t.OnSomethingB
	default:
		h1 := t.OnSomethingB
		h2 := c.OnSomethingB
		fn = func(in0 int8, in1 int16) func(int32, int64) {
			r1 := h1(in0, in1)
			r2 := h2(in0, in1)
			switch {
			case r1 == nil:
				return r2
			case r2 == nil:
				return r1
			default:
				return func(in0 int32, in1 int64) {
					r1(in0, in1)
					r2(in0, in1)
				}
			}
		}
	}
	if fn == nil {
		return func(int32, int64) {
			return
		}
	}
	res := fn(in0, in1)
	if res == nil {
		return func(int32, int64) {
			return
		}
	}
	return res
}
func (t AnotherBuildTagTrace) onSomethingC(ctx context.Context, in0 Type) func(Type) {
	c := ContextAnotherBuildTagTrace(ctx)
	var fn func(Type) func(Type) 
	switch {
	case t.OnSomethingC == nil:
		fn = c.OnSomethingC
	case c.OnSomethingC == nil:
		fn = t.OnSomethingC
	default:
		h1 := t.OnSomethingC
		h2 := c.OnSomethingC
		fn = func(in0 Type) func(Type) {
			r1 := h1(in0)
			r2 := h2(in0)
			switch {
			case r1 == nil:
				return r2
			case r2 == nil:
				return r1
			default:
				return func(in0 Type) {
					r1(in0)
					r2(in0)
				}
			}
		}
	}
	if fn == nil {
		return func(Type) {
			return
		}
	}
	res := fn(in0)
	if res == nil {
		return func(Type) {
			return
		}
	}
	return res
}
func anotherBuildTagTraceOnSomethingA(ctx context.Context, t AnotherBuildTagTrace) func() {
	res := t.onSomethingA(ctx)
	return func() {
		res()
	}
}
func anotherBuildTagTraceOnSomethingB(ctx context.Context, t AnotherBuildTagTrace, i int8, i1 int16) func(int32, int64) {
	res := t.onSomethingB(ctx, i, i1)
	return func(i2 int32, i3 int64) {
		res(i2, i3)
	}
}
func anotherBuildTagTraceOnSomethingC(ctx context.Context, t AnotherBuildTagTrace, s string, i int, b bool, e error) func(string, int, bool, error) {
	var p Type
	p.String = s
	p.Integer = i
	p.Boolean = b
	p.Error = e
	res := t.onSomethingC(ctx, p)
	return func(s1 string, i1 int, b1 bool, e1 error) {
		var p Type
		p.String = s1
		p.Integer = i1
		p.Boolean = b1
		p.Error = e1
		res(p)
	}
}