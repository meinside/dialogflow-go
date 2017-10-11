package dialogflow

// https://dialogflow.com/docs/reference/agent/entities

import (
	"encoding/json"
	"fmt"
)

// Get all entities.
func (c *Client) AllEntities() (result Entities, err error) {
	var bytes []byte
	if bytes, err = c.httpGet("entities", nil, nil); err == nil {
		if err = json.Unmarshal(bytes, &result); err == nil {
			return result, nil
		}
	}

	return Entities{}, err
}

// Get an entitiy with given eid.
func (c *Client) Entity(eidOrName string) (result EntityObject, err error) {
	var bytes []byte
	if bytes, err = c.httpGet(fmt.Sprintf("entities/%s", eidOrName), nil, nil); err == nil {
		if err = json.Unmarshal(bytes, &result); err == nil {
			return result, nil
		}
	}

	return EntityObject{}, err
}

// Create a new entity.
//
// (do not fill Id, IsEnum, AutomatedExpansion value in EntityObject)
func (c *Client) CreateEntity(entity EntityObject) (result ApiResponse, err error) {
	var bytes []byte
	if bytes, err = c.httpPost("entities", nil, nil, entity); err == nil {
		if err = json.Unmarshal(bytes, &result); err == nil {
			return result, nil
		}
	}

	return ApiResponse{}, err
}

// Add entires to an entity.
func (c *Client) AddEntityEntries(eidOrName string, entries []EntityEntryObject) (result ApiResponse, err error) {
	var bytes []byte
	if bytes, err = c.httpPost(fmt.Sprintf("entities/%s/entries", eidOrName), nil, nil, entries); err == nil {
		if err = json.Unmarshal(bytes, &result); err == nil {
			return result, nil
		}
	}

	return ApiResponse{}, err
}

// Create/update entities.
//
// (do not fill Id, IsEnum, AutomatedExpansion value in EntityObject)
func (c *Client) CreateOrUpdateEntities(entities []EntityObject) (result ApiResponse, err error) {
	var bytes []byte
	if bytes, err = c.httpPut("entities", nil, entities); err == nil {
		if err = json.Unmarshal(bytes, &result); err == nil {
			return result, nil
		}
	}

	return ApiResponse{}, err
}

// Update an entity.
func (c *Client) UpdateEntity(eidOrName string, entity EntityObject) (result ApiResponse, err error) {
	var bytes []byte
	if bytes, err = c.httpPut(fmt.Sprintf("entities/%s", eidOrName), nil, entity); err == nil {
		if err = json.Unmarshal(bytes, &result); err == nil {
			return result, nil
		}
	}

	return ApiResponse{}, err
}

// Update entries of an entity.
func (c *Client) UpdateEntityEntries(eidOrName string, entries []EntityEntryObject) (result ApiResponse, err error) {
	var bytes []byte
	if bytes, err = c.httpPut(fmt.Sprintf("entities/%s/entries", eidOrName), nil, entries); err == nil {
		if err = json.Unmarshal(bytes, &result); err == nil {
			return result, nil
		}
	}

	return ApiResponse{}, err
}

// Delete an entity.
func (c *Client) DeleteEntity(eidOrName string) (result ApiResponse, err error) {
	var bytes []byte
	if bytes, err = c.httpDelete(fmt.Sprintf("entities/%s", eidOrName), nil, nil, nil); err == nil {
		if err = json.Unmarshal(bytes, &result); err == nil {
			return result, nil
		}
	}

	return ApiResponse{}, err
}

// Delete entries of an entity.
func (c *Client) DeleteEntityEntries(eidOrName string, entries []string) (result ApiResponse, err error) {
	var bytes []byte
	if bytes, err = c.httpDelete(fmt.Sprintf("entities/%s/entries", eidOrName), nil, nil, entries); err == nil {
		if err = json.Unmarshal(bytes, &result); err == nil {
			return result, nil
		}
	}

	return ApiResponse{}, err
}
