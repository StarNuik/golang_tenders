package handler

import (
	"fmt"
	"strconv"
	test "tenders/pkg/tenderstest"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
)

func TestRollbackTender(t *testing.T) {
	require := require.New(t)

	// setup
	tenderFrom := test.GetTenders(t, "/tenders/my?username=user01&offset=26&limit=1")[0]

	url := fmt.Sprintf("/tenders/%s/edit?username=user01", uuid.UUID(tenderFrom.ID).String())
	req := struct {
		Description string `json:"description"`
	}{}

	req.Description = "Edited"
	code, _ := test.Patch(url, req)
	require.Equal(200, code)

	req.Description = "Edited Twice"
	test.Patch(url, req)
	require.Equal(200, code)

	//
	rollback := func(id string, version int, user string) (int, string) {
		url := fmt.Sprintf("/tenders/%s/rollback/%s?username=%s",
			id, strconv.Itoa(version), user)
		return test.Put(url)
	}

	id := uuid.UUID(tenderFrom.ID).String()

	code, _ = rollback(uuid.NewString(), 2, "user15")
	require.Equal(404, code)

	code, _ = rollback(id, 100, "user02")
	require.Equal(404, code)

	code, _ = rollback(id, 2, "userXX")
	require.Equal(401, code)

	code, _ = rollback(id, 2, "user15")
	require.Equal(403, code)

	// happy path
	code, resp := rollback(id, 2, "user02")
	require.Equal(200, code)

	have, err := test.ParseTender(resp)
	require.Nil(err)
	require.Equal("Edited", string(have.Description))
	require.Equal(4, int(have.Version))

	code, resp = rollback(id, 1, "user02")
	require.Equal(200, code)

	have, err = test.ParseTender(resp)
	require.Nil(err)
	require.Equal(tenderFrom.Description, have.Description)
	require.Equal(5, int(have.Version))
}
