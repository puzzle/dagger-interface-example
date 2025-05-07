package main

import (
	"context"
	"fmt"
)

type Example struct{}

func (m *Example) Foo(ctx context.Context, bar int) (string, error) {
	return fmt.Sprintf("number is: %d", bar), nil
}

func (m *Example) HasBar() bool {
	return false
}
