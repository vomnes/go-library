package tests

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http/httptest"
	"reflect"

	libPretty "github.com/vomnes/go-library/pretty"

	"github.com/kylelemons/godebug/pretty"
)

// ChargeResponse allows to mode http body in structure, used for tests
func ChargeResponse(w *httptest.ResponseRecorder, response interface{}) error {
	res := w.Result()
	defer res.Body.Close()
	decoder := json.NewDecoder(res.Body)
	err := decoder.Decode(response)
	return err
}

type responseBodyError struct {
	Error string
}

// ReadBodyError allows to read body in error case, used for tests
// Return the error string
func ReadBodyError(r io.Reader) string {
	body, err := ioutil.ReadAll(r)
	if err != nil {
		log.Fatal(libPretty.Error(err.Error()))
	}
	var responseBody responseBodyError
	err = json.Unmarshal(body, &responseBody)
	if err != nil {
		return ""
	}
	return responseBody.Error
}

// CompareResponseJSONCode check the http response in the tests
// Check the http code
// Check the http json
// Return an error, nil if the are no error
func CompareResponseJSONCode(w *httptest.ResponseRecorder, expectedCode int, expectedJSONResponse interface{}) []string {
	var errorArray []string
	if w.Result().StatusCode != expectedCode {
		errorArray = append(errorArray, fmt.Sprintf("Must return an error with http code \x1b[1;32m%d\033[0m not \x1b[1;31m%d\033[0m.\n", expectedCode, w.Result().StatusCode))
	}
	var response interface{}
	// Handle array and non-array response
	if reflect.TypeOf(expectedJSONResponse).Kind() == reflect.Slice {
		response = []map[string]interface{}{}
	} else {
		response = map[string]interface{}{}
	}
	if err := ChargeResponse(w, &response); err != nil {
		if err.Error() != "EOF" {
			errorArray = append(errorArray, "\x1b[1;31m"+err.Error()+"\033[0m\n")
		}
	}
	if compare := pretty.Compare(&expectedJSONResponse, response); compare != "" {
		errorArray = append(errorArray, compare)
	}
	return errorArray
}
