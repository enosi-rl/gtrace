// Code generated by gtrace. DO NOT EDIT.

// +build !gtrace

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

func gtraceNoopBdbd803d() {
}
func (AnotherBuildTagTrace) onSomethingA(context.Context) func() {
	return gtraceNoopBdbd803d
}
func gtraceNoop86ffa8ac(int32, int64) {
}
func (AnotherBuildTagTrace) onSomethingB(context.Context, int8, int16) func(int32, int64) {
	return gtraceNoop86ffa8ac
}
func gtraceNoop8a9c608a(Type) {
}
func (AnotherBuildTagTrace) onSomethingC(context.Context, Type) func(Type) {
	return gtraceNoop8a9c608a
}
func gtraceNoopBdbd803d1() {
}
func anotherBuildTagTraceOnSomethingA(context.Context, AnotherBuildTagTrace) func() {
	return gtraceNoopBdbd803d1
}
func gtraceNoop86ffa8ac1(int32, int64) {
}
func anotherBuildTagTraceOnSomethingB(context.Context, AnotherBuildTagTrace, int8, int16) func(int32, int64) {
	return gtraceNoop86ffa8ac1
}
func gtraceNoop8a9c608a1(string, int, bool, error) {
}
func anotherBuildTagTraceOnSomethingC(context.Context, AnotherBuildTagTrace, string, int, bool, error) func(string, int, bool, error) {
	return gtraceNoop8a9c608a1
}