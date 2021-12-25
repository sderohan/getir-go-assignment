package db

import (
	"fmt"
)

type ConnectionURL struct {
	UserName string
	Password string
	Host     string
	Database string
	Port     uint
	Options  map[string]string
}

func (c ConnectionURL) String() string {
	prefix := "mongodb+srv"
	options := ""

	for key, value := range c.Options {
		options += fmt.Sprintf("%s=%s&", key, value)
	}

	connectionString := fmt.Sprintf("%s://%s:%s@%s/%s?%s", prefix, c.UserName, c.Password, c.Host, c.Database, options)
	return connectionString
}

// mongodb+srv://challengeUser:WUMglwNBaydH8Yvu@challenge-xzwqd.mongodb.net/getir-case-study?retryWrites=true
