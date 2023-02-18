#!/bin/bash
GOOS=linux GOARCH=amd64 go build ./cmd/crashreporter
zip -r crashreporter.zip crashreporter web
