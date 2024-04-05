package types

type Config struct {
	WhatsApp WhatsAppConfig `yaml:"whatsapp"`
	Logger   LoggerConfig   `yaml:"logger"`
}

// whatsapp client configuration
type WhatsAppConfig struct {
	Version           string                 `yaml:"version"`
	BusinessAccountID string                 `yaml:"business_account_id"`
	PhoneNumberID     string                 `yaml:"phone_number_id"`
	AccessToken       string                 `yaml:"access_token"`
	TemplateIDs       map[string]string      `yaml:"template_ids"`
	Others            map[string]interface{} `yaml:",inline"`
}

// logger configuration
type LoggerConfig struct {
	Mode             string `yaml:"mode"`
	Encoding         string `yaml:"encoding"`
	Level            string `yaml:"level"`
	EnableStacktrace bool   `yaml:"enable_stacktrace"`
}
