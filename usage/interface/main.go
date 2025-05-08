package main

import (
	"context"
	"dagger/my-module/internal/dagger"
)

type MyModule struct{}

type Fooer interface {
	dagger.DaggerObject
	Foo(ctx context.Context, bar int) (string, error)
	HasBar(ctx context.Context) (bool, error)
	Lint(dir *dagger.Directory, pass bool) *dagger.Directory
}

func (m *MyModule) Foo(ctx context.Context, fooer Fooer) (string, error) {
	return fooer.Foo(ctx, 42)
}

func (m *MyModule) HasBar(ctx context.Context, fooer Fooer) (bool, error) {
	return fooer.HasBar(ctx)
}

func (m *MyModule) Lint(
	dir *dagger.Directory,
	pass bool,
	fooer Fooer,
) *dagger.Directory {
	return fooer.Lint(dir, pass)
}
