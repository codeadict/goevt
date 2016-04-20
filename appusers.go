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
    "encoding/json"
    "fmt"
)

const (
	users_url         = "/users"                              // Get users list
  create_user_url   = "/auth/evrythng/users"                // Create a new user
  activate_user_url = "/auth/evrythng/users/:id/validate"   // Activate a user
	user_url          = "/users/:id"                          // Get, Delete, Update a single user.
)


type UserCustomFields struct {
  Key           string            `json:"key,omitempty"`
  Value         string            `json:"value,omitempty"`
}
// User represent a user returned in the /users/ call.
type User struct {
	ID              string            `json:"id"`
  EMail           string            `json:"email"`
	FirstName       string            `json:"firstName,omitempty"`
  LastName        string            `json:"lastName,omitempty"`
  Password        string            `json:"password,omitempty"`
  Birthday        string            `json:"birthday,omitempty"`
  Gender          string            `json:"gender,omitempty"`
  Timezone        string            `json:"timezone,omitempty"`
  Locale          string            `json:"locale,omitempty"`
	Photo           string            `json:"photo,omitempty"` //Todo: check how to convert it to base 64
  CustomFields    *UserCustomFields `json:"customFields,omitempty"`
  Tags            []string          `json:"tags,omitempty"`
}


/*
Get a list of users.
    GET /users/
*/
func (c *Client) Users() ([]*User, error) {

	url := c.ResourceUrl(users_url)

	var users []*User

	contents, err := c.buildAndExecRequest("GET", url, nil)
	if err == nil {
		err = json.Unmarshal(contents, &users)
	}

	return users, err
}

/*
Create a new user.
    POST /auth/evrythng/users
*/
func (c *Client) CreateUser(data *User) error {
	url := c.ResourceUrl(create_user_url)
  encodedRequest, err := json.Marshal(data)
	if err != nil {
		return
	}

	created, err = c.buildAndExecRequest("POST", url, encodedRequest)
  if err != nil {
   return
  }

  user = new(User)
  err = json.Unmarshal(created, &user)
  return user, err
}

/*
Get a user details.
    GET /users/:id
Parameters:
    id The ID of a user
Usage:
	user, err := goevt.User("your_user_id")
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Printf("%+v\n", user)
*/
func (c *Client) User(id string) (*User, error) {

	url := c.ResourceUrl(user_url, map[string]string{":id": id})

	user := new(User)

	contents, err := c.buildAndExecRequest("GET", url, nil)
	if err == nil {
		err = json.Unmarshal(contents, &user)
	}

	return user, err
}

/*
Delete a user.
    DELETE /users/:id
Parameters:
    id The ID of a user
Usage:
	user, err := goevt.DeleteUser("your_user_id")
	if err != nil {
		fmt.Println(err.Error())
	}
*/
func (c *Client) DeleteUser(id string) error {
	url := c.ResourceUrl(user_url, map[string]string{":id": id})
	var err error
	_, err = c.buildAndExecRequest("DELETE", url, nil)
	return err
}
