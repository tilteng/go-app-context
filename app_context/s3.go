package app_context

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

type s3JSONConfig struct {
	Environment map[string]string `json:"environment"`
}

func (self *baseAppContext) configureFromS3() error {
	s3_config := os.Getenv("APPCTX_S3_CONFIG")
	if s3_config == "" {
		return nil
	}

	keyparts := strings.SplitN(s3_config, "::", 3)
	if len(keyparts) != 3 || keyparts[0] == "" || keyparts[1] == "" || keyparts[2] == "" {
		return fmt.Errorf("Invalid APPCTX_S3_CONFIG format. Should be: region::bucket::key")
	}

	sess, err := session.NewSession()
	if err != nil {
		return fmt.Errorf("Error creating new aws session: %s", err)
	}

	svc := s3.New(sess, (&aws.Config{}).WithRegion(keyparts[0]))

	output, err := svc.GetObject(
		(&s3.GetObjectInput{}).SetBucket(keyparts[1]).SetKey(keyparts[2]),
	)

	if err != nil {
		return fmt.Errorf("Error getting s3 object from %s", s3_config, err)
	}

	defer output.Body.Close()

	decoder := json.NewDecoder(output.Body)
	conf := &s3JSONConfig{}

	err = decoder.Decode(conf)
	if err != nil {
		return fmt.Errorf(
			"Error reading/decoding json from s3 location '%s': %s",
			s3_config,
			err,
		)
	}

	if conf.Environment != nil {
		for k, v := range conf.Environment {
			if _, ok := os.LookupEnv(k); ok {
				// Don't overwrite existing environemnt
				continue
			}
			os.Setenv(k, v)
		}
	}

	return nil
}
