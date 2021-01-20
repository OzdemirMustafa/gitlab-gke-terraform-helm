//+build provider

package main

import (
	"fmt"
	"os"
	"strings"
	"testing"

	"backend-app/config"
	"backend-app/handler"
	"backend-app/model"
	"backend-app/server"
	"backend-app/service"

	mocks "backend-app/.mocks"
	"github.com/golang/mock/gomock"
	"github.com/joho/godotenv"
	"github.com/pact-foundation/pact-go/dsl"
	"github.com/pact-foundation/pact-go/types"

	"github.com/stretchr/testify/assert"
)

func TestProvider(t *testing.T) {
	if os.Getenv("CI") == "" {
		godotenv.Load("../../.env")
	}

	mockController := gomock.NewController(t)
	mockTodoRepository := mocks.NewMockTodoRepository(mockController)
	defer mockController.Finish()

	TodoService := service.TodoService{TodoRepository: mockTodoRepository}
	Todo := handler.TodoHandler{TodoService: TodoService}

	conf, _ := config.New("../../.config/", os.Getenv("APP_ENV"))
	conf.Print()

	server, err := server.New(&conf)
	assert.Nil(t, err)

	server.E.POST("/addTodoListItem/:title", Todo.AddTodoItem)
	server.E.POST("/updateItemCompleted/:id/:completed", Todo.UpdateTodoItemCompleted)
	server.E.POST("/deleteTodoItem/:id", Todo.DeleteTodoItem)
	server.E.GET("/getTodoList", Todo.GetTodoList)

	go server.Start()

	stateHandlers := types.StateHandlers{
		"i have not empty product list which matches given barcode": func() error {
			mockTodoRepository.EXPECT().
				AddTodoListItemByTitle(gomock.Any()).
				Return(model.TodoItem{ID: 1, Title: "FirstTodo", Completed: false}, nil).
				AnyTimes()

			return nil
		},
		"i have a todo list": func() error {
			mockTodoRepository.EXPECT().
				GetTodoListItems().
				Return([]model.TodoList{{ID: 1, Title: "First Todo Item", Completed: false}, {ID: 1, Title: "Second Todo Item", Completed: false}}, nil).
				AnyTimes()
			return nil
		},
		"i have will update todo item": func() error {
			mockTodoRepository.EXPECT().
				UpdateTodoListItem(gomock.Any(), gomock.Any()).
				Return(model.TodoItem{ID: 1, Title: "FirstTodo", Completed: true}, nil).
				AnyTimes()
			return nil
		},
		"i have will delete todo item": func() error {
			mockTodoRepository.EXPECT().
				DeleteTodoItemByID(gomock.Any()).
				Return(model.TodoItem{ID: 1, Title: "", Completed: true}, nil).
				AnyTimes()
			return nil
		},
	}

	verifyProvider(t, conf.Server.Port, stateHandlers)
}

func verifyProvider(t *testing.T, providerPort int, stateHandlers types.StateHandlers) {
	godotenv.Load("./.env")

	brokerBaseURL := os.Getenv("PACT_BROKER_BASE_URL")
	brokerToken := os.Getenv("PACT_BROKER_TOKEN")
	consumerName := os.Getenv("PACT_CONSUMER_NAME")
	consumerVersion := os.Getenv("PACT_CONSUMER_VERSION")
	providerName := os.Getenv("PACT_PROVIDER_NAME")
	providerVersion := os.Getenv("PACT_PROVIDER_VERSION")

	fmt.Println(strings.Repeat("-", 20))
	fmt.Println("consumer-name: ", consumerName)
	fmt.Println("provider-name: ", providerName)
	fmt.Println("provider-version: ", providerVersion)
	fmt.Println("base-url: ", brokerBaseURL)
	fmt.Println(strings.Repeat("-", 20))

	pact := &dsl.Pact{
		Host:                     "localhost",
		DisableToolValidityCheck: true,
		Consumer:                 consumerName,
		Provider:                 providerName,
		LogDir:                   "../../logs",
		PactDir:                  "../../pacts",
	}

	pactURL := fmt.Sprintf("%s/pacts/provider/%s/consumer/%s/version/%s.json", brokerBaseURL, providerName, consumerName, consumerVersion)

	verifyRequest := types.VerifyRequest{
		ProviderBaseURL:            fmt.Sprintf("http://localhost:%d", providerPort),
		ProviderVersion:            providerVersion,
		BrokerToken:                brokerToken,
		BrokerURL:                  brokerBaseURL,
		StateHandlers:              stateHandlers,
		Provider:                   providerName,
		PublishVerificationResults: true,
		PactURLs:                   []string{pactURL},
	}

	_, err := pact.VerifyProvider(t, verifyRequest)
	if err != nil {
		fmt.Println(err.Error())
	}
}
