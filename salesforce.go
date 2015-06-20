// Package salesforce is based on the Python "Simple-Salesforce" library. It provides the
// basic CRUD operations for Salesforce. Currently only OAuth2 authentication is supported
package salesforce

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/cookiejar"
	"net/url"
)

const Version = 31

// SalesforceClient struct contains information about a salesforce
// OAuth 2 session
type SalesforceClient struct {
	InstanceURL string
	SessionID   string
	BaseURL     string
	ObjectName  string
	Header      map[string]string
	HttpClient  *http.Client
}

func (s *SalesforceClient) getJsonDataFromURL(requestURL string) (interface{}, error) {

	var data interface{}
	resp, err := s.HttpClient.Get(requestURL)
	if err != nil {
		return data, err
	}

	data, err = httpResponseToJson(resp)
	return data, err
}

func (s *SalesforceClient) setJsonDataForURL(httpMethod string, requestURL string,
	data interface{}) (interface{}, error) {

	retData := make(map[string]interface{})
	jsonData, err := json.Marshal(data)
	if err != nil {
		return retData, err
	}

	req, err := http.NewRequest(httpMethod, requestURL, bytes.NewReader(jsonData))
	if err != nil {
		return retData, err
	}
	// set the HTTP headers
	for k, v := range s.Header {
		req.Header.Set(k, v)
	}

	resp, err := s.HttpClient.Do(req)
	if err != nil {
		return retData, err
	}

	data, err = httpResponseToJson(resp)
	return data, err
}

func (s *SalesforceClient) deleteDataForURL(requestURL string) error {

	req, err := http.NewRequest("DELETE", requestURL, nil)
	if err != nil {
		return err
	}
	// set the HTTP headers
	for k, v := range s.Header {
		req.Header.Set(k, v)
	}

	_, err = s.HttpClient.Do(req)
	if err != nil {
		return err
	}

	return nil
}

// NewSalesforceClient returns an instance of SalesforceClient
// initialied with the required fields
func NewSalesforceClient(instanceURL string, sessionID string) *SalesforceClient {

	baseURL := fmt.Sprintf("https://%s/services/data/v%d/", instanceURL, Version)
	header := map[string]string{
		"Content-Type":  "application/json",
		"Authorization": "Bearer " + sessionID,
	}

	cookieJar, _ := cookiejar.New(nil)
	httpClient := &http.Client{
		Jar: cookieJar,
	}

	return &SalesforceClient{
		InstanceURL: instanceURL,
		SessionID:   sessionID,
		BaseURL:     baseURL,
		Header:      header,
		HttpClient:  httpClient,
	}
}

// Describe returns all the objects present for the Salesforce instance
func (s *SalesforceClient) Describe() (interface{}, error) {

	requestURL := s.BaseURL + "sobjects"
	data, err := s.getJsonDataFromURL(requestURL)
	if err != nil {
		return data, err
	}

	return data, nil
}

// Search returns the result from a raw SOQL query string
func (s *SalesforceClient) Search(query string) (interface{}, error) {

	v := url.Values{}
	v.Set("q", query)

	requestURL := s.BaseURL + "search/" + v.Encode()
	data, err := s.getJsonDataFromURL(requestURL)
	if err != nil {
		return data, err
	}

	return data, nil
}

// Create will create an object and it's corresponding data with the POST method
func (s *SalesforceClient) Create(objectName string,
	data interface{}) (interface{}, error) {

	requestURL := s.BaseURL + objectName
	data, err := s.setJsonDataForURL("POST", requestURL, data)
	if err != nil {
		return data, err
	}

	return data, nil
}

// Get will retrieve the data for an object using the GET method
func (s *SalesforceClient) Get(objectName string,
	recordID string) (interface{}, error) {

	requestURL := s.BaseURL + objectName + "/" + recordID
	data, err := s.getJsonDataFromURL(requestURL)
	if err != nil {
		return data, err
	}

	return data, nil
}

// Upsert will update an existing object's data with the PATCH method
func (s *SalesforceClient) Upsert(objectName string, recordID string,
	data interface{}) (interface{}, error) {

	requestURL := s.BaseURL + objectName + "/" + recordID
	data, err := s.setJsonDataForURL("PATCH", requestURL, data)
	if err != nil {
		return data, err
	}

	return data, nil
}

// Update will update an object's metadata with the PATCH method
func (s *SalesforceClient) Update(objectName string, recordID string,
	data interface{}) (interface{}, error) {

	requestURL := s.BaseURL + objectName + "/" + recordID
	data, err := s.setJsonDataForURL("PATCH", requestURL, data)
	if err != nil {
		return data, err
	}

	return data, nil
}

// Delete will delete an object's metadata with the DELETE method
func (s *SalesforceClient) Delete(objectName string,
	recordID string) error {

	requestURL := s.BaseURL + objectName + "/" + recordID
	err := s.deleteDataForURL(requestURL)
	if err != nil {
		return err
	}

	return nil
}
