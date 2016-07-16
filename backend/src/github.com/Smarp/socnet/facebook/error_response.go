package facebook

import (
	"encoding/json"
	"io"
	"io/ioutil"
)

type FaceookErrorResponse struct {
	Message string `json:"message"`
	Type    string `json:"type"`
	Code    int    `json:"code"`
	SubCode int    `json:"error_subcode"`
}

type wrapper struct {
	Error *FaceookErrorResponse `json:"error"`
}

var ParseError = func(body io.Reader) (errorResponse *FaceookErrorResponse, stringBody string, err error) {
	var byteBody []byte
	byteBody, err = ioutil.ReadAll(body)
	if err != nil {
		return nil, "", err
	}
	stringBody = string(byteBody)
	var w wrapper
	err = json.Unmarshal(byteBody, &w)
	if err != nil {
		return nil, stringBody, err
	}
	return w.Error, stringBody, nil
}
