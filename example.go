package main

import (
	"Pawelek242/httpclient-go/gohttp"
	"fmt"
	"io/ioutil"
	"net/http"
)

var githubHttpClient = getGithubClient()

func getGithubClient() gohttp.HttpClient {
	client := gohttp.New()

	commonHeaders := make(http.Header)
	commonHeaders.Set("Authorization", "Bearer ABC-123")
	client.SetHeaders(commonHeaders)

	return client
}

func main() {
	getUrls()
	getUrls()
	getUrls()


}

type User struct {
	FirstName string `json:"first_name""`
	LastName string `json:"last_name"`
}

func user(user User) {
	response, err := githubHttpClient.Post("https//api.github.com", nil, user)
	if err != nil {
		panic(err)

	}

	fmt.Println(response.StatusCode)

	bytes, _ := ioutil.ReadAll(response.Body)
	fmt.Println(string(bytes))

}

func getUrls() {

	response, err := githubHttpClient.Post("https://api.github.com", nil, nil )
	if err != nil {
		panic(err)
	}
	fmt.Println(response.StatusCode)

	bytes, _ := ioutil.ReadAll(response.Body)
	fmt.Println(string(bytes))


}