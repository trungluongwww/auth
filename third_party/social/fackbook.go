package social

import (
	"encoding/json"
	fb "github.com/huandu/facebook/v2"
)

type FacebookInfo struct {
	ID     string `json:"id"`
	Name   string `json:"name"`
	Email  string `json:"email"`
	Gender string `json:"gender"`
	Token  string `json:"token"`
}

func (*socialImpl) GetFacebookInfo(accessToken string) (res *FacebookInfo, err error) {
	params := fb.Params{
		"access_token": accessToken,
		"fields":       []string{"id", "name", "email", "gender"},
	}

	fbResponse, err := fb.Get("/me", params)
	if err != nil {
		return nil, err
	}

	data, err := json.Marshal(fbResponse)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(data, &res)
	if err != nil {
		return nil, err
	}

	res.Token = accessToken

	return res, nil
}
