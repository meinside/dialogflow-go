package apiai

// https://docs.api.ai/docs/userentities

// XXX - not working properly yet... something is missing in the docs

import (
	"encoding/json"
	"fmt"
)

// create new user entities
func (c *Client) CreateUserEntities(sessionId string, entities []EntityObject) (result ApiResponse, err error) {
	var bytes []byte
	if bytes, err = c.httpPost("userEntities", nil, nil, NewUserEntitiesObject{
		SessionId: sessionId,
		Entities:  entities,
	}); err == nil {
		if err = json.Unmarshal(bytes, &result); err == nil {
			return result, nil
		}
	}

	return ApiResponse{}, err
}

// update user entity
func (c *Client) UpdateUserEntity(name string, entity EntityObject) (result ApiResponse, err error) {
	var bytes []byte
	if bytes, err = c.httpPut(fmt.Sprintf("userEntities/%s", name), nil, entity); err == nil {
		if err = json.Unmarshal(bytes, &result); err == nil {
			return result, nil
		}
	}

	return ApiResponse{}, err
}

// get user entity
func (c *Client) UserEntity(name string) (result UserEntityObject, err error) {
	var bytes []byte
	if bytes, err = c.httpGet(fmt.Sprintf("userEntities/%s", name), nil, nil); err == nil {
		if err = json.Unmarshal(bytes, &result); err == nil {
			return result, nil
		}
	}

	return UserEntityObject{}, err
}

// delete user entity
func (c *Client) DeleteUserEntity(name string) (result ApiResponse, err error) {
	var bytes []byte
	if bytes, err = c.httpDelete(fmt.Sprintf("userEntities/%s", name), nil, nil, nil); err == nil {
		if err = json.Unmarshal(bytes, &result); err == nil {
			return result, nil
		}
	}

	return ApiResponse{}, err
}
