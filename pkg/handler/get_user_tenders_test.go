package handler_test

import (
	"slices"
	"strings"
	test "tenders/pkg/tenderstest"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestGetUserTenders(t *testing.T) {
	require := require.New(t)

	// basic
	code, _ := test.Get("/tenders/my")
	require.Equal(401, code)

	code, _ = test.Get("/tenders/my?username=abc")
	require.Equal(401, code)

	code, _ = test.Get("/tenders/my?username=user01")
	require.Equal(200, code)

	haveAll := test.GetTenders(t, "/tenders/my?username=user01")
	require.Len(haveAll, 5)

	// offset
	first := haveAll[0]
	haveAll = test.GetTenders(t, "/tenders/my?username=user01&offset=1")
	require.NotEqual(first, haveAll[0])

	// limit
	haveAll = test.GetTenders(t, "/tenders/my?username=user01&limit=10")
	require.Len(haveAll, 10)

	// sorting
	for idx := 0; idx < len(haveAll)-1; idx++ {
		left := haveAll[idx]
		right := haveAll[idx+1]
		diff := strings.Compare(string(left.Name), string(right.Name))
		require.True(diff < 0)
	}

	// users of the same organization
	haveLeft := test.GetTenders(t, "/tenders/my?username=user01")
	haveRight := test.GetTenders(t, "/tenders/my?username=user02")
	for _, tender := range haveLeft {
		require.True(slices.Contains(haveRight, tender))
	}

	// users of different organizations
	haveRight = test.GetTenders(t, "/tenders/my?username=user04")
	for _, tender := range haveLeft {
		require.True(!slices.Contains(haveRight, tender))
	}
}
