package main

import (
  "os"
  "fmt"
  "time"
  "strings"
  "encoding/base64"

  "github.com/aws/aws-sdk-go/aws"
  "github.com/aws/aws-sdk-go/aws/credentials"
  "github.com/aws/aws-sdk-go/aws/session"
  "github.com/aws/aws-sdk-go/service/ecr"
)

type Auth struct {
	Username string
	Password string
	ProxyEndpoint string
	ExpiresAt time.Time
}

func getEnv(key, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	} else {
		return value
	}
}

func main() {
  key := getEnv("AWS_ACCESS_KEY_ID", "")
  secret := getEnv("AWS_SECRET_ACCESS_KEY", "")
  region := getEnv("AWS_DEFAULT_REGION", "")
  registryID := getEnv("AWS_ECS_REGISTRY_ID", "")
  max := 1

  if key == "" {
    panic("wrong key")
    return
  }

  if secret == "" {
    panic("wrong secret")
    return
  }

  if region == "" {
    panic("wrong region")
    return
  }

  if registryID == "" {
    panic("wrong registry id")
    return
  }

  sess, err := session.NewSession(&aws.Config{
    Region: aws.String(region),
    Credentials: credentials.NewStaticCredentials(key, secret, ""),
  })

	svc := ecr.New(sess, aws.NewConfig().WithMaxRetries(max).WithRegion(region))

  params := &ecr.GetAuthorizationTokenInput{
    RegistryIds: []*string{aws.String(registryID)},
  }

  res, err := svc.GetAuthorizationToken(params)
  if err != nil {
    panic(err.Error())
    return
  }

  data := res.AuthorizationData
  auths := []*Auth{}

	for _, v := range data {
		token, err := base64.StdEncoding.DecodeString(*v.AuthorizationToken)
    if err != nil {
      panic(err.Error())
      return
    }

		splited := strings.SplitN(string(token), ":", 2)
    if len(splited) != 2 {
      panic("wrong token")
      return
    }

    auth := &Auth{
      Username: splited[0],
      Password: splited[1],
			ProxyEndpoint: *(v.ProxyEndpoint),
			ExpiresAt: *(v.ExpiresAt),
    }

    auths = append(auths, auth)
	}

  for _, auth := range auths {
    cmd := fmt.Sprintf("docker login -u %s -p %s %s", auth.Username, auth.Password, auth.ProxyEndpoint)
    fmt.Println(cmd)
  }
}
