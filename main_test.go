package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_SayHello(t *testing.T) {
	require.Equal(t, "hello world!", sayHello())
}
