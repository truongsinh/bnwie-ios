package facebook

import (
	"encoding/json"
	"fmt"
	// "io/ioutil"
	"net/http"
	"net/url"
)

type batchRequestStruct struct {
	Method      string `json:"method"`
	RelativeUrl string `json:"relative_url"`
}

func newBatchRequest(method string, edge string, params map[string]string) *batchRequestStruct {
	b := batchRequestStruct{method, ""}
	uParams := url.Values{}
	for k, v := range params {
		uParams.Add(k, v)
	}
	b.RelativeUrl = edge + "?" + uParams.Encode()
	return &b
}

func BatchRequest(c *http.Client, accessToken string, input []*batchRequestStruct) (res *http.Response, err error) {

	fbParam := url.Values{}
	fbParam.Add("access_token", accessToken)
	fbParam.Add("format", "json")
	fbParam.Add("include_headers", "false")
	b, err := json.Marshal(input)
	if err != nil {
		return nil, err
	}
	fbParam.Add("batch", string(b))

	u := url.URL{
		Scheme: "https",
		Host:   endpoint,
		Path:   graphApiVersion,
	}

	res, err = c.PostForm(u.String(), fbParam)
	return res, err
}

type batchResponseStruct struct {
	StatusCode int    `json:"code"`
	Body       string `json:"body"`
}

func BatchParse(resp *http.Response, ierr error) ([]batchResponseStruct, error) {
	if ierr != nil {
		return nil, ierr
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		err := fmt.Errorf("Err Facebook Graph Batch statusCode %d", resp.StatusCode)
		return nil, err
	}
	batchr := []batchResponseStruct{}
	decoder := json.NewDecoder(resp.Body)
	err := decoder.Decode(&batchr)
	return batchr, err
}
