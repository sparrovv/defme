package wordnik

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func EncodeSpace(str string) string {
	return strings.Replace(str, " ", "%20", -1)
}

func makeRequest(url string) (body []byte, err error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		err = fmt.Errorf("Error fetching assignments: [%v]", err)
		return
	}

	body, err = ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()

	return
}
