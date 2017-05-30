# Simple Go library for api.ai

## Install

```
$ go get -u github.com/meinside/api.ai-go
```

## Usage / Example

```go
package main

import (
	"fmt"

	ai "github.com/meinside/api.ai-go"
)

func main() {
	// variables for test
	token := "000000aaaaaa111111bbbbbbcccccc" // XXX - your token here
	sessionId := "test_0123456789"
	newEntityName := "test-new-entity"

	// setup a client
	client := ai.NewClient(token)
	//client.Verbose = false
	client.Verbose = true // for verbose messages

	/////////////////////////////////////////////
	// Query
	//
	////////////////
	// query text
	if response, err := client.QueryText(ai.QueryRequest{
		Query:     []string{"May I test?"},
		SessionId: sessionId,
		Language:  ai.English,
	}); err == nil {
		fmt.Printf(">>> response = %+v\n", response)
	} else {
		fmt.Printf("*** error: %s\n", err)
	}

	/////////////////////////////////////////////
	// Entity
	//
	////////////////
	// create entity
	if response, err := client.CreateEntity(ai.EntityObject{
		Name: newEntityName,
		Entries: []ai.EntityEntryObject{
			ai.EntityEntryObject{
				Value: "Roadhog",
				Synonyms: []string{
					"Roadhog",
					"pig",
					"bacon",
				},
			},
		},
	}); err == nil {
		fmt.Printf(">>> created entity = %+v\n", response)
	} else {
		fmt.Printf("*** error: %s\n", err)
	}

	/////////////////////////////////////////////
	// Intents
	//
	////////////////
	// get all intents
	if intents, err := client.AllIntents(); err == nil {
		fmt.Printf(">>> all intents = %+v\n", intents)

		for _, intent := range intents {
			////////////////
			// get an intent
			if response, err := client.Intent(intent.Id); err == nil {
				fmt.Printf(">>> intent = %+v\n", response)
			} else {
				fmt.Printf("*** error: %s\n", err)
			}
		}

		////////////////
		// create an intent
		if response, err := client.CreateIntent(ai.IntentObject{
			Name:     "test-intent",
			Auto:     true,
			Contexts: []string{},
			Templates: []string{
				"test 1",
				"test 2",
			},
			UserSays: []ai.UserSays{
				ai.UserSays{
					Data: []ai.UserSaysData{
						ai.UserSaysData{
							Text: "test says 1",
						},
					},
				},
			},
			Responses: []ai.IntentResponse{
				ai.IntentResponse{
					Action:        "test-action1",
					ResetContexts: false,
					AffectedContexts: []ai.IntentAffectedContext{
						ai.IntentAffectedContext{
							Name:     "test1",
							Lifespan: 1,
						},
					},
				},
			},
		}); err == nil {
			fmt.Printf(">>> created intent = %+v\n", response)

			////////////////
			// update an intent
			if response, err := client.UpdateIntent(response.Id, ai.IntentObject{
				Name:     "test-intent-updated",
				Auto:     true,
				Contexts: []string{},
				Templates: []string{
					"test 1",
					"test 2",
				},
				UserSays: []ai.UserSays{
					ai.UserSays{
						Data: []ai.UserSaysData{
							ai.UserSaysData{
								Text: "test says 1",
							},
						},
					},
				},
				Responses: []ai.IntentResponse{
					ai.IntentResponse{
						Action:        "test-action1",
						ResetContexts: false,
						AffectedContexts: []ai.IntentAffectedContext{
							ai.IntentAffectedContext{
								Name:     "test1",
								Lifespan: 1,
							},
						},
					},
				},
			}); err == nil {
				fmt.Printf(">>> updated intent = %+v\n", response)
			} else {
				fmt.Printf("*** error: %s\n", err)
			}

			////////////////
			// delete an intent
			if response, err := client.DeleteIntent(response.Id); err == nil {
				fmt.Printf(">>> deleted intent = %+v\n", response)
			} else {
				fmt.Printf("*** error: %s\n", err)
			}
		} else {
			fmt.Printf("*** error: %s\n", err)
		}
	} else {
		fmt.Printf("*** error: %s\n", err)
	}

	/////////////////////////////////////////////
	// Contexts
	//
	////////////////
	// create a context
	if response, err := client.CreateContexts(sessionId, []ai.ContextObject{
		ai.ContextObject{
			Name:     "test-context",
			Lifespan: 3,
			Parameters: map[string]interface{}{
				"name":  "some-name",
				"value": "some-value",
			},
		},
	}); err == nil {
		fmt.Printf(">>> created context = %+v\n", response)
	} else {
		fmt.Printf("*** error: %s\n", err)
	}
	////////////////
	// get all contexts
	if contexts, err := client.AllContexts(sessionId); err == nil {
		fmt.Printf(">>> all contexts = %+v\n", contexts)

		for _, context := range contexts {
			////////////////
			// get a context
			if response, err := client.Context(sessionId, context.Name); err == nil {
				fmt.Printf(">>> context = %+v\n", response)
			} else {
				fmt.Printf("*** error: %s\n", err)
			}
		}
	} else {
		fmt.Printf("*** error: %s\n", err)
	}
	////////////////
	// delete all contexts
	if response, err := client.DeleteContexts(sessionId); err == nil {
		fmt.Printf(">>> deleted contexts = %+v\n", response)
	} else {
		fmt.Printf("*** error: %s\n", err)
	}
}
```

## License

MIT

