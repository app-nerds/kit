package database

/*
BaseDatabase defines the simplest things a database can do: connect and disconnect
*/
type BaseDatabase interface {
	Connect(connection *Connection) error
	Disconnect()
}
