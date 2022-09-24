package routes

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
	"github.com/stretchr/testify/require"
)

func TestRegisterBookStoreRoutes(t *testing.T){

	m:= mux.NewRouter()
	RegisterBookstoreRoutes(m)
	
	respRec := httptest.NewRecorder()

	req,err := http.NewRequest("GET","/book/",nil)
	require.NoError(t, err,"error creating a new request")

	m.ServeHTTP(respRec,req)

	if respRec.Code != http.StatusBadRequest{
		t.Fatal("Server error",respRec.Code,"instead of ", http.StatusBadRequest)
	}



}