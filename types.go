package apiai

// https://docs.api.ai/docs/languages
type LanguageTag string

const (
	BrazilianProtuguese LanguageTag = "pt-BR"
	ChineseCantonese    LanguageTag = "zh-HK"
	ChineseSimplified   LanguageTag = "zh-CN"
	ChineseTraditional  LanguageTag = "zh-TW"
	Dutch               LanguageTag = "nl"
	English             LanguageTag = "en"
	French              LanguageTag = "fr"
	German              LanguageTag = "de"
	Italian             LanguageTag = "it"
	Japanese            LanguageTag = "ja"
	Korean              LanguageTag = "ko"
	Portugese           LanguageTag = "pt"
	Russian             LanguageTag = "ru"
	Spanish             LanguageTag = "es"
	Ukrainian           LanguageTag = "uk"
)

// https://docs.api.ai/docs/status-and-error-codes
type ErrorType string

const (
	Success ErrorType = "success"

	Deprecated      ErrorType = "deprecated"
	BadRequest      ErrorType = "bad_request"
	Unauthorized    ErrorType = "unauthorized"
	NotFound        ErrorType = "not_found"
	NotAllowed      ErrorType = "not_allowed"
	NotAcceptable   ErrorType = "not_acceptable"
	Conflict        ErrorType = "conflict"
	TooManyRequests ErrorType = "too_many_requests"
)

// https://docs.api.ai/docs/status-object
type StatusObject struct {
	Code         int       `json:"code"`
	ErrorType    ErrorType `json:"errorType"`
	ErrorId      string    `json:"errorId,omitempty"`
	ErrorDetails string    `json:"errorDetails,omitempty"`
}

type ApiResponse struct {
	Id     string       `json:"id,omitempty"`
	Status StatusObject `json:"status,omitempty"`
}

///////////////////////////////
//
// https://docs.api.ai/docs/query
type QueryRequest struct {
	Query         []string        `json:"query,omitempty"`
	SessionId     string          `json:"sessionId"`
	Language      LanguageTag     `json:"lang"`
	Contexts      []ContextObject `json:"contexts,omitempty"`
	ResetContexts bool            `json:"resetContexts,omitempty"`
	Entities      []Entity        `json:"entities,omitempty"`
	Timezone      string          `json:"timezone,omitempty"`
	Location      struct {
		Latitude  float32 `json:"latitude"`
		Longitude float32 `json:"longitude"`
	} `json:"location,omitempty"`
	OriginalRequest struct {
		Source string      `json:"source"`
		Data   interface{} `json:"data"`
	} `json:"originalRequest"`
}

type Entity struct {
	Id      string `json:"id"`
	Name    string `json:"name"`
	Count   int    `json:"count"`
	Preview string `json:"preview"`
}

type Entities struct {
	Entities []Entity     `json:"entities"`
	Status   StatusObject `json:"status"`
}

type QueryResponse struct {
	ApiResponse

	Timestamp string      `json:"timestamp"`
	Language  LanguageTag `json:"lang"`
	Result    struct {
		Source           string                 `json:"source"`
		ResolvedQuery    string                 `json:"resolvedQuery"`
		Action           string                 `json:"action"`
		ActionIncomplete bool                   `json:"actionIncomplete"`
		Parameters       map[string]interface{} `json:"parameters"`
		Contexts         []ContextObject        `json:"contexts"`
		Fulfillment      struct {
			Speech   string    `json:"speech"`
			Messages []Message `json:"messages"` // https://docs.api.ai/docs/query#section-message-objects
		} `json:"fulfillment"`
		Score    float32  `json:"score"`
		Metadata Metadata `json:"metadata"`
	} `json:"result"`
	Status    StatusObject `json:"status"`
	SessionId string       `json:"sessionId"`
}

// https://docs.api.ai/docs/query#section-message-objects
type MessageType int

