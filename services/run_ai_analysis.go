package services

import (
	models "aurora-borealis/models"
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	http "net/http"
)

const (
	// NLPAnalysisURL is the URL for the NLP analysis service
	CORE_NLP_ANALYSIS = "http://localhost:8888/bots/journal/post/analyze"
	NEXT_NLP_ANALYSIS = "http://localhost:8200/next-analyze"
)

func RunNLPAnalysis(post models.PostRequestForAnalysis) (*models.PostAnalysisResponse, error) {
	jsonData, err := json.Marshal(post)
	if err != nil {
		fmt.Println("Error marshalling post data", err)
		return nil, err
	}

	req, err := http.NewRequest("POST", CORE_NLP_ANALYSIS, bytes.NewBuffer(jsonData))
	if err != nil {
		fmt.Println("Error creating request", err)
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error making request", err)
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response", err)
		return nil, err
	}

	var analysisMetaActual models.PostAnalysisResponseActual
	var analysisMeta models.PostAnalysisResponse
	err = json.Unmarshal(body, &analysisMetaActual)
	analysisMeta = analysisMetaActual.Response
	if err != nil {
		fmt.Println("Error unmarshalling response", err)
		return nil, err
	}

	return &analysisMeta, nil
}

func fetchDataGet(url string) (string, error) {
	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	return "", nil
}

func fetchDataViaPost(url string) (string, error) {
	//resp, err := http.Post(url)
	//if err != nil {
	//	return "", err
	//}
	//defer resp.Body.Close()
	return "", nil
}
