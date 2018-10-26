package database

/*
A Connection struct stores information about how to connect to a database
*/
type Connection struct {
	Host         string `json:"host"`
	DatabaseName string `json:"databaseName"`
}
