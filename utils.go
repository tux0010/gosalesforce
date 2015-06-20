package gosalesforce

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

func httpResponseToJson(resp *http.Response) (interface{}, error) {
	var data interface{}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return data, err
	}

	err = json.Unmarshal(body, &data)
	if err != nil {
		return data, err
	}

	return data, nil
}
