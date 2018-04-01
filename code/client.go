package code

import (
	"net/http"
	"io/ioutil"
)

func FetchData(url, id string) (string, error) {
	resp, err := http.DefaultClient.Get(url+"/data/"+id)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	return string(body), nil
}