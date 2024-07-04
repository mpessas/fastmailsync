package main

import (
	"bytes"
	"log"
	"os"
	"path"
	"text/template"
)

type MessageInfo struct {
	AccountId string
	Id        string
	BlobId    string
}

const DOWNLOAD_TEMPLATE = "https://www.fastmailusercontent.com/publicjmap/download/{{.AccountId}}/{{.BlobId}}/email?type=application/octet-stream"

func download(maildir string, info MessageInfo, auth BasicAuthInfo) {
	tmpl, err := template.New("download_template").Parse(DOWNLOAD_TEMPLATE)
	if err != nil {
		log.Fatalln(err)
	}
	var buf bytes.Buffer
	tmpl.Execute(&buf, info)

	content := get(buf.String(), auth)

	f, err := os.OpenFile(path.Join(maildir, "inbox/tmp", info.Id), os.O_RDWR|os.O_CREATE, 0600)
	if err != nil {
		log.Println("Error opening local file:", err)
	}
	_, err = f.Write(content)
	if err != nil {
		log.Fatalln("Error storing email locally:", err)
	}
}
