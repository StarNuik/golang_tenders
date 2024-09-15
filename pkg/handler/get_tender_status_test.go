package handler_test

import (
	"encoding/json"
	"strconv"
	"testing"

	"tenders/pkg/api"
	test "tenders/pkg/tenderstest"

	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
)

func TestGetTenderStatus(t *testing.T) {
	require := require.New(t)

	tenderId := func(offset int) string {
		offsetStr := strconv.Itoa(offset)
		tender := test.GetTenders(t, "/tenders/my?username=user01&limit=1&offset="+offsetStr)[0]
		return uuid.UUID(tender.ID).String()
	}

	created := tenderId(3)
	published := tenderId(4)
	closed := tenderId(7)

	tenderStatus := func(id string) (int, string) {
		return test.Get("/tenders/" + id + "/status")
	}

	// published tender
	want, _ := json.Marshal(string(api.TenderStatusPublished))
	code, resp := tenderStatus(published)
	require.Equal(200, code)
	require.Equal(string(want), resp)

	// forbidden
	code, resp = tenderStatus(created)
	require.Equal(403, code, resp)

	code, _ = tenderStatus(closed)
	require.Equal(403, code)

	// invalid uuid
	code, _ = tenderStatus("abc")
	require.Equal(400, code)

	// non-existent uuid
	code, _ = tenderStatus(uuid.New().String())
	require.Equal(404, code)

	testAuth := func(id string, status api.TenderStatus) {
		want, _ := json.Marshal(string(status))

		// correct users
		code, resp := test.Get("/tenders/" + id + "/status?username=user01")
		require.Equal(200, code)
		require.Equal(string(want), resp)

		code, resp = test.Get("/tenders/" + id + "/status?username=user03")
		require.Equal(200, code)
		require.Equal(string(want), resp)

		// incorrect user
		code, _ = test.Get("/tenders/" + id + "/status?username=user04")
		require.Equal(403, code)

		// non-existent user
		code, _ = test.Get("/tenders/" + id + "/status?username=doesntexist")
		require.Equal(401, code)
	}

	testAuth(created, api.TenderStatusCreated)
	testAuth(closed, api.TenderStatusClosed)
}
