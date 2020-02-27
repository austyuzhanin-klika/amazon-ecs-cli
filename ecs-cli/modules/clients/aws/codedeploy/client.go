package codedeploy

import (
	"github.com/aws/amazon-ecs-cli/ecs-cli/modules/clients"
	"github.com/aws/amazon-ecs-cli/ecs-cli/modules/config"
	awscd "github.com/aws/aws-sdk-go/service/codedeploy"
)

// Client CodeDeploy interface
type CodeDeployClient interface {
	CreateDeployment(input *awscd.CreateDeploymentInput) (*awscd.CreateDeploymentOutput, error)
}

// codeDeployClient implements Client
type codeDeployClient struct {
	client *awscd.CodeDeploy
	config *config.CommandConfig
}

func (c *codeDeployClient) CreateDeployment(input *awscd.CreateDeploymentInput) (*awscd.CreateDeploymentOutput, error) {
	output, err := c.client.CreateDeployment(input)
	if err != nil {
		return nil, err
	}
	return output, nil
}

// NewCodeDeployClient Creates a new CodeDeployClient client
func NewCodeDeployClient(config *config.CommandConfig) CodeDeployClient {
	client := awscd.New(config.Session)
	client.Handlers.Build.PushBackNamed(clients.CustomUserAgentHandler())

	return &codeDeployClient{
		client: client,
		config: config,
	}
}
