package repository

import (
	"auth/internal/repository/api/users/dto"
	"auth/internal/service/auth/adapters/repository"
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (r *Repository) ReadUserByCredetinals(ctx context.Context, params *repository.ReadUserByCredetinalsParams) (user *repository.User, err error) {
	request := &dto.ReadUserByCredetinalsRequest{
		Login: params.Login,
		Pass:  params.Pass,
	}

	//var id, _ = uuid.Parse("8d6f93fa-b198-413c-a40d-7fc462083e52")

	// response := &dto.ReadUserByCredetinalsResponse{
	// 	Id:         id,
	// 	Login:      "Léa.Fernandez",
	// 	Name:       "Lea",
	// 	MiddleName: "N",
	// 	Surname:    "Fernandez",
	// 	Phone:      "123456789",
	// 	City:       "X",
	// 	CreatedAt:  time.Now(),
	// 	ModifiedAt: time.Now(),
	// }

	response := &dto.ReadUserByCredetinalsResponse{}
	_, err = sendRequest(r.baseUrl+"/user/creds", http.MethodPost, "application/json", request, response)
	if err != nil {
		return
	}

	fmt.Println("api->user->auth: ReadUserByCredetinals(): SUCCESS! Name = " + response.Name)

	return &repository.User{
		Id:         response.Id,
		Login:      response.Login,
		Name:       response.Name,
		MiddleName: response.MiddleName,
		Surname:    response.Surname,
		Phone:      response.Phone,
		City:       response.City,
		CreatedAt:  response.CreatedAt,
		ModifiedAt: response.ModifiedAt,
	}, nil
}

func sendRequest(url, method, contentType string, data interface{}, response interface{}) (code int, err error) {
	bodyBuffer, err := prepareBody(data)
	if err != nil {
		return
	}
	req, err := http.NewRequest(method, url, bodyBuffer)
	if err != nil {
		return 0, err
	}

	req.Header.Set("Content-Type", contentType)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return 0, err
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return 0, fmt.Errorf("no body in response")
	}

	if resp.StatusCode > 201 {
		return resp.StatusCode, fmt.Errorf("bad response code from server %s - code %d, body of response: %s", url, code, string(body))
	}

	if response != nil {
		err = json.Unmarshal(body, response)
		if err != nil {
			return 0, fmt.Errorf("error decode response from %s - error %s, body of response: %s", url, err.Error(), string(body))
		}
	}

	return resp.StatusCode, err
}

func prepareBody(data interface{}) (buffer *bytes.Buffer, err error) {
	var sendData []byte
	if data != nil {
		sendData, err = json.Marshal(data)
		if err != nil {
			return
		}
		buffer = bytes.NewBuffer(sendData)
	}
	return
}
