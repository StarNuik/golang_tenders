package middleware

import (
	"log"
	"strconv"

	"github.com/ogen-go/ogen/middleware"
)

func Logging() middleware.Middleware {
	return func(
		req middleware.Request,
		next func(req middleware.Request) (middleware.Response, error),
	) (middleware.Response, error) {
		// log.Println("Handling request")

		resp, err := next(req)

		content := []string{}
		content = append(content, req.OperationName)

		if err != nil {
			content = append(content, "Error:", err.Error())
		} else {
			// Some response types may have a status code.
			// ogen provides a getter for it.
			//
			// You can write your own interface to match any response type.
			if tresp, ok := resp.Type.(interface{ GetStatusCode() int }); ok {
				str := strconv.Itoa(tresp.GetStatusCode())
				content = append(content, "Success", str)
			}
		}

		log.Println(content)

		return resp, err
	}
}
