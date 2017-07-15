package main

import "fmt"
import "encoding/json"

// GcmURL url for GCM
const GcmURL string = "https://gcm-http.googleapis.com/gcm/send"

// FcmURL url for FCM
const FcmURL string = "https://gcm-http.googleapis.com/gcm/send"

func main() {
	jsonBytes, err := generatePayload("{\"Hello\":1, \"obj\":[{\"a\":1}]}", "")
	if err != nil {
		fmt.Println(err)
	}

	println(string(jsonBytes))
}

func generatePayload(data string, to string) ([]byte, error) {
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
