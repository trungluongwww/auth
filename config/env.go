package config

import (
	"github.com/kelseyhightower/envconfig"
)

type (
	Env struct {
		// App
		AppEnv                      string `envconfig:"APP_ENV"`
		AppName                     string `envconfig:"APP_NAME"`
		APIKey                      string `envconfig:"API_KEY"`
		UserAppDomain               string `envconfig:"USER_APP_DOMAIN"`
		AdminAppDomain              string `envconfig:"ADMIN_APP_DOMAIN"`
		SupportEmailAddress         string `envconfig:"SUPPORT_EMAIL_ADDRESS"`
		InfoEmailAddress            string `envconfig:"INFO_EMAIL_ADDRESS"`
		ContactMailReceive          string `envconfig:"CONTACT_EMAIL_RECEIVE"`
		CustomerSupportEmailAddress string `envconfig:"CUSTOMER_SUPPORT_EMAIL_ADDRESS"` // customer-support@kaitainaviibaraki.com
		InvoiceEmailAddress         string `envconfig:"INVOICE_EMAIL_ADDRESS"`          // invoice@kaitainaviibaraki.com
		PrivacyEmailAddress         string `envconfig:"PRIVACY_EMAIL_ADDRESS"`          // privacy@kaitainaviibaraki.com

		// Mysql
		MysqlUser     string `envconfig:"MYSQL_USER"`
		MysqlPassword string `envconfig:"MYSQL_PASSWORD"`
		MysqlDatabase string `envconfig:"MYSQL_DATABASE"`
		MysqlProtocol string `envconfig:"MYSQL_PROTOCOL"`

		// Third party
		SendgridAPIKey string `envconfig:"SENDGRID_API_KEY"`

		// S3
		S3Region            string `envconfig:"S3_REGION"`
		S3BucketName        string `envconfig:"S3_BUCKET_NAME"`
		S3AccessKeyId       string `envconfig:"S3_ACCESS_KEY_ID"`
		S3SecretAccessKeyId string `envconfig:"S3_SECRET_ACCESS_KEY_ID"`
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
