package apiai

// https://docs.api.ai/docs/languages
type LanguageTag string

const (
	BrazilianProtuguese LanguageTag = "pt-BR"
	ChineseCantonese    LanguageTag = "zh-HK"
	ChineseSimplified   LanguageTag = "zh-CN"
	ChineseTraditional  LanguageTag = "zh-TW"
	English             LanguageTag = "en"
	Dutch               LanguageTag = "nl"
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
	Confidence    float32         `json:"confidence,omitempty"`
	SessionId     string          `json:"sessionId"`
	Language      LanguageTag     `json:"lang"`
	Contexts      []ContextObject `json:"contexts,omitempty"`
	ResetContexts bool            `json:"resetContexts,omitempty"`
	Entities      []Entity        `json:"entities,omitempty"`
	Timezone      string          `json:"timezone,omitempty"`
	Location      Location        `json:"location,omitempty"`
}

type Entity struct {
	Id      string `json:"id"`
	Name    string `json:"name"`
	Count   int    `json:"count"`
	Preview string `json:"preview"`
}

type Location struct {
	Latitude  float32 `json:"latitude"`
	Longitude float32 `json:"longitude"`
}

type QueryResponse struct {
	ApiResponse

	Timestamp string       `json:"timestamp"`
	Result    QueryResult  `json:"result"`
	Status    StatusObject `json:"status"`
	SessionId string       `json:"sessionId"`
}

type QueryResult struct {
	Source           string                 `json:"source"`
	ResolvedQuery    string                 `json:"resolvedQuery"`
	Action           string                 `json:"action"`
	ActionIncomplete bool                   `json:"actionIncomplete"`
	Parameters       map[string]interface{} `json:"parameters"`
	Contexts         []ContextObject        `json:"contexts"`
	Fulfillment      Fulfillment            `json:"fulfillment"`
	Metadata         Metadata               `json:"metadata"`
}

type Fulfillment struct {
	Speech string  `json:"speech"`
	Score  float32 `json:"score"`
}

type Metadata struct {
	IntentId    string `json:"intentId"`
	WebhookUsed string `json:"webhookUsed"`
	IntentName  string `json:"intentName"`
}

///////////////////////////////
//
// https://docs.api.ai/docs/entities
type EntityObject struct {
	ApiResponse

	Name               string              `json:"name"`
	IsOverridable      bool                `json:"isOverridable"`
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
	SessionId string         `json:"sessionId"`
	Entities  []EntityObject `json:"entities"`
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
	ApiResponse

	Name             string           `json:"name"`
	Auto             bool             `json:"auto"`
	Contexts         []string         `json:"contexts"`
	Templates        []string         `json:"templates"`
	UserSays         []UserSays       `json:"userSays"`
	Responses        []IntentResponse `json:"responses"`
	Priority         int              `json:"priority"`
	WebhookUsed      bool             `json:"webhookUsed"`
	FallbackIntent   bool             `json:"fallbackIntent"`
	AssistantCommand AssistantCommand `json:"assistantCommand"`
	CortanaCommand   CortanaCommand   `json:"cortanaCommand"`
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
	Action           string                    `json:"action"`
	ResetContexts    bool                      `json:"resetContexts"`
	AffectedContexts []IntentAffectedContext   `json:"affectedContexts"`
	Parameters       []IntentResponseParameter `json:"parameters"`
	Speech           []string                  `json:"speech"`
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
}

type AssistantCommand struct {
	UrlCommand string `json:"urlCommand"`
	DoCommand  string `json:"doCommand"`
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
	Parameters map[string]interface{} `json:"parameters,omitempty"`
	Lifespan   int                    `json:"lifespan,omitempty"`
}

type ContextResponseCreated struct {
	ApiResponse

	Names []string `json:"names"`
}

type ContextResponseDeleted struct {
	ApiResponse

	Deleted []string `json:"deleted"`
}

type ContextParameter struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}
