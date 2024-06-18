package session

type Config struct {
	SigningKey    string `yaml:"signing_key"`
	EncryptingKey string `yaml:"encrypting_key"`
}
