package email

/*
A Config object tells us how to configure our email server connection
*/
type Config struct {
	Host     string
	Password string
	Port     int
	UserName string
}
