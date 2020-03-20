package srv

import (
	"encoding/json"
	"net/http"

	"cartoon-user/model"
	"golang.org/x/net/context"
)

type loginRequest struct {
	DeviceToken string `json:"device_token"`
}

type loginResponse struct {
	User model.User `json:"user"`
	Token string `json:"token"`
	TokenType string `json:"token_type"`
}

func decodeLoginRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	var req loginRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		return nil,err
	}
	return req, nil
}

func encodeLoginResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	return json.NewEncoder(w).Encode(response)
}