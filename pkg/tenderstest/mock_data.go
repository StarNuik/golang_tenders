package tenderstest

import (
	"strconv"
	"tenders/pkg/api"

	"github.com/google/uuid"
)

func TenderRequests() []*api.CreateTenderReq {
	out := []*api.CreateTenderReq{}

	idx := 0
	for _, row := range reps() {
		orgStr := row[0]
		for _, username := range row[1:] {
			for _, tenderStatus := range api.TenderStatusClosed.AllValues() {
				for _, tenderType := range api.TenderServiceTypeConstruction.AllValues() {
					req := baseReq()
					name := string(req.Name) + " " + strconv.Itoa(idx)

					req.Name = api.TenderName(name)
					req.OrganizationId = api.OrganizationId(uuid.MustParse(orgStr))
					req.CreatorUsername = api.Username(username)
					req.Status = tenderStatus
					req.ServiceType = tenderType

					out = append(out, &req)
					idx++
				}
			}
		}
	}

	return out
}

func TenderTypes() []string {
	out := []string{}
	for _, tenderType := range api.TenderServiceTypeConstruction.AllValues() {
		out = append(out, string(tenderType))
	}
	return out
}

func baseReq() api.CreateTenderReq {
	return api.CreateTenderReq{
		Name:            "Tender",
		Description:     "Tender Description",
		ServiceType:     "Construction",
		Status:          "Created",
		OrganizationId:  api.OrganizationId(uuid.MustParse("00000000-0000-0000-0000-000000000011")),
		CreatorUsername: "user01",
	}
}

// this was a map, but I actually
// need this to be in a reproducible order
func reps() [][]string {
	return [][]string{
		{"00000000-0000-0000-0000-000000000011", "user01", "user02", "user03"},
		{"00000000-0000-0000-0000-000000000012", "user04", "user05", "user06"},
		{"00000000-0000-0000-0000-000000000013", "user07", "user08", "user09"},
		{"00000000-0000-0000-0000-000000000014", "user10", "user11", "user12"},
	}
}
