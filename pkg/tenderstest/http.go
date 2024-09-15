package tenderstest

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"net/http"
	"os"
	"tenders/pkg/api"
	"testing"

	"github.com/jackc/pgx/v5"
	"github.com/stretchr/testify/require"
)

var baseUrl = "http://localhost:8080/api"
var pgUrl = "postgres://user:insecure@localhost:5432"

func init() {
	if url, ok := os.LookupEnv("SERVER_URL"); ok {
		baseUrl = url + "/api"
	}
	if url, ok := os.LookupEnv("DB_URL"); ok {
		pgUrl = url
	}
}

func ClearDb() {
	ctx := context.Background()

	db, err := pgx.Connect(ctx, pgUrl)
	panicIf(err)
	defer db.Close(ctx)

	_, err = db.Exec(ctx, `TRUNCATE tender_content_ref, tender_content, tender`)
	panicIf(err)
}

func GetTenders(t *testing.T, url string) []api.Tender {
	require := require.New(t)

	var out []api.Tender
	code, resp := Get(url)
	require.Equal(200, code, resp)

	err := json.Unmarshal([]byte(resp), &out)
	require.Nil(err, err)
	return out
}

func Post(url string, body any) (int, string) {
	return callJson("POST", url, body)
}

func Patch(url string, body any) (int, string) {
	return callJson("PATCH", url, body)
}

func Get(url string) (int, string) {
	return call("GET", url, "", nil)
}

func Put(url string) (int, string) {
	return call("PUT", url, "", nil)
}

func callJson(method string, url string, body any) (int, string) {
	json, err := json.Marshal(body)
	panicIf(err)

	headers := make(map[string]string)
	headers["Content-Type"] = "application/json"

	return call(method, url, string(json), headers)
}

func call(method string, url string, body string, headers map[string]string) (int, string) {
	bytes := bytes.NewBufferString(body)
	req, err := http.NewRequest(method, baseUrl+url, bytes)
	panicIf(err)

	for key, value := range headers {
		req.Header.Set(key, value)
	}

	resp, err := http.DefaultClient.Do(req)
	panicIf(err)
	defer resp.Body.Close()

	out, err := io.ReadAll(resp.Body)
	panicIf(err)

	return resp.StatusCode, string(out)
}
