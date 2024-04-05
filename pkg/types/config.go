package types

type Config struct {
	WhatsApp WhatsAppConfig `yaml:"whatsapp"`
	Logger   LoggerConfig   `yaml:"logger"`
}

// whatsapp client configuration
type WhatsAppConfig struct {
	Version           string `yaml:"version"`
	PhoneNumberID     string `yaml:"phone_number_id"`
	BusinessAccountID string `yaml:"business_account_id"`
	AccessToken       string `yaml:"access_token"`
}

// logger configuration
type LoggerConfig struct {
	Mode             string `yaml:"mode"`
	Encoding         string `yaml:"encoding"`
	Level            string `yaml:"level"`
	EnableStacktrace bool   `yaml:"enable_stacktrace"`
}
