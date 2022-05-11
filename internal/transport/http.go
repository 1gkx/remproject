package transport

import (
	"github.com/go-kit/kit/log"
	ht "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"
)

type HttpServer struct {
	IncreaseBalance *ht.Server
	TransferMoney   *ht.Server
}

func New(log log.Logger) *mux.Router {

	// e := endpoints.MakeEndpoints(services.New(repo, log))
	r := mux.NewRouter()

	// r.Methods("GET").Path("/api/v1/get_all").Handler(ht.NewServer(
	// 	e.GetAll,
	// 	decodeEmptyRequest,
	// 	encodeResponse,
	// ))

	return r
}
