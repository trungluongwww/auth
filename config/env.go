package config

import (
	"github.com/kelseyhightower/envconfig"
)

type (
	Env struct {
		// App
		Port    string `envconfig:"PORT" default:"5000"`
		AppEnv  string `envconfig:"APP_ENV"`
		AppName string `envconfig:"APP_NAME"`
		APIKey  string `envconfig:"API_KEY"`

		// token
		SecretUserJWTToken  string `envconfig:"PRIVATE_USER_JWT_TOKEN"`
		SecretAdminJWTToken string `envconfig:"SECRET_ADMIN_JWT_TOKEN"`

		// Mysql
		MysqlUser     string `envconfig:"MYSQL_USER"`
		MysqlPassword string `envconfig:"MYSQL_PASSWORD"`
		MysqlDatabase string `envconfig:"MYSQL_DATABASE"`
		MysqlProtocol string `envconfig:"MYSQL_PROTOCOL"`
	}
	KeyPair struct {
		PrivateKey []byte
		PublicKey  []byte
	}
)

const (
	LocalEnv   = "local"
	StagingEnv = "staging"
	ProdEnv    = "prod"
)

func NewEnv() (*Env, error) {
	var env Env
	err := envconfig.Process("", &env)
	if err != nil {
		return nil, err
	}
	return &env, nil
}

func (e *Env) IsLocal() bool {
	return e.AppEnv == LocalEnv
}

func (e *Env) IsStaging() bool {
	return e.AppEnv == StagingEnv
}

func (e *Env) IsProd() bool {
	return e.AppEnv == ProdEnv
}
