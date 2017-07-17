package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"time"
)

// GcmURL url for GCM
const GcmURL string = "https://gcm-http.googleapis.com/gcm/send"

// FcmURL url for FCM
const FcmURL string = "https://gcm-http.googleapis.com/gcm/send"

func sendData(serverKey string, data string, to string) error {
	jsonBytes, err := createPayload(data, to)
	if err != nil {
		return errors.New("invalid data json")
	}

	header := map[string]string{
		"Content-Type":  "application/json",
		"Authorization": fmt.Sprintf("key=%v", serverKey),
	}

	_, err = sendMessage(FcmURL, header, jsonBytes)
	if err != nil {
		return err
	}

	return nil
}

func sendMessage(url string, header map[string]string, json []byte) (string, error) {
	req, err := http.NewRequest("POST", url, strings.NewReader(string(json)))
	if err != nil {
		return "", err
	}

	for k, v := range header {
		req.Header.Add(k, v)
	}

	client := &http.Client{Timeout: time.Duration(10) * time.Second}
	res, err := client.Do(req)
	if err != nil {
		return "", err
	}

	defer res.Body.Close()

	b, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return "", err
	}

	fmt.Println(string(b))

	return string(b), nil
}

func createPayload(data string, to string) ([]byte, error) {
	dataJSON := json.RawMessage(data)

	m := map[string]interface{}{
		"to":   to,
		"data": &dataJSON,
	}

	for k, v := range m {
		if v == "" {
			delete(m, k)
		}
	}

	bytes, err := json.MarshalIndent(m, "", "\t")
	return bytes, err
}
