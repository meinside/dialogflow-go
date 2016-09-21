package apiai

// text to speech
//
// https://docs.api.ai/docs/tts
func (c *Client) Tts(text, language string) (wav []byte, err error) {
	return c.httpGet("tts", map[string]string{
		"Accept-language": language,
	}, map[string]string{
		"text": text,
	})
}
