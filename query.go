package apiai

import (
	"encoding/json"
)

// query text
//
// https://docs.api.ai/docs/query
func (c *Client) QueryText(query QueryRequest) (result QueryResponse, err error) {
	var bytes []byte
	if bytes, err = c.httpPost("query", nil, nil, query); err == nil {
		if err = json.Unmarshal(bytes, &result); err == nil {
			return result, nil
		}
	}

	return QueryResponse{}, err
}

// query voice in .wav(16000Hz, signed PCM, 16 bit, mono) format
//
// NOTE: this api requires paid plan
//
// https://docs.api.ai/docs/query
func (c *Client) QueryVoice(query QueryRequest, filepath string) (result QueryResponse, err error) {
	var bytes []byte
	if bytes, err = c.httpPostMultipart(
		"query",
		nil,
		map[string]interface{}{
			"request": query,
		},
		map[string]string{
			"voiceData": filepath,
		},
	); err == nil {
		if err = json.Unmarshal(bytes, &result); err == nil {
			return result, nil
		}
	}

	return QueryResponse{}, err
}