const (
	TextResponseMessageObjectType  MessageType = 0
	CardMessageObjectType          MessageType = 1
	QuickRepliesMessageObjectType  MessageType = 2
	ImageMessageObjectType         MessageType = 3
	CustomPayloadMessageObjectType MessageType = 4
	NoSuchObjectType               MessageType = -1
)

type MessageObject struct {
	Type     MessageType `json:"type"`
	Platform string      `json:"platform"`
}

type Message map[string]interface{}

func (f Message) Type() MessageType {
	if v, exists := map[string]interface{}(f)["type"]; exists {
		if v, ok := v.(MessageType); ok {
			return v
		}
	}
	return NoSuchObjectType
}

// https://docs.api.ai/docs/intents#section-text-response-message-object
type TextResponseMessageObject struct {
	MessageObject
	Speech []string `json:"speech"`
}

// https://docs.api.ai/docs/intents#section-card-message-object
type CardMessageObject struct {
	MessageObject
	Title    string              `json:"title"`
	Subtitle string              `json:"subtitle"`
	Buttons  []CardMessageButton `json:"buttons"`
}

type CardMessageButton struct {
	Text     string `json:"text"`
	Postback string `json:"postback"`
}

// https://docs.api.ai/docs/intents#section-quick-replies-message-object
type QuickRepliesMessageObject struct {
	MessageObject
	Title   string   `json:"title"`
	Replies []string `json:"replies"`
}

// https://docs.api.ai/docs/intents#section-image-message-object
type ImageMessageObject struct {
	MessageObject
	ImageUrl string `json:"imageUrl"`
}

// https://docs.api.ai/docs/intents#section-custom-payload-message-object
type CustomPayloadMessageObject struct {
	MessageObject
	Payload interface{} `json:"payload"`
}

// https://docs.api.ai/docs/intents#section-text-response-message-object
func (f Message) ToTextResponseMessage() TextResponseMessageObject {
	m := map[string]interface{}(f)

	var typ MessageType = -1
	var platform string = ""
	var speech []string = nil

	if v, exists := m["type"]; exists {
		switch v.(type) {
		case MessageType:
			typ = v.(MessageType)
		}
	}
	if v, exists := m["platform"]; exists {
		switch v.(type) {
		case string:
			platform = v.(string)
		}
	}
	if v, exists := m["speech"]; exists {
		switch v.(type) {
		case []string:
			speech = v.([]string)
		case string:
			speech = []string{v.(string)}
		}
	}

	return TextResponseMessageObject{
		MessageObject: MessageObject{
			Type:     typ,
			Platform: platform,
		},
		Speech: speech,
	}
}

// https://docs.api.ai/docs/intents#section-card-message-object
func (f Message) ToCardMessage() CardMessageObject {
	m := map[string]interface{}(f)

	var typ MessageType = -1
	var platform string = ""
	var title string = ""
	var subtitle string = ""
	var buttons []CardMessageButton = []CardMessageButton{}

	if v, exists := m["type"]; exists {
		switch v.(type) {
		case MessageType:
			typ = v.(MessageType)
		}
	}
	if v, exists := m["platform"]; exists {
		switch v.(type) {
		case string:
			platform = v.(string)
		}
	}
	if v, exists := m["title"]; exists {
		switch v.(type) {
		case string:
			title = v.(string)
		}
	}
	if v, exists := m["subtitle"]; exists {
		switch v.(type) {
		case string:
			subtitle = v.(string)
		}
	}
	if v, exists := m["buttons"]; exists {
		switch v.(type) {
		case []map[string]string:
			for _, m := range v.([]map[string]string) {
				text := ""
				postback := ""
				if v, exists := m["text"]; exists {
					text = v
				}
				if v, exists := m["postback"]; exists {
					postback = v
				}
				buttons = append(buttons, CardMessageButton{
					Text:     text,
					Postback: postback,
				})
			}
		}
	}

	return CardMessageObject{
		MessageObject: MessageObject{
			Type:     typ,
			Platform: platform,
		},
		Title:    title,
		Subtitle: subtitle,
		Buttons:  buttons,
	}
}

