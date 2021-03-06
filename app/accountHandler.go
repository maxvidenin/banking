package app

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/maxvidenin/banking/dto"
	"github.com/maxvidenin/banking/service"
)

type AccountHandler struct {
	service service.AccountService
}

func (ah AccountHandler) NewAccount(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	customerId := vars["customer_id"]
	var req dto.NewAccountRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		writeResponse(w, http.StatusBadRequest, err.Error())
	} else {
		req.CustomerId = customerId
		accRes, appErr := ah.service.NewAccount(req)
		if appErr != nil {
			writeResponse(w, http.StatusBadRequest, appErr.Message)
		} else {
			writeResponse(w, http.StatusCreated, accRes)
		}

	}
}

func (ah AccountHandler) MakeTransaction(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	accountId := vars["account_id"]
	customerId := vars["customer_id"]
	var req dto.TransactionRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		writeResponse(w, http.StatusBadRequest, err.Error())
	} else {
		req.AccountId = accountId
		req.CustomerId = customerId
		transRes, appErr := ah.service.MakeTransaction(req)
		if appErr != nil {
			writeResponse(w, http.StatusBadRequest, appErr.AsMessage())
		} else {
			writeResponse(w, http.StatusCreated, transRes)
		}

	}
}
