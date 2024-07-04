package main

import (
	"bytes"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/mpessas/fastmailsync/jmap"
)

type BasicAuthInfo struct {
	username string
	password string
}

func get(url string, auth BasicAuthInfo) []byte {
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		log.Fatalln(err)
	}
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
	return body
}

func post(url string, auth BasicAuthInfo, payload *jmap.JmapPayload, debug bool) []byte {
	json_payload, err := payload.ToJson()
	if err != nil {
		log.Fatalln(err)
	}
	if debug {
		os.Stdout.Write(json_payload)
	}
	buf := bytes.NewBuffer(json_payload)
	req, err := http.NewRequest(http.MethodPost, url, buf)
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

	if debug {
		os.Stdout.Write(body)
	}

	return body
}
