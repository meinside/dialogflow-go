package dialogflow

// https://dialogflow.com/docs/reference/agent/query

import (
	"encoding/json"
)

// Query text.
func (c *Client) QueryText(query QueryRequest) (result QueryResponse, err error) {
	var bytes []byte
	if bytes, err = c.httpPost("query", nil, nil, query); err == nil {
		if err = json.Unmarshal(bytes, &result); err == nil {
			return result, nil
		}
	}

	return QueryResponse{}, err
}

/*
// Query voice in .wav(16000Hz, signed PCM, 16 bit, mono) format.
//
// - file: filepath(string) or opened file (*os.File)
//
// NOTE: this api requires paid plan
func (c *Client) QueryVoice(query QueryRequest, file interface{}) (result QueryResponse, err error) {
	var bytes []byte
	if bytes, err = c.httpPostMultipart(
		"query",
		nil,
		map[string]interface{}{
			"request": query,
		},
		map[string]interface{}{
			"voiceData": file,
		},
	); err == nil {
		if err = json.Unmarshal(bytes, &result); err == nil {
			return result, nil
		}
	}

	return QueryResponse{}, err
}
*/
