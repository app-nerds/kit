package database

import (
	"github.com/globalsign/mgo"
	"github.com/pkg/errors"
)

/*
MongoDatabaseService encapsulates behaviors for working with MongoDB databases
*/
type MongoDatabaseService struct {
	DB      *mgo.Database
	Session *mgo.Session
}

func (d *MongoDatabaseService) Connect(connection *Connection) error {
	var err error

	if d.Session, err = mgo.Dial(connection.Host); err != nil {
		return errors.Wrapf(err, "Error connecting to database server %s", connection.Host)
	}

	d.DB = d.Session.DB(connection.DatabaseName)
	return nil
}

func (d *MongoDatabaseService) Disconnect() {
	d.Session.Close()
}

func (d *MongoDatabaseService) GetCollection(name string) *mgo.Collection {
	return d.DB.C(name)
}

func (d *MongoDatabaseService) GetDB() *mgo.Database {
	return d.DB
}
