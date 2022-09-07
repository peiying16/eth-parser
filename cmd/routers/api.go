package routers

import (
	"encoding/json"
	"eth/internal/services"
	"eth/internal/structs"
	"io/ioutil"
	"net/http"
)

type Server struct {
	s *services.TransationSvc
}

func NewServer(port string, s *services.TransationSvc) (*Server, error) {

	server := Server{
		s: s,
	}
	http.HandleFunc("/block", server.GetCurrentBlock)
	http.HandleFunc("/subscribe", server.Subscribe)
	http.HandleFunc("/transactions", server.GetTransactions)

	err := http.ListenAndServe(port, nil)
	return &server, err
}

// last parsed block
// GetCurrentBlock() int
func (s *Server) GetCurrentBlock(w http.ResponseWriter, r *http.Request) {
	num, err := s.s.GetCurrentBlock()

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Add("Content-Type", "application/json")

	resp := structs.GetBlockResp{Number: num}

	respByte, err := json.Marshal(resp)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Write(respByte)
}

// add address to observer
// Subscribe(address string) bool
func (s *Server) Subscribe(w http.ResponseWriter, r *http.Request) {

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	req := structs.SubscribeReq{}
	err = json.Unmarshal(body, &req)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	success, err := s.s.Subscribe(req.Address)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	resp := structs.SubscribeResp{
		Success: success,
	}

	respByte, err := json.Marshal(resp)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Add("Content-Type", "application/json")
	w.Write(respByte)
}

// list of inbound or outbound transactions for an address
// GetTransactions(address string) []Transaction
func (s *Server) GetTransactions(w http.ResponseWriter, r *http.Request) {
	req := structs.GetTransactionsReq{
		Address: r.URL.Query().Get("address"),
	}

	txs, err := s.s.GetTransactions(req.Address)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	resp := structs.GetTransactionsResp{
		Txs: txs,
	}

	respByte, err := json.Marshal(resp)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Add("Content-Type", "application/json")
	w.Write(respByte)

}
