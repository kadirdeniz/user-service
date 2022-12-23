package contract

import (
	"fmt"
	"github.com/pact-foundation/pact-go/dsl"
	"github.com/pact-foundation/pact-go/types"
	"github.com/pact-foundation/pact-go/utils"
	"testing"
	"user-service/tools/fiber"
)

type Settings struct {
	Host            string
	ProviderName    string
	BrokerBaseURL   string
	BrokerUsername  string // Basic authentication
	BrokerPassword  string // Basic authentication
	ConsumerName    string
	ConsumerVersion string // a git sha, semantic version number
	ConsumerTag     string // dev, staging, prod
	ProviderVersion string
}

func (s *Settings) create() {
	s.Host = "127.0.0.1"
	s.ProviderName = "UserService"
	s.ConsumerName = "Client"
	//s.BrokerBaseURL = "http://localhost"
	s.ConsumerTag = "main"
	s.ProviderVersion = "1.0.0"
	s.ConsumerVersion = "1.0.0"
}

func Test_CreateNewUser(t *testing.T) {

	port, _ := utils.GetFreePort()
	go fiber.StartServer(port)

	var settings Settings
	settings.create()

	pact := dsl.Pact{
		Host:                     settings.Host,
		Provider:                 settings.ProviderName,
		Consumer:                 settings.ConsumerName,
		DisableToolValidityCheck: true,
	}

	verifyRequest := types.VerifyRequest{
		ProviderBaseURL:            fmt.Sprintf("http://%s:%d", settings.Host, port),
		ProviderVersion:            settings.ProviderVersion,
		Tags:                       []string{settings.ConsumerTag},
		PactURLs:                   []string{"test/contract/pacts/Client-UserService.json"},
		PublishVerificationResults: false,
		FailIfNoPactsFound:         true,
		StateHandlers: map[string]types.StateHandler{
			"create user": func() error {
				return nil
			},
			"create user email already exists": func() error {
				return nil
			},
			"create user nickname already exists": func() error {
				return nil
			},
			"delete user": func() error {
				return nil
			},
			"get user": func() error {
				return nil
			},
			"get user not found": func() error {
				return nil
			},
			"get users": func() error {
				return nil
			},
			"update user": func() error {
				return nil
			},
		},
	}

	_, err := pact.VerifyProvider(t, verifyRequest)
	if err != nil {
		fmt.Println("Error on VerifyProvider: ", err)
		t.Fatal(err)
	}

	fmt.Println("Pact Verification succeeded!")
}
