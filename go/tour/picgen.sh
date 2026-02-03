#!/bin/bash
go run types.go | cut -c 7- | base64 -d > temp.png
