/*Package mydb - to define the structure of DB
 */
package mydb

import (
	"database/sql"
)

// DB - DB struct
type DB struct {
	Master       *sql.DB
	Readreplicas []interface{}
	Count        int
}

// NewDB - constructor
func NewDB(master *sql.DB, readreplicas ...interface{}) *DB {
	return &DB{
		Master:       master,
		Readreplicas: readreplicas,
	}
}
