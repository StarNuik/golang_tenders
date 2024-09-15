package handler_test

import (
	"strings"
	"tenders/pkg/api"
	test "tenders/pkg/tenderstest"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestGetTenders(t *testing.T) {
	require := require.New(t)

	// basic
	code, _ := test.Get("/tenders")
	require.Equal(200, code)

	haveAll := test.GetTenders(t, "/tenders")
	require.Len(haveAll, 5)

	// offset
	first := haveAll[0]
	haveAll = test.GetTenders(t, "/tenders?offset=1")
	require.NotEqual(first, haveAll[0])

	// limit
	haveAll = test.GetTenders(t, "/tenders?limit=25")
	require.Len(haveAll, 25)

	// sorting
	for idx := 0; idx < len(haveAll)-1; idx++ {
		left := haveAll[idx]
		right := haveAll[idx+1]
		diff := strings.Compare(string(left.Name), string(right.Name))
		require.True(diff < 0)
	}

	// only published
	for _, tender := range haveAll {
		require.Equal(tender.Status, api.TenderStatusPublished)
	}

	// filters
	haveAll = test.GetTenders(t, "/tenders?service_type=Delivery&limit=50")
	for _, tender := range haveAll {
		require.Equal(tender.ServiceType, api.TenderServiceTypeDelivery)
	}
}
