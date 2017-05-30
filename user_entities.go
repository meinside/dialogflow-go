package apiai

// https://docs.api.ai/docs/userentities

import (
	"encoding/json"
	"fmt"
)

// Create new user entities.
func (c *Client) CreateUserEntities(sessionId string, entities []UserEntityObject) (result ApiResponse, err error) {
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

// Update user entity.
func (c *Client) UpdateUserEntity(name string, entity UserEntityObject) (result ApiResponse, err error) {
	var bytes []byte
	if bytes, err = c.httpPut(fmt.Sprintf("userEntities/%s", name), nil, entity); err == nil {
		if err = json.Unmarshal(bytes, &result); err == nil {
			return result, nil
		}
	}

	return ApiResponse{}, err
}

// Get a user entity.
func (c *Client) UserEntity(name string) (result UserEntityObject, err error) {
	var bytes []byte
	if bytes, err = c.httpGet(fmt.Sprintf("userEntities/%s", name), nil, nil); err == nil {
		if err = json.Unmarshal(bytes, &result); err == nil {
			return result, nil
		}
	}

	return UserEntityObject{}, err
}

// Delete user entity.
func (c *Client) DeleteUserEntity(name string) (result ApiResponse, err error) {
	var bytes []byte
	if bytes, err = c.httpDelete(fmt.Sprintf("userEntities/%s", name), nil, nil, nil); err == nil {
		if err = json.Unmarshal(bytes, &result); err == nil {
			return result, nil
		}
	}

	return ApiResponse{}, err
}
