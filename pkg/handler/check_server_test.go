package handler_test

import (
	"testing"

	test "tenders/pkg/tenderstest"

	"github.com/stretchr/testify/require"
)

func TestCheckServer(t *testing.T) {
	require := require.New(t)

	code, resp := test.Get("/ping")
	require.Equal(200, code)
	require.Equal("ok", resp)
}
