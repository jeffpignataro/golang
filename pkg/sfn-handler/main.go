package sfnhandler

import (
	"context"

	log "github.com/BillHeroInc/logrus"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/sfn"
)

type Config struct {
	client *sfn.Client
}

func NewConfig() (*Config, error) {
	cfg, err := config.LoadDefaultConfig(context.TODO(), config.WithRegion("us-west-2"))
	if err != nil {
		log.Error(err.Error())
		return nil, err
	}

	return &Config{
		client: sfn.NewFromConfig(cfg, sfn.WithAPIOptions()),
	}, nil
}

func (c Config) GetActivities() (*sfn.ListActivitiesOutput, error) {
	a, err := c.client.ListActivities(context.TODO(), &sfn.ListActivitiesInput{}, sfn.WithAPIOptions())
	if err != nil {
		log.Error(err.Error())
	}

	return a, nil
}
