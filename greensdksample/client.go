package greensdksample

import (
	"encoding/json"
	//"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

type DefaultClient struct {
	Profile Profile
}

func (defaultClient DefaultClient) GetResponse(path string, clinetInfo ClinetInfo, bizData BizData) string {
	clientInfoJson, _ := json.Marshal(clinetInfo)
	bizDataJson, _ := json.Marshal(bizData)

	client := &http.Client{}
	req, err := http.NewRequest(method, host+path+"?clientInfo="+url.QueryEscape(string(clientInfoJson)), strings.NewReader(string(bizDataJson)))

	if err != nil {
		// handle error
		return ErrorResult(err)
	} else {
		addRequestHeader(string(bizDataJson), req, string(clientInfoJson), path, defaultClient.Profile.AccessKeyId, defaultClient.Profile.AccessKeySecret)
		response, err := client.Do(req);
		if  err != nil {
			return ErrorResultWithCode(err, 500)
		}

		defer response.Body.Close()

		body, err := ioutil.ReadAll(response.Body)
		if err != nil {
			// handle error
			return ErrorResult(err)
		} else {
			// fmt.Println(string(body))
			m := map[string]string{}
			if err := json.Unmarshal(body, &m); err == nil {
				if _, ok := m["code"]; !ok {
					m["code"] = "500"
					m["msg"] = "third-party service request failed"
					if body, err = json.Marshal(m); err == nil {
						return string(body)
					}
				}
			}
			return string(body)
		}
	}
}

type IAliYunClient interface {
	GetResponse(path string, clinetInfo ClinetInfo, bizData BizData) string
}
