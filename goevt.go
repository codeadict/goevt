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
	"net/http"
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


}
