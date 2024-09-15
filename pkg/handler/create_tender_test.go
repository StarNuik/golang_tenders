package handler_test

import (
	"encoding/json"
	"testing"

	"tenders/pkg/api"
	test "tenders/pkg/tenderstest"

	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
)

func TestCreateTender(t *testing.T) {
	require := require.New(t)

	baseReq := api.CreateTenderReq{
		Name:            "Tender Name",
		Description:     "Tender Description",
		ServiceType:     "Construction",
		Status:          "Created",
		OrganizationId:  api.OrganizationId(uuid.MustParse("00000000-0000-0000-0000-000000000011")),
		CreatorUsername: "user01",
	}

	call := func(req *api.CreateTenderReq) (int, string) {
		return test.Post("/tenders/new", req)
	}

	req := baseReq
	code, resp := call(&req)
	require.Equal(200, code, resp)

	// non-existent user
	req = baseReq
	req.CreatorUsername = "user42"
	code, resp = call(&req)
	require.Equal(401, code, resp)

	// not a rep of any org
	req = baseReq
	req.CreatorUsername = "user15"
	code, resp = call(&req)
	require.Equal(403, code, resp)

	// rep of a different org
	req = baseReq
	req.CreatorUsername = "user04"
	code, resp = call(&req)
	require.Equal(403, code, resp)

	// different org
	req = baseReq
	req.OrganizationId = api.OrganizationId(uuid.MustParse("00000000-0000-0000-0000-000000000012"))
	code, resp = call(&req)
	require.Equal(403, code, resp)

	// non-existent org
	req = baseReq
	req.OrganizationId = api.OrganizationId(uuid.MustParse("ffffffff-0000-0000-0000-000000000000"))
	code, resp = call(&req)
	require.Equal(403, code, resp)
}

func TestCreateTenderContents(t *testing.T) {
	require := require.New(t)

	want := api.CreateTenderReq{
		Name:            "Tender Name",
		Description:     "Tender Description",
		ServiceType:     "Construction",
		Status:          "Created",
		OrganizationId:  api.OrganizationId(uuid.MustParse("00000000-0000-0000-0000-000000000011")),
		CreatorUsername: "user01",
	}

	code, resp := test.Post("/tenders/new", want)
	require.Equal(200, code, resp)

	var have api.Tender
	err := json.Unmarshal([]byte(resp), &have)
	require.Nil(err, err)

	require.Equal(want.Name, have.Name)
	require.Equal(want.Description, have.Description)
	require.Equal(want.ServiceType, have.ServiceType)
	require.Equal(want.Status, have.Status)
	require.Equal(want.OrganizationId, have.OrganizationId)
	require.Equal(api.TenderVersion(1), have.Version)
	require.True(test.ApproxNow(have.CreatedAt))
}
