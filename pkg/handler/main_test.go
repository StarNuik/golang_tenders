package handler_test

import (
	"os"
	test "tenders/pkg/tenderstest"
	"testing"
)

func TestMain(m *testing.M) {
	test.ClearDb()

	codes := make(chan int)
	reqs := test.TenderRequests()
	for _, req := range reqs {
		go func(codes chan<- int) {
			code, _ := test.Post("/tenders/new", req)
			codes <- code
		}(codes)
	}
	for range len(reqs) {
		code := <-codes
		if code != 200 {
			panic("TestMain: /tenders/new failed")
		}
	}

	code := m.Run()

	test.ClearDb()

	os.Exit(code)
}