// https://docs.api.ai/docs/intents#section-quick-replies-message-object
func (f Message) ToQuickRepliesMessage() QuickRepliesMessageObject {
	m := map[string]interface{}(f)

	var typ MessageType = -1
	var platform string = ""
	var title string = ""
	var replies []string = []string{}

	if v, exists := m["type"]; exists {
		switch v.(type) {
		case MessageType:
			typ = v.(MessageType)
		}
	}
	if v, exists := m["platform"]; exists {
		switch v.(type) {
		case string:
			platform = v.(string)
		}
	}
	if v, exists := m["title"]; exists {
		switch v.(type) {
		case string:
			title = v.(string)
		}
	}
	if v, exists := m["replies"]; exists {
		switch v.(type) {
		case []string:
			for _, s := range v.([]string) {
				replies = append(replies, s)
			}
		}
	}

	return QuickRepliesMessageObject{
		MessageObject: MessageObject{
			Type:     typ,
			Platform: platform,
		},
		Title:   title,
		Replies: replies,
	}
}

// https://docs.api.ai/docs/intents#section-image-message-object
func (f Message) ToImageMessage() ImageMessageObject {
	m := map[string]interface{}(f)

	var typ MessageType = -1
	var platform string = ""
	var imageUrl string = ""

	if v, exists := m["type"]; exists {
		switch v.(type) {
		case MessageType:
			typ = v.(MessageType)
		}
	}
	if v, exists := m["platform"]; exists {
		switch v.(type) {
		case string:
			platform = v.(string)
		}
	}
	if v, exists := m["imageUrl"]; exists {
		switch v.(type) {
		case string:
			imageUrl = v.(string)
		}
	}

	return ImageMessageObject{
		MessageObject: MessageObject{
			Type:     typ,
			Platform: platform,
		},
		ImageUrl: imageUrl,
	}
}

// https://docs.api.ai/docs/intents#section-custom-payload-message-object
func (f Message) ToCustomPayloadMessage() CustomPayloadMessageObject {
	m := map[string]interface{}(f)

	var typ MessageType = -1
	var platform string = ""
	var payload interface{} = nil

	if v, exists := m["type"]; exists {
		switch v.(type) {
		case MessageType:
			typ = v.(MessageType)
		}
	}
	if v, exists := m["platform"]; exists {
		switch v.(type) {
		case string:
			platform = v.(string)
		}
	}
	if v, exists := m["payload"]; exists {
		payload = v
	}

	return CustomPayloadMessageObject{
		MessageObject: MessageObject{
			Type:     typ,
			Platform: platform,
		},
		Payload: payload,
	}
}

type Metadata struct {
	IntentId                  string `json:"intentId"`
	WebhookUsed               string `json:"webhookUsed"`
	WebhookForSlotFillingUsed string `json:"webhookForSlotFillingUsed"`
	WebhookResponseTime       int    `json:"webhookResponseTime"`
	IntentName                string `json:"intentName"`
}

///////////////////////////////
//
// https://docs.api.ai/docs/entities
type EntityObject struct {
	ApiResponse

	Name               string              `json:"name"`
	Entries            []EntityEntryObject `json:"entries"`
	IsEnum             bool                `json:"isEnum,omitempty"`
	AutomatedExpansion bool                `json:"automatedExpansion,omitempty"`
}

type EntityEntryObject struct {
	Value    string   `json:"value"`
	Synonyms []string `json:"synonyms"`
}

///////////////////////////////
//
// https://docs.api.ai/docs/userentities
type UserEntityObject struct {
	ApiResponse

	SessionId string              `json:"sessionId"`
	Name      string              `json:"name"`
	Extend    bool                `json:"extend,omitempty"`
	Entries   []EntityEntryObject `json:"entries"`
}

