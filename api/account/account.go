package account

import (
	"net/http"

	"github.com/asiainfoLDP/datafoundry_payment/api"
	"github.com/asiainfoLDP/datafoundry_payment/fake"
	"github.com/julienschmidt/httprouter"
	"github.com/zonesan/clog"
)

func Account(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

	clog.Info("from", r.RemoteAddr, r.Method, r.URL.RequestURI(), r.Proto)

	account := fake.Account(r)

	api.RespOK(w, account)

	return
}