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

// GetAccessToken returns struct of TokenResponse
// No need to call SetAccessToken to apply new access token for current Client
// Endpoint: POST /token/
// func (c *Client) GetAccessToken() (*TokenResponse, error) {
// 	buf := bytes.NewBuffer([]byte("grant_type=client_credentials"))
// 	req, err := http.NewRequest("POST", fmt.Sprintf("%s%s", c.APIBase, "/token/"), buf)
// 	if err != nil {
// 		return &TokenResponse{}, err
// 	}
//
// 	req.SetBasicAuth(c.ClientID, c.Secret)
// 	req.Header.Set("Content-type", "application/x-www-form-urlencoded")
//
// 	t := TokenResponse{}
// 	err = c.Send(req, &t)
//
// 	// Set Token fur current Client
// 	if t.Token != "" {
// 		c.Token = &t
// 	}
//
// 	return &t, err
// }
