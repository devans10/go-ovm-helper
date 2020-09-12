package ovmhelper

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
)

// Client - Interface to connect with OVM Manager
type Client struct {
	User       string
	Password   string
	BaseURL    *url.URL
	client     *http.Client
	Vms        *VMService
	Vmcds      *VmcdService
	Vmcsms     *VmcsmService
	Vmcnms     *VmcnmService
	Vns        *VnService
	Vds        *VdService
	Vdms       *VdmService
	Jobs       *JobService
	Repos      *RepoService
	ServerPool *ServerPoolService
}

// NewClient - instanciate a new client object
func NewClient(user string, password string, endpoint string) *Client {
	baseURL, _ := url.Parse(endpoint)
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	c := &Client{User: user, Password: password, BaseURL: baseURL}
	c.client = &http.Client{Transport: tr}
	c.Vms = &VMService{client: c}
	c.Vmcds = &VmcdService{client: c}
	c.Vmcsms = &VmcsmService{client: c}
	c.Vmcnms = &VmcnmService{client: c}
	c.Vds = &VdService{client: c}
	c.Vdms = &VdmService{client: c}
	c.Vns = &VnService{client: c}
	c.Jobs = &JobService{client: c}
	c.Repos = &RepoService{client: c}
	c.ServerPool = &ServerPoolService{client: c}
	return c
}

// NewRequest - Create a New API request to send to OVM Manager
func (pc *Client) NewRequest(method string, rsc string, params map[string]string, data interface{}) (*http.Request, error) {

	baseURL, err := url.Parse(pc.BaseURL.String() + rsc)
	if err != nil {
		return nil, err
	}

	if params != nil {
		ps := url.Values{}
		for k, v := range params {
			log.Printf("[DEBUG] key: %s, value: %s \n", v, k)
			ps.Set(k, v)
		}
		baseURL.RawQuery = ps.Encode()
	}
	req, err := http.NewRequest(method, baseURL.String(), nil)
	if data != nil {
		jsonString, _ := json.Marshal(data)
		req, err = http.NewRequest(method, baseURL.String(), bytes.NewBuffer(jsonString))
	}
	req.SetBasicAuth(pc.User, pc.Password)
	req.Header.Add("content-type", "application/json; charset=utf-8")
	req.Header.Add("Accept", "application/json")
	req.Header.Set("Content-Type", "application/json")
	return req, err
}

// Do - Execute the request to OVM Manager
func (pc *Client) Do(req *http.Request, v interface{}) (*http.Response, error) {
	resp, err := pc.client.Do(req)
	if err != nil {
		fmt.Println("Do request failed")
		return nil, err
	}
	defer resp.Body.Close()

	//body, err := ioutil.ReadAll(resp.Body)
	//log.Printf("[ERROR] Response code: %v", resp.Status)
	//fmt.Println(string(body))

	log.Printf("[INFO] Response code: %v", resp.Status)

	if err := validateResponse(resp); err != nil {
		return resp, err
	}

	err = decodeResponse(resp, v)
	return resp, err

}

func decodeResponse(r *http.Response, v interface{}) error {
	if v == nil {
		return fmt.Errorf("nil interface provided to decodeResponse")
	}

	bodyBytes, _ := ioutil.ReadAll(r.Body)
	bodyString := string(bodyBytes)
	err := json.Unmarshal([]byte(bodyString), &v)
	//err := json.Unmarshal([]byte(bodyBytes), &v)
	return err
}

// Takes an HTTP response and determines whether it was successful.
// Returns nil if the HTTP status code is within the 2xx range.  Returns
// an error otherwise.
func validateResponse(r *http.Response) error {
	if c := r.StatusCode; 200 <= c && c <= 299 {
		return nil
	}

	bodyBytes, _ := ioutil.ReadAll(r.Body)
	bodyString := string(bodyBytes)
	return fmt.Errorf("Response code: %d, ResponeBody: %s", r.StatusCode, bodyString)
}
