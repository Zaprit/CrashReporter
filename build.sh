#!/bin/bash
GOOS=linux GOARCH=amd64 go build ./cmd/crashreporter
rm crashreporter.zip
zip -r crashreporter.zip crashreporter static
