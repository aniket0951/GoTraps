package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

// func TestMain(t *testing.T) {
// 	type Args struct {
// 		Pannumber string
// 	}

// 	tests := []struct {
// 		Name string
// 		Args Args
// 		Want bool
// 	}{
// 		{
// 			Name: "Aniket PanCard",
// 			Args: Args{Pannumber: "1234567890"},
// 			Want: false,
// 		},
// 		{
// 			Name: "Mayur PanCard",
// 			Args: Args{Pannumber: "AWSER12AAA"},
// 			Want: false,
// 		},
// 		{
// 			Name: "Mr.Programmer",
// 			Args: Args{Pannumber: "EAIPS9832A"},
// 			Want: false,
// 		},
// 	}

// 	for _, tt := range tests {
// 		t.Run(tt.Name, func(t *testing.T) {
// 			got, _ := ValidatePanCard(tt.Args.Pannumber)
// 			assert.Equal(t, tt.Want, got)
// 		})
// 	}
// }

func TestArticalAPI(t *testing.T) {
	url := "http://test-api.af-south-1.elasticbeanstalk.com/api/login-user"
	type Params struct {
		Email           string `json:"email"`
		Password        string `json:"passowrd"`
		ResourceAddress string `json:"resource_address"`
	}
	tests := []struct {
		Name         string
		Url          string
		RequestBody  Params
		ExpectedCode int
	}{
		{
			Name: "Aniket Suryawanshi",
			Url:  url,
			RequestBody: Params{
				Email:           "aniketsuryawanshixz1@gmail.com",
				Password:        "aniket123",
				ResourceAddress: "http://test-dashboard.af-south-1.elasticbeanstalk.com/admin/forgot-password/",
			},
			ExpectedCode: 401,
		},
	}

	for _, tt := range tests {
		t.Run(tt.Name, func(t *testing.T) {
			c := http.Client{}

			requestBody, marErr := json.Marshal(tt.RequestBody)
			if marErr != nil {
				fmt.Println("Marshal Error : ", marErr)
			}
			response, err := c.Post(url, "application/json", bytes.NewBuffer(requestBody))

			if err != nil {
				panic(err)
			}
			defer response.Body.Close()
			responseBody, _ := io.ReadAll(response.Body)

			fmt.Println(string(responseBody))
			assert.Equal(t, tt.ExpectedCode, response.StatusCode)
			// if assert.Equal(t, tt.ExpectedCode, response.StatusCode) {
			//

			// 	if bodyErr != nil {
			// 		fmt.Println(bodyErr)
			// 	}

			// 	article := Articles{}

			// 	json.Unmarshal(responseBody, &article)

			// 	assert.Equal(t, "Best Programer", article.Name)
			// }
		})
	}
}

type Address struct {
	City string
}

type Users struct {
	Name string
}

func BenchmarkSetUsers(b *testing.B) {
	count := make(chan int)

	// var myFunc func(chanData chan int)

	myFunc := func(chanData chan int) {
		chanData <- 3
	}

	go myFunc(count)

	for {
		val, ok := <-count
		if !ok {
			break
		}
		fmt.Println(val)
	}
}
