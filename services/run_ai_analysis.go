package services

import (
	http "net/http"
)

const (
	// NLPAnalysisURL is the URL for the NLP analysis service
	CORE_NLP_ANALYSIS = "http://localhost:8100/core-analyze"
	NEXT_NLP_ANALYSIS = "http://localhost:8200/next-analyze"
)

func RunNLPAnalysis(post_contents string, tags []string, mediaLinks []string) (string, error) {
	// Fetch data from the URL
	//data, err := fetchDataViaPost(CORE_NLP_ANALYSIS, post_contents)
	//if err != nil {
	//	return "", err
	//}
	//
	//// Perform NLP analysis
	////analysis, err := (data)
	//if err != nil {
	//	return "", err
	//}
	//
	//return analysis, nil
	return "", nil
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
