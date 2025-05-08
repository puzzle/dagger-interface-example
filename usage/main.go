package main

import (
	"context"
)

type Usage struct{}

func (m *Usage) Test(ctx context.Context) (string, error) {
	// Because `Example` implements `Fooer`, the conversion function
	// `AsMyModuleFooer` has been generated.
	return dag.MyModule().Foo(ctx, dag.Example().AsMyModuleFooer())
}

func (m *Usage) HasBar(ctx context.Context) (bool, error) {
	return dag.MyModule().HasBar(ctx, dag.Example().AsMyModuleFooer())
}

func (m *Usage) Lint(dir *dagger.Directory,
	//+optional
	//+default=false
	pass bool) *dagger.Directory {
	return dag.MyModule().Lint(dir, pass, dag.Example().AsMyModuleFooer())
}
