# Simple Go library for Dialogflow

This repository is a rework of [api.ai-go](https://github.com/meinside/api.ai-go), due to the [change of service name from API.AI to Dialogflow](https://blog.dialogflow.com/post/apiai-new-name-dialogflow-new-features/).

Dialogflow's API reference is [here](https://dialogflow.com/docs/reference/agent/).

## Install

```
$ go get -u github.com/meinside/dialogflow-go
```

## Usage / Example

```go
package main

import (
	"fmt"

	df "github.com/meinside/dialogflow-go"
)

func main() {
	// variables for test
	token := "000000aaaaaa111111bbbbbbcccccc" // XXX - your token here
	sessionId := "test_0123456789"
	newEntityName := "test-new-entity"

	// setup a client
	client := df.NewClient(token)
	//client.Verbose = false
	client.Verbose = true // for verbose messages

	/////////////////////////////////////////////
	// Query
	//
	////////////////
	// query text
	if response, err := client.QueryText(df.QueryRequest{
		Query:     []string{"May I test?"},
		SessionId: sessionId,
		Language:  df.English,
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
	if response, err := client.CreateEntity(df.EntityObject{
		Name: newEntityName,
		Entries: []df.EntityEntryObject{
			df.EntityEntryObject{
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
		if response, err := client.CreateIntent(df.IntentObject{
			Name:     "test-intent",
			Auto:     true,
			Contexts: []string{},
			Templates: []string{
				"test 1",
				"test 2",
			},
			UserSays: []df.UserSays{
				df.UserSays{
					Data: []df.UserSaysData{
						df.UserSaysData{
							Text: "test says 1",
						},
					},
				},
			},
			Responses: []df.IntentResponse{
				df.IntentResponse{
					Action:        "test-action1",
					ResetContexts: false,
					AffectedContexts: []df.IntentAffectedContext{
						df.IntentAffectedContext{
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
			if response, err := client.UpdateIntent(response.Id, df.IntentObject{
				Name:     "test-intent-updated",
				Auto:     true,
				Contexts: []string{},
				Templates: []string{
					"test 1",
					"test 2",
				},
				UserSays: []df.UserSays{
					df.UserSays{
						Data: []df.UserSaysData{
							df.UserSaysData{
								Text: "test says 1",
							},
						},
					},
				},
				Responses: []df.IntentResponse{
					df.IntentResponse{
						Action:        "test-action1",
						ResetContexts: false,
						AffectedContexts: []df.IntentAffectedContext{
							df.IntentAffectedContext{
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
	if response, err := client.CreateContexts(sessionId, []df.ContextObject{
		df.ContextObject{
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

## Todos

- [ ] Add tests

## License

MIT

