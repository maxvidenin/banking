package app

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/gorilla/mux"
	"github.com/maxvidenin/banking/dto"
	"github.com/maxvidenin/banking/errs"
	"github.com/maxvidenin/banking/mocks/service"
)

var mockService *service.MockCustomerService
var ch CustomerHandlers
var router *mux.Router

func setup(t *testing.T) func() {
	ctrl := gomock.NewController(t)
	mockService = service.NewMockCustomerService(ctrl)
	ch = CustomerHandlers{mockService}
	router = mux.NewRouter()
	return func() {
		router = nil
		defer ctrl.Finish()
	}
}

func TestShouldReturnCustomersWithSuccess(t *testing.T) {
	// Arrange
	teardown := setup(t)
	defer teardown()
	dummyCustomers := []dto.CustomerResponse{
		{1001, "Max", "Amsterdam", "1122XZ", "1980-11-21", "1"},
		{1002, "John", "London", "1122XZ", "1990-05-12", "1"},
	}
	mockService.EXPECT().GetAllCustomers("1").Return(dummyCustomers, nil)

	router.HandleFunc("/customers", ch.getAllCustomers).Methods(http.MethodGet)

	request, _ := http.NewRequest(http.MethodGet, "/customers?status=1", nil)
	// Act
	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, request)

	// Assert
	if recorder.Code != http.StatusOK {
		t.Errorf("Expected status code %d but got %d", http.StatusOK, recorder.Code)
	}
}

func TestShouldReturnErrorMessage(t *testing.T) {
	// Arrange
	teardown := setup(t)
	defer teardown()

	mockService.EXPECT().GetAllCustomers("1").Return(nil, errs.NewUnexpectedError("Error"))

	router.HandleFunc("/customers", ch.getAllCustomers).Methods(http.MethodGet)

	request, _ := http.NewRequest(http.MethodGet, "/customers?status=1", nil)
	// Act
	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, request)

	// Assert
	if recorder.Code != http.StatusInternalServerError {
		t.Errorf("Expected status code %d but got %d", http.StatusOK, recorder.Code)
	}
}
