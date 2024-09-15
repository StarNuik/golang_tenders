package tenderstest

import (
	"encoding/json"
	"tenders/pkg/api"
	"time"
)

func ParseTender(s string) (api.Tender, error) {
	var have api.Tender
	err := json.Unmarshal([]byte(s), &have)
	return have, err
}

func ApproxNow(stamp time.Time) bool {
	return time.Now().UTC().Sub(stamp).Seconds() <= 2
}

func panicIf(err error) {
	if err != nil {
		panic(err)
	}
}
