package handler_test

import (
	"fmt"
	"testing"

	test "tenders/pkg/tenderstest"

	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
)

func TestEditTender(t *testing.T) {
	require := require.New(t)

	want := test.GetTenders(t, "/tenders/my?username=user01&limit=1&offset=4")[0]
	want.Description = "Edited"
	want.Version += 1

	editTender := func(id string, user string) (int, string) {
		url := fmt.Sprintf("/tenders/%s/edit?username=%s", id, user)
		req := struct {
			Description string `json:"description"`
		}{string(want.Description)}
		return test.Patch(url, req)
	}

	id := uuid.UUID(want.ID).String()

	code, resp := editTender(uuid.NewString(), "user15")
	require.Equal(404, code, resp)

	code, _ = editTender(id, "userXX")
	require.Equal(401, code)

	code, _ = editTender(id, "user15")
	require.Equal(403, code)

	// happy path
	code, resp = editTender(id, "user03")
	require.Equal(200, code)

	have, err := test.ParseTender(resp)
	require.Nil(err)
	require.Equal(want, have)

	code, resp = editTender(id, "user02")
	require.Equal(200, code)

	have, err = test.ParseTender(resp)
	require.Nil(err)
	require.Equal(want.Version+1, have.Version)
}
