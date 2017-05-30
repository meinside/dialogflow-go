package apiai

// https://docs.api.ai/docs/contexts

import (
	"encoding/json"
	"fmt"
)

// Get all contexts with given session id.
func (c *Client) AllContexts(sid string) (result []ContextObject, err error) {
	var bytes []byte
	if bytes, err = c.httpGet("contexts", nil, map[string]string{"sessionId": sid}); err == nil {
		if err = json.Unmarshal(bytes, &result); err == nil {
			return result, nil
		}
	}

	return []ContextObject{}, err
}

// Get a context.
func (c *Client) Context(sid, contextName string) (result ContextObject, err error) {
	var bytes []byte
	if bytes, err = c.httpGet(fmt.Sprintf("contexts/%s", contextName), nil, map[string]string{"sessionId": sid}); err == nil {
		if err = json.Unmarshal(bytes, &result); err == nil {
			return result, nil
		}
	}

	return ContextObject{}, err
}

// Create contexts.
func (c *Client) CreateContexts(sid string, contexts []ContextObject) (result ContextResponseCreated, err error) {
	var bytes []byte
	if bytes, err = c.httpPost("contexts", nil, map[string]string{"sessionId": sid}, contexts); err == nil {
		if err = json.Unmarshal(bytes, &result); err == nil {
			return result, nil
		}
	}

	return ContextResponseCreated{}, err
}

// Delete all contexts.
func (c *Client) DeleteContexts(sid string) (result ContextResponseDeleted, err error) {
	var bytes []byte
	if bytes, err = c.httpDelete("contexts", nil, map[string]string{"sessionId": sid}, nil); err == nil {
		if err = json.Unmarshal(bytes, &result); err == nil {
			return result, nil
		}
	}

	return ContextResponseDeleted{}, err
}

// Delete a context.
func (c *Client) DeleteContext(sid, contextName string) (result ApiResponse, err error) {
	var bytes []byte
	if bytes, err = c.httpDelete(fmt.Sprintf("contexts/%s", contextName), nil, map[string]string{"sessionId": sid}, nil); err == nil {
		if err = json.Unmarshal(bytes, &result); err == nil {
			return result, nil
		}
	}

	return ApiResponse{}, err
}
