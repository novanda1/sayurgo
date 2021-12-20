package utils

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"os"
	"time"
)

type WhatsappMessage struct {
	To         string `json:"to" bson:"to"`
	Content    string `json:"content" bson:"content"`
	InstanceID string `json:"instances_id" bson:"instances_id"`
}

func httpClient() *http.Client {
	client := &http.Client{Timeout: 10 * time.Second}
	return client
}

func SendOtpCodeToWhatsapp(to string, message string) {
	main_url := "https://api.zuwinda.com"
	send_text_url := "/v1.2/message/whatsapp/send-text"
	full_url := main_url + send_text_url

	client := httpClient()
	payload := WhatsappMessage{
		To:         to,
		InstanceID: os.Getenv("ZUWINDA_INSTANCE"),
		Content:    "Kode OTP SayurGO anda adalah: " + message,
	}
	body, err := json.Marshal(payload)

	request, _ := http.NewRequest(http.MethodPost, full_url, bytes.NewBuffer(body))
	request.Header.Add("x-access-key", os.Getenv("ZUWINDA_TOKEN"))
	request.Header.Set("Content-Type", "application/json")

	response, err := client.Do(request)
	if err != nil {
		log.Fatal(err)
	}

	defer response.Body.Close()
}
