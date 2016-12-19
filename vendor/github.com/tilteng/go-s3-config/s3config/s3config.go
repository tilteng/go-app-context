package s3config

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

type EnvironmentConfig struct {
	Environment map[string]string `json:"environment"`
}

func FetchConfig(region, bucket, key string, v interface{}) error {
	sess, err := session.NewSession()
	if err != nil {
		return fmt.Errorf("Error creating new aws session: %s", err)
	}

	svc := s3.New(sess, (&aws.Config{}).WithRegion(region))

	output, err := svc.GetObject(
		(&s3.GetObjectInput{}).SetBucket(bucket).SetKey(key),
	)

	if err != nil {
		return fmt.Errorf("Error getting object from s3: %s", err)
	}

	defer output.Body.Close()

	decoder := json.NewDecoder(output.Body)

	err = decoder.Decode(v)
	if err != nil {
		return fmt.Errorf("Error reading/decoding json from s3: %s", err)
	}

	return nil
}

func SetEnvironment(region, bucket, key string, overwrite bool) error {
	conf := &EnvironmentConfig{}

	err := FetchConfig(region, bucket, key, conf)
	if err != nil {
		return err
	}

	if conf.Environment == nil {
		return errors.New("No environment found in s3 content")

	}

	for k, v := range conf.Environment {
		if !overwrite {
			if _, ok := os.LookupEnv(k); ok {
				// Don't overwrite existing environemnt
				continue
			}
		}
		os.Setenv(k, v)
	}

	return nil
}
