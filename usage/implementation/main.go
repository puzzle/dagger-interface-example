package main

import (
	"context"
	"dagger/example/internal/dagger"
	"fmt"
)

type Example struct{}

func (m *Example) Foo(ctx context.Context, bar int) (string, error) {
	return fmt.Sprintf("number is: %d", bar), nil
}

func (m *Example) HasBar() (bool, error) {
	return false, nil
}

func (m *Example) Lint(dir *dagger.Directory, pass bool) *dagger.Directory {
	var anotherDir *dagger.Directory
	return anotherDir
}
