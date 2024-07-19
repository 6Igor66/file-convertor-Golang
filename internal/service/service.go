package service

import (
	"encoding/json"
	"io"
	"log"
	convertioapi "tgbot/ConvertioAPI"
	"time"
)

func TransformFile(a convertioapi.Api, payload []string) (string, error) {
	resp, err := a.MethodPost(payload)
	if err != nil {
		log.Printf("Error post req: %s", err)
		return "", err
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Printf("Error reading resp body: %s", err)
		return "", err
	}

	var postResp PostResponse
	err = json.Unmarshal(body, &postResp)
	if err != nil {
		log.Printf("Error unmarshaling JSON: %s", err)
		return "", err
	}

	log.Println(postResp)

	time.Sleep(15 * time.Second)

	resp, err = a.MethodGet("result", postResp.Data.ID)
	if err != nil {
		log.Printf("Error get res req: %s", err)
		return "", err
	}

	body, err = io.ReadAll(resp.Body)
	if err != nil {
		log.Printf("Error reading resp body: %s", err)
		return "", err
	}

	var result GetFileResponse
	err = json.Unmarshal(body, &result)
	//log.Println(result)
	if err != nil {
		log.Printf("Error unmarshaling JSON: %s", err)
		return "", err
	}

	if result.Status == "ok" {
		return result.Data.Content, nil
	}

	return "", err // TODO обработать норм ошибку
}
