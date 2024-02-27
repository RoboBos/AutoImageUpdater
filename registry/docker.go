package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
	// log "github.com/sirupsen/logrus"
)

type Layer struct {
	Digest      string `json:"digest"`
	Size        int    `json:"size"`
	Instruction string `json:"instruction"`
}

type Images struct {
	Architecture string    `json:"architecture"`
	Features     string    `json:"features"`
	Variant      string    `json:"variant"`
	Digest       string    `json:"digest"`
	Layers       []Layer   `json:"layers"`
	OS           string    `json:"os"`
	OSFeatures   string    `json:"os_features"`
	OSVersion    string    `json:"os_version"`
	Size         int       `json:"size"`
	Status       string    `json:"status"`
	LastPulled   time.Time `json:"last_pulled"`
	LastPushed   time.Time `json:"last_pushed"`
}

type Result struct {
	ID                  int       `json:"id"`
	Images              []Images  `json:"images"`
	Creator             int       `json:"creator"`
	LastUpdated         time.Time `json:"last_updated"`
	LastUpdater         int       `json:"last_updater"`
	LastUpdaterUsername string    `json:"last_updater_username"`
	Name                string    `json:"name"`
	Repository          int       `json:"repository"`
	FullSize            int       `json:"full_size"`
	V2                  bool      `json:"v2"`
	Status              string    `json:"status"`
	TagLastPulled       time.Time `json:"tag_last_pulled"`
	TagLastPushed       time.Time `json:"tag_last_pushed"`
}

type Response struct {
	Count    int      `json:"count"`
	Next     string   `json:"next"`
	Previous string   `json:"previous"`
	Results  []Result `json:"results"`
}

var BASE_URL = "https://hub.docker.com"

func GetImageVersions(namespace string, repository string, page_size string, versionValidtor VersioningSchemesValidtor) ([]string, error) {
	var versions []string

	tagsApi := fmt.Sprintf("%s/v2/namespaces/%s/repositories/%s/tags?page_size=%s", BASE_URL, namespace, repository, page_size)
	reponse, err := http.Get(tagsApi)
	if err != nil {
		return nil, err
	}
	defer reponse.Body.Close()

	var tags Response
	err = json.NewDecoder(reponse.Body).Decode(&tags)
	if err != nil {
		return nil, err
	}


	for _, results := range tags.Results {
		if versionValidtor(results.Name){
			versions = append(versions, results.Name)
		}
	}
	return versions, nil
}
