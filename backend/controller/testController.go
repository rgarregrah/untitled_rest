package controller

import (
	"encoding/json"
	"net/http"
)

type Test struct {
	Msg string `json:"msg"`
}

func TestHandler(w http.ResponseWriter, r *http.Request) {
	test1 := Test{
		Msg: "Hello, Hello, Hello",
	}
	res := createJsonResponse(&test1)
	w.Header().Set("Content-Type", "application/json")
	w.Write(res)
}

func createJsonResponse(r *Test) []byte {
	json, err := json.Marshal(*r)
	if err != nil {
		panic("json cannot marshal")
	}
	return json
}
