package resolver

import "net/http"

func (res *Resolver) LogHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(res.ShowStat()))
}
