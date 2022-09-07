package core

import (
	"context"
	"net/http"
	"strconv"
)

// Midware is simple white magic
// for wrapping functional options
type Midware func(http.Handler) http.Handler

func Wrap(hdl http.Handler, midware ...Midware) http.Handler {
	for _, mid := range midware {
		hdl = mid(hdl)
	}
	return hdl
}

// CheckMethod will check if method is GET and
// if it is return 405 error
func CheckMethod(hdl http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		if req.Method == http.MethodPost {
			http.Error(rw, "wrong http method", http.StatusMethodNotAllowed)
			return
		}
		hdl.ServeHTTP(rw, req)
	})
}

// ValidateParams check if all necessary parameters for
// main random logic are present in query
// In case all parameters are OK
// they will convert to int and send to next layer via context
// (not sure if it is a good choice though)
func ValidateParams(params ...string) Midware {
	return func(hdl http.Handler) http.Handler {
		return http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
			ctx := context.Background()

			q := req.URL.Query()
			for _, param := range params {
				val := q.Get(param)
				if val == "" {
					http.Error(rw, "missing parameter: "+param, http.StatusBadRequest)
					return // return error is at least one param is missing
				}

				// TODO: Not sure about that
				num, err := strconv.Atoi(val)
				if err != nil {
					http.Error(rw, "error parsing param: "+param, http.StatusInternalServerError)
					return
				}
				ctx = context.WithValue(ctx, param, num)

			}

			hdl.ServeHTTP(rw, req.WithContext(ctx)) // all params are OK
		})
	}
}
