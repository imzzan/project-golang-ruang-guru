package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type Animechan struct {
	Anime     string `json:"anime"`
	Character string `json:"character"`
	Quote     string `json:"quote"`
}

func ClientGet() ([]Animechan, error) {
	res, err := http.Get("https://animechan.xyz/api/quotes/anime?title=naruto")
	if err != nil {
		return nil, err
	}

	body, err := io.ReadAll(res.Body)

	if err != nil {
		return nil, err
	}
	datas := []Animechan{}
	err = json.Unmarshal(body, &datas)
	if err != nil {
		return nil, err
	}
	// Hit API https://animechan.vercel.app/api/quotes/anime?title=naruto with method GET:
	return datas, nil // TODO: replace this
}

type data struct {
	Email string `json:"email"`
	Name  string `json:"name"`
}

type Postman struct {
	Data data
	Url  string `json:"url"`
}

func ClientPost() (Postman, error) {
	postBody, _ := json.Marshal(map[string]string{
		"name":  "Dion",
		"email": "dionbe2022@gmail.com",
	})
	requestBody := bytes.NewBuffer(postBody)

	res, err := http.Post("https://postman-echo.com/post", "application/json", requestBody)

	if err != nil {
		return Postman{}, err
	}

	body, err := io.ReadAll(res.Body)

	if err != nil {
		return Postman{}, err
	}
	datas := Postman{}
	err = json.Unmarshal(body, &datas)
	if err != nil {
		return Postman{}, err
	}

	// Hit API https://postman-echo.com/post with method POST:
	return datas, nil // TODO: replace this
}

func main() {
	get, _ := ClientGet()
	fmt.Println(get)

	// post, _ := ClientPost()
	// fmt.Println(post)
}
