package config

var (
	TokenPreURL = "http://10.121.8.95:5000/v2.0/"
	Token       = ""
	SDSPreURL   = "http://10.121.8.95:9999/v1/"
)

func SetToken(token string) {
	Token = token
}

func GetToken() string {
	return Token
}
