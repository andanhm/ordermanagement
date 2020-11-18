package api

import (
	"bytes"
	"context"
	"encoding/csv"
	"io/ioutil"
	"os"
	"strings"
	"time"
)

var (
	// gitVersionFilePath determines git generated version path
	gitVersionFilePath = "./version"
)

// HealthStatus represent health status of the application
type HealthStatus struct {
	Host    string `json:"host"`
	Version string `json:"version"`
	Commit  string `json:"commit"`
	Started string `json:"startedAt"`
}

// git method which allows fetch the git HEAD tag version and commit number
func git() (version, commitHead string) {

	b, _ := ioutil.ReadFile(gitVersionFilePath)

	if b != nil {
		reader := bytes.NewReader(b)
		r := csv.NewReader(reader)
		rows, _ := r.ReadAll()
		for idx, row := range rows {
			if len(row) < 1 {
				continue
			}
			switch idx {
			case 0:
				{
					commitHead = strings.TrimSpace(row[0])
				}
			case 1:
				{
					version = strings.TrimSpace(row[0])
					// ignore following rows
					break
				}
			}
		}
	}
	return
}

// Health return the status
func (api *API) Health(context context.Context) HealthStatus {
	version, commitHead := git()
	host, _ := os.Hostname()
	return HealthStatus{
		Host:    host,
		Version: version,
		Commit:  commitHead,
		Started: time.Now().UTC().String(),
	}
}
