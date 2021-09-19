package api

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"time"
)

const (
	token_url = "https://accounts.spotify.com/api/token"
)

type AuthorizationResponse struct {
	AccessToken string `json:"access_token"`
	TokenType   string `json:"token_type"`
	ExpiresIn   int    `json:"expires_in"`
	ExpiresAt   time.Time    `json:"expires_at"`
}

// GenerateAuthorizeToken
//get Spotiy token
func GenerateAuthorizeToken(clientId, clientSecret string) string{
	//get auth string

	cli_cred_encoded := fmt.Sprintf("%s:%s", clientId, clientSecret )

	//convert client string to base64 encoding
	cli_cred_encoded_b64 := base64.StdEncoding.EncodeToString([]byte(cli_cred_encoded))

	//client authorization string
	authorizationStr := fmt.Sprintf("Basic %s", cli_cred_encoded_b64)

	return authorizationStr
}

// New
// New instance of authorization class
func New() *AuthorizationResponse {
	return &AuthorizationResponse{}
}

//Authorization
//authorise the Spotify Api
// Return error
func (s *AuthorizationResponse) Authorization(clientId, clientSecret string) error{
	//parse params in request
	params := url.Values{}
	params.Add("grant_type", `client_credentials`)
	reqbody := strings.NewReader(params.Encode())


	// create request
	req, err := http.NewRequest("POST", token_url, reqbody)
	if err != nil {
		return err
	}
	//set request headers
	req.Header.Set("Authorization", GenerateAuthorizeToken(clientId,clientSecret))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	//fmt.Println(req)
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	defer res.Body.Close()


	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return err
	}

	err = json.Unmarshal(body, &s);

	if err != nil {
		return err
	}
	return nil
}





