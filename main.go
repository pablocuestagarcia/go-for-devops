package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

type User struct {
	Id       int    `json:"id"`
	Name     string `json:"name"`
	Username string `json:"username"`
	Email    string `json:"email"`
}

func fetchUsers(url string) ([]User, error) {

	resp, err := http.Get(url)

	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("Error: received status code %d", resp.StatusCode)
	}

	body, err := io.ReadAll(resp.Body)

	if err != nil {
		return nil, err
	}

	var users []User

	if err := json.Unmarshal(body, &users); err != nil {
		return nil, err
	}

	return users, nil
}

func printUsers(users []User) {
	for i, user := range users {
		fmt.Printf("%d- User %d\n", i, user.Name)
	}
}

func main() {
	users, err := fetchUsers("https://jsonplaceholder.typicode.com/users")

	if err != nil {
		os.Exit(1)
	}

	printUsers(users)

}