type NewUserEntitiesObject struct {
	SessionId string             `json:"sessionId"`
	Entities  []UserEntityObject `json:"entities"`
}

///////////////////////////////
//
// https://docs.api.ai/docs/intents
type Intent struct {
	Id             string            `json:"id"`
	Name           string            `json:"name"`
	ContextIn      []string          `json:"contextIn"`
	ContextOut     []ContextOut      `json:"contextOut"`
	Actions        []string          `json:"actions"`
	Parameters     []IntentParameter `json:"parameters"`
	Priority       int               `json:"priority"`
	FallbackIntent bool              `json:"fallbackIntent"`
}

type ContextOut struct {
	Name     string `json:"name"`
	Lifespan int    `json:"lifespan"`
}

type IntentParameter struct {
	Name         string   `json:"name"`
	Value        string   `json:"value"`
	DefaultValue string   `json:"defaultValue"`
	Required     bool     `json:"required"`
	DataType     string   `json:"dataType"`
	Prompts      []string `json:"prompts"`
}

type IntentObject struct {
	ApiResponse // may not present when api call was successful

	Name                  string           `json:"name"`
	Auto                  bool             `json:"auto"`
	Contexts              []string         `json:"contexts"`
	Templates             []string         `json:"templates"`
	UserSays              []UserSays       `json:"userSays"`
	Responses             []IntentResponse `json:"responses"`
	Priority              int              `json:"priority"`
	WebhookUsed           bool             `json:"webhookUsed"`
	WebhookForSlotFilling bool             `json:"webhookForSlotFilling"`
	FallbackIntent        bool             `json:"fallbackIntent"`
	CortanaCommand        CortanaCommand   `json:"cortanaCommand"`
	Events                []struct {
		Name string `json:"name"`
	} `json:"events"`
}

type UserSays struct {
	Id         string         `json:"id,omitempty"`
	Data       []UserSaysData `json:"data"`
	IsTemplate bool           `json:"isTemplate"`
	Count      int            `json:"count"`
}

type UserSaysData struct {
	Text        string `json:"text"`
	Meta        string `json:"meta"`
	Alias       string `json:"alias"`
	UserDefined bool   `json:"userDefined"`
}

type IntentResponse struct {
	Action                   string                    `json:"action"`
	ResetContexts            bool                      `json:"resetContexts"`
	AffectedContexts         []IntentAffectedContext   `json:"affectedContexts"`
	Parameters               []IntentResponseParameter `json:"parameters"`
	Messages                 []Message                 `json:"messages"`
	DefaultResponsePlatforms []string                  `json:"defaultResponsePlatforms"`
}

type IntentAffectedContext struct {
	Name     string `json:"name"`
	Lifespan int    `json:"lifespan"`
}

type IntentResponseParameter struct {
	Name         string   `json:"name"`
	Value        string   `json:"value"`
	DefaultValue string   `json:"defaultValue"`
	Required     bool     `json:"required"`
	DataType     string   `json:"dataType"`
	Prompts      []string `json:"prompts"`
	IsList       bool     `json:"isList"`
}

type CortanaCommand struct {
	NavigateOrService string `json:"navigateOrService"`
	Target            string `json:"target"`
}

///////////////////////////////
//
// https://docs.api.ai/docs/contexts
type ContextObject struct {
	Name       string                 `json:"name,omitempty"`
	Lifespan   int                    `json:"lifespan,omitempty"`
	Parameters map[string]interface{} `json:"parameters,omitempty"` // XXX - document's specification and its sample request is different...?
}

type ContextResponseCreated struct {
	ApiResponse // may not present when api call was successful

	Names []string `json:"names"`
}

type ContextResponseDeleted struct {
	ApiResponse // may not present when api call was successful

	Deleted []string `json:"deleted"`
}
