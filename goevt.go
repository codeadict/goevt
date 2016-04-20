// Copyright 2016 Dairon Medina <http://github.com/codeadict>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
package goevt

import (
	"bytes"
	"crypto/tls"
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

const (
  libraryVersion = "0.0.1"
	defaultBaseURL = "https://api.evrythng.com"
	userAgent      = "goevt/" + libraryVersion
	responseType   = "application/json"
)

// Client manages all the communication with Everythng API.
type Client struct {
  // The HTTP client that communicates with the API.
  client *http.Client

	BaseURL *url.URL

	// User agent for client
	UserAgent string

}

var (
	skipCertVerify = flag.Bool("goevt.skip-cert-check", false,
		`If set to true, Everythng client will skip certificate checking for https, possibly exposing your system to MITM attack.`)
)

func NewClient(baseUrl, token string) *Client {
	config := &tls.Config{InsecureSkipVerify: *skipCertVerify}
	transport := &http.Transport{
		Proxy:           http.ProxyFromEnvironment,
		TLSClientConfig: config,
	}

	client := &http.Client{Transport: transport}

	if baseUrl == nil {
		baseUrl = url.Parse(defaultBaseURL)
	}

	return &Client{
		BaseUrl: baseUrl,
		Token:   token,
		Client:  client,
	}
}

func (c *Client) ResourceUrl(url string, params map[string]string) string {

	if params != nil {
		for key, val := range params {
			url = strings.Replace(url, key, val, -1)
		}
	}

	url = c.BaseUrl + url

	return url
}

func (c *Client) execRequest(method, url string, body []byte) (*http.Response, error) {
	var req *http.Request
	var err error

	if body != nil {
		reader := bytes.NewReader(body)
		req, err = http.NewRequest(method, url, reader)
	} else {
		req, err = http.NewRequest(method, url, nil)
	}

	if err != nil {
		panic("Error while building request for EVT.")
	}

	resp, err := c.Client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("Client.Do error: %q", err)
	}

	if resp.StatusCode >= http.StatusBadRequest {
		err = fmt.Errorf("*EVT.buildAndExecRequest failed: <%d> %s", resp.StatusCode, req.URL)
	}

	return resp, err
}

func (c *Client) buildAndExecRequest(method, url string, body []byte) ([]byte, error) {
	resp, err := c.execRequest(method, url, body)

	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	contents, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("%s", err)
	}

	return contents, err
}
