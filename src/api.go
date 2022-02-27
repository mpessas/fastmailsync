package main

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/mpessas/fastmailsync/queries"
)

const Url = "https://jmap.fastmail.com/api/"

type BasicAuthInfo struct {
	username string
	password string
}

func post(auth BasicAuthInfo, payload *queries.JmapPayload) {
	json_payload, err := payload.ToJson()
	if err != nil {
		log.Fatalln(err)
	}
	os.Stdout.Write(json_payload)
	buf := bytes.NewBuffer(json_payload)
	req, err := http.NewRequest(http.MethodPost, Url, buf)
	if err != nil {
		log.Fatalln(err)
	}
	req.Header.Add("Content-Type", "application/json")
	req.SetBasicAuth(auth.username, auth.password)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(body))
}
