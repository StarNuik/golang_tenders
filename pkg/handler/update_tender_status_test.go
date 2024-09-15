package handler_test

import (
	"tenders/pkg/api"
	"testing"

	test "tenders/pkg/tenderstest"

	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
)

func TestUpdateTenderStatus(t *testing.T) {
	require := require.New(t)

	closeTender := func(id string, user string) (int, string) {
		return test.Put("/tenders/" + id + "/status?username=" + user + "&status=Closed")
	}

	want := test.GetTenders(t, "/tenders/my?username=user01&limit=1&offset=4")[0]
	id := uuid.UUID(want.ID).String()

	code, _ := closeTender("abc", "user15")
	require.Equal(400, code)

	code, _ = closeTender(uuid.NewString(), "non-existent")
	require.Equal(404, code)

	code, _ = closeTender(id, "non-existent")
	require.Equal(401, code)

	code, _ = closeTender(id, "user15")
	require.Equal(403, code)

	// happy path
	code, resp := closeTender(id, "user03")
	require.Equal(200, code)

	have, err := test.ParseTender(resp)
	require.Nil(err)
	require.Equal(want.Name, have.Name)
	require.Equal(api.TenderStatusClosed, have.Status)
}
