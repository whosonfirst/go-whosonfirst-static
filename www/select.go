package www

import (
	"encoding/json"
	"github.com/aaronland/go-http-sanitize"
	"github.com/tidwall/gjson"
	"github.com/whosonfirst/go-reader"
	"net/http"
	"regexp"
)

type SelectHandlerOptions struct {
	Pattern *regexp.Regexp
}

func SelectHandler(r reader.Reader, opts *SelectHandlerOptions) (http.Handler, error) {

	fn := func(rsp http.ResponseWriter, req *http.Request) {

		query, err := sanitize.GetString(req, "select")

		if err != nil {
			http.Error(rsp, "Invalid parameters", http.StatusBadRequest)
			return
		}

		if query == "" {
			http.Error(rsp, "Missing select", http.StatusBadRequest)
			return
		}

		if !opts.Pattern.MatchString(query) {
			http.Error(rsp, "Invalid select", http.StatusBadRequest)
			return
		}

		uri, err, status := ParseURIFromRequest(req, r)

		if err != nil {
			http.Error(rsp, err.Error(), status)
			return
		}

		f := uri.Feature

		query_rsp := gjson.GetBytes(f.Bytes(), query)

		var rsp_body []byte

		if query_rsp.Exists() {

			enc, err := json.Marshal(query_rsp.Value())

			if err != nil {
				http.Error(rsp, err.Error(), http.StatusInternalServerError)
				return
			}

			rsp_body = enc
		}

		rsp.Header().Set("Content-Type", "application/json")
		rsp.Header().Set("Access-Control-Allow-Origin", "*")

		rsp.Write(rsp_body)
	}

	h := http.HandlerFunc(fn)
	return h, nil
}
