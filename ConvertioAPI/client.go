package convertioapi

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type Api struct {
	client *http.Client
	apiKey string
}

func NewApi(apikey string) Api {
	return Api{
		client: &http.Client{},
		apiKey: apikey,
	}
}

func (a *Api) MethodPost(payload []string) (*http.Response, error) {
	//file, format, filename, "base64"
	data := make(map[string]string)
	data["apikey"] = a.apiKey
	data["file"] = payload[0]
	if len(payload) > 2 {
		data["input"] = payload[3]
		data["filename"] = payload[2]
	}
	data["outputformat"] = payload[1]

	jsonData, err := json.Marshal(data)
	if err != nil {
		log.Printf("Error marshaling data: %s", err)
		return nil, err
	}

	req, err := http.NewRequest("POST", "https://api.convertio.co/convert", bytes.NewBuffer(jsonData))
	if err != nil {
		log.Printf("Error creating request: %s", err)
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")

	//log.Println(req) //TODO delete
	resp, err := a.client.Do(req)
	if err != nil {
		log.Printf("Error sending request: %s", err)
		return nil, err
	}

	return resp, nil
}

func (a *Api) MethodGet(endpoint, id string) (*http.Response, error) {
	var link string
	switch endpoint {
	case "status":
		link = fmt.Sprintf("https://api.convertio.co/convert/%s/status", id)
	case "result":
		link = fmt.Sprintf("http://api.convertio.co/convert/%s/dl", id)
	default:
		log.Printf("Unsupported endpoint")
		return nil, nil
	}

	req, err := http.NewRequest("GET", link, nil)
	if err != nil {
		log.Printf("Error creating request: %s", err)
		return nil, err
	}

	resp, err := a.client.Do(req)
	if err != nil {
		log.Printf("Error sending request: %s", err)
		return nil, err
	}

	return resp, nil
}
