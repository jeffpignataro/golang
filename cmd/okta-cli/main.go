package main

import (
	"encoding/json"
	"fmt"
	"os"
	"time"

	"github.com/jeffpignataro/golang/pkg/rest"
)

type OktaResponse struct {
	ExpiresAt    time.Time `json:"expiresAt"`
	Status       string    `json:"status"`
	SessionToken string    `json:"sessionToken"`
	Embedded     struct {
		User struct {
			ID              string    `json:"id"`
			PasswordChanged time.Time `json:"passwordChanged"`
			Profile         struct {
				Login     string `json:"login"`
				FirstName string `json:"firstName"`
				LastName  string `json:"lastName"`
				Locale    string `json:"locale"`
				TimeZone  string `json:"timeZone"`
			} `json:"profile"`
		} `json:"user"`
	} `json:"_embedded"`
	ErrorCode    string        `json:"errorCode"`
	ErrorSummary string        `json:"errorSummary"`
	ErrorLink    string        `json:"errorLink"`
	ErrorID      string        `json:"errorId"`
	ErrorCauses  []interface{} `json:"errorCauses"`
}

type DeviceResponse struct {
	UserCode                string `json:"user_code"`
	DeviceCode              string `json:"device_code"`
	Interval                int    `json:"interval"`
	VerificationURIComplete string `json:"verification_uri_complete"`
	VerificationURI         string `json:"verification_uri"`
	ExpiresIn               int    `json:"expires_in"`
	Error                   string `json:"error"`
	ErrorDescription        string `json:"error_description"`
}

func main() {
	token := os.Getenv("OKTA_TOKEN")
	pword := os.Getenv("OKTA_PASSWORD")
	authType := "SSWS"
	// ctx, client := okta.NewClient(os.Getenv("OKTA_TOKEN"))

	// fmt.Print(ctx)
	// fmt.Print(client)

	// jsonStr := `
	// {
	// 	"token": "26q43Ak9Eh04p7H6Nnx0m69JqYOrfVBY"
	// }`

	jsonStr := fmt.Sprintf(`
	{
		"username": "jeffp2662@gmail.com",
		"password": "%s"
	}`, pword)
	a := rest.NewAuthentication(true, &authType, &token)
	h := map[string]string{
		"Content-Type": "application/x-www-form-urlencoded",
	}
	r, err := rest.PostWithHeaders("https://dev-20308777-admin.okta.com/api/v1/authn", jsonStr, &a, &h)
	if err != nil {
		panic(err.Error())
	}

	u := fmt.Sprintf("https://dev-20308777.okta.com/oauth2/aus8l8qyb7b8zsZIM5d7/v1/authorize?client_id=%s&client_secret=%s&scope=okta.myAccount.profile.read", os.Getenv("CLIENT_ID"), os.Getenv("CLIENT_SECRET"))
	g, err := rest.PostWithHeaders(u, "", nil, &h)
	if err != nil {
		panic(err.Error())
	}

	resp := OktaResponse{}
	err = json.Unmarshal(r, &resp)
	if err != nil {
		panic(err.Error())
	}

	s := DeviceResponse{}
	json.Unmarshal(g, &s)

	// fmt.Printf("%+v\n", resp)
	fmt.Printf("%+v\n", s)
}
