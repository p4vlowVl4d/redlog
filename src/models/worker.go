package models

import "net/http"

type Worker struct {
	Server http.Server
}