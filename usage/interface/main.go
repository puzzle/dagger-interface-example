package main

import (
	"context"
)

type MyModule struct{}

type Fooer interface {
	DaggerObject
	Foo(ctx context.Context, bar int) (string, error)
	// Generation fails if return type is not a core type
	HasBar() bool
}

func (m *MyModule) Foo(ctx context.Context, fooer Fooer) (string, error) {
	return fooer.Foo(ctx, 42)
}

func (m *MyModule) HasBar(fooer Fooer) bool {
    return fooer.HasBar()
}