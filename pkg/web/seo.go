package web

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

// Seo contain SEO data correspond to route
type Seo map[string]map[string]string

// initSEO to initialize SEO Data in memory
func initSEO(path string) (Seo, error) {
	var (
		seoData Seo
	)

	seoFileData, err := os.Open(path)
	if err != nil {
		return nil, err
	}

	defer seoFileData.Close()
	// read our opened jsonFile as a byte array.
	byteValue, _ := ioutil.ReadAll(seoFileData)

	jsonErr := json.Unmarshal(byteValue, &seoData)
	if jsonErr != nil {
		return nil, jsonErr
	}

	return seoData, nil

}
