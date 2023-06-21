// Calling GitHub API
package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"
)

// User is a github user information
type User struct {
	Login    string
	Name     string
	NumRepos int `json:"public_repos"`
}

func main() {
	// user is a pointer to a User struct
	user, err := userInfo("tebeka")
	if err != nil {
		log.Fatalf("error: %s", err)
	}

	fmt.Printf("%#v\n", user)

	// can also access the fields of User
	// GO handles dereferencing automatically when access fields through the pointer
	// so same as (*user).Login
	fmt.Println("Login:", user.Login)
	fmt.Println("Name:", (*user).Name)
	fmt.Println("NumRepos:", user.NumRepos)
}

// userInfo return information on github user
func userInfo(login string) (*User, error) { // returns a pointer to a User struct
	// HTTP call
	url := fmt.Sprintf("https://api.github.com/users/%s", url.PathEscape(login))
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf(resp.Status)
	}

	// Decode JSON
	user := User{Login: login}
	dec := json.NewDecoder(resp.Body)
	if err := dec.Decode(&user); err != nil {
		return nil, err
	}

	return &user, nil
	// &user is the address of the user struct, which is a pointer to the user struct
}
