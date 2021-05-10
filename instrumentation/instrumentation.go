package instrumentation

import (
	"encoding/json"
	"net/http"
)

type Pong struct {
	Respo string
	Stat  string
}

func Ping() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		responseforping := Pong{"pong", "ok"}
		js, err := json.Marshal(responseforping)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.Write(js)
	}
}
