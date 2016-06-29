package bart

import (
	"encoding/xml"
	"github.com/google/go-querystring/query"
	"net/http"
)

const (
	baseEndpointUrl = "http://api.bart.gov/api/"
	cmdParam        = "cmd"
	bartApiKey      = "MW9S-E7SL-26DU-VV8V"
	keyParam        = "key"
)

type BartClient struct {
	Advisories  *AdvisoryService
	StationInfo *StationInfoService
}

type BartResponseMeta struct {
	Uri     string `xml:"uri"`
	Date    string `xml:"date"`
	Time    string `xml:"time"`
	Message string `xml:"message"`
}

func New() BartClient {
	return BartClient{}
}

func makeRequest(endpoint string, command string, requestParams interface{}, response interface{}) error {
	v, _ := query.Values(requestParams)
	v.Add(cmdParam, command)
	v.Add(keyParam, bartApiKey)
	values := v.Encode()
	url := baseEndpointUrl + endpoint + "?" + values

	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	xmlDecoder := xml.NewDecoder(resp.Body)

	err = xmlDecoder.Decode(response)
	if err != nil {
		return err
	}

	return nil
}
