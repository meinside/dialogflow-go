package dialogflow

// https://dialogflow.com/docs/reference/language
type LanguageTag string

const (
	BrazilianProtuguese LanguageTag = "pt-BR"
	ChineseCantonese    LanguageTag = "zh-HK"
	ChineseSimplified   LanguageTag = "zh-CN"
	ChineseTraditional  LanguageTag = "zh-TW"
	Dutch               LanguageTag = "nl"
	English             LanguageTag = "en"
	EnglishAustralian   LanguageTag = "en-AU"
	EnglishCanadian     LanguageTag = "en-CA"
	EnglishUK           LanguageTag = "en-GB"
	EnglishUS           LanguageTag = "en-US"
	French              LanguageTag = "fr"
	FrenchCanadian      LanguageTag = "fr-CA"
	FrenchFrench        LanguageTag = "fr-FR"
	German              LanguageTag = "de"
	Italian             LanguageTag = "it"
	Japanese            LanguageTag = "ja"
	Korean              LanguageTag = "ko"
	Portugese           LanguageTag = "pt"
	Russian             LanguageTag = "ru"
	Spanish             LanguageTag = "es"
	Ukrainian           LanguageTag = "uk"
)

// https://dialogflow.com/docs/reference/agent/#status_and_error_codes
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

// https://dialogflow.com/docs/reference/agent/#status_object
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
// https://dialogflow.com/docs/reference/agent/query#query_parameters_and_json_fields
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
			Messages []Message `json:"messages"` // https://dialogflow.com/docs/reference/agent/query#message_objects
		} `json:"fulfillment"`
		Score    float32  `json:"score"`
		Metadata Metadata `json:"metadata"`
	} `json:"result"`
	Status    StatusObject `json:"status"`
	SessionId string       `json:"sessionId"`
}

// https://dialogflow.com/docs/reference/agent/query#message_objects
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
	Platform string      `json:"platform,omitempty"`
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

type TextResponseMessageObject struct {
	MessageObject
	Speech []string `json:"speech"`
}

// helper function for creating a TextResponseMessage
func TextResponseMessage(platform string, speech []string) Message {
	msg := map[string]interface{}{
		"type":   TextResponseMessageObjectType,
		"speech": speech,
	}
	if len(platform) > 0 {
		msg["platform"] = platform
	}

	return Message(msg)
}

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

type QuickRepliesMessageObject struct {
	MessageObject
	Title   string   `json:"title"`
	Replies []string `json:"replies"`
}

type ImageMessageObject struct {
	MessageObject
	ImageUrl string `json:"imageUrl"`
}

type CustomPayloadMessageObject struct {
	MessageObject
	Payload interface{} `json:"payload"`
}

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
// https://dialogflow.com/docs/reference/agent/entities#entity_object
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
// https://dialogflow.com/docs/reference/agent/userentities#user_entity_object
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
// https://dialogflow.com/docs/reference/agent/intents#intent_object
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
	Auto                  bool             `json:"auto,omitempty"`
	Contexts              []string         `json:"contexts,omitempty"`
	Templates             []string         `json:"templates,omitempty"`
	UserSays              []UserSays       `json:"userSays,omitempty"`
	Responses             []IntentResponse `json:"responses,omitempty"`
	Priority              int              `json:"priority,omitempty"`
	WebhookUsed           bool             `json:"webhookUsed,omitempty"`
	WebhookForSlotFilling bool             `json:"webhookForSlotFilling,omitempty"`
	FallbackIntent        bool             `json:"fallbackIntent,omitempty"`
	CortanaCommand        CortanaCommand   `json:"cortanaCommand,omitempty"`
	Events                []struct {
		Name string `json:"name"`
	} `json:"events,omitempty"`
}

type UserSays struct {
	Id         string         `json:"id,omitempty"`
	Data       []UserSaysData `json:"data"`
	IsTemplate bool           `json:"isTemplate"`
	Count      int            `json:"count"`
}

type UserSaysData struct {
	Text        string `json:"text"`
	Meta        string `json:"meta,omitempty"`
	Alias       string `json:"alias,omitempty"`
	UserDefined bool   `json:"userDefined,omitempty"`
}

type IntentResponse struct {
	Action                   string                    `json:"action,omitempty"`
	ResetContexts            bool                      `json:"resetContexts,omitempty"`
	AffectedContexts         []IntentAffectedContext   `json:"affectedContexts,omitempty"`
	Parameters               []IntentResponseParameter `json:"parameters,omitempty"`
	Messages                 []Message                 `json:"messages,omitempty"`
	DefaultResponsePlatforms []string                  `json:"defaultResponsePlatforms,omitempty"`
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
	DataType     string   `json:"dataType,omitempty"`
	Prompts      []string `json:"prompts,omitempty"`
	IsList       bool     `json:"isList"`
}

type CortanaCommand struct {
	NavigateOrService string `json:"navigateOrService"`
	Target            string `json:"target"`
}

///////////////////////////////
//
// https://dialogflow.com/docs/reference/agent/contexts#context_object
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
