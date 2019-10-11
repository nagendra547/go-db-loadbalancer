package dbadmin

import (
	"database/sql"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/nagendra547/go-db-loadbalancer/log"
	"github.com/nagendra547/go-db-loadbalancer/mydb"
	"github.com/stretchr/testify/assert"
)

var master *sql.DB
var readReplica1 *sql.DB
var readReplica2 *sql.DB
var readReplica3 *sql.DB

/*This init is executed all time while creating a new mock instance of database
 */
func init() {
	var err0 error
	master, _, err0 = sqlmock.New()
	checkError(err0)

	readReplica1, _, err0 = sqlmock.New()
	checkError(err0)
	readReplica2, _, err0 = sqlmock.New()
	checkError(err0)
	readReplica3, _, err0 = sqlmock.New()
	checkError(err0)

}
func checkError(err error) {
	if err != nil {
		log.Error("an error '%s' was not expected when opening a stub database connection", err)
	}
}

//TestSetMaxOpenConns - testing SetMaxOpenConns
func TestSetMaxOpenConns(t *testing.T) {

	mydatabase := mydb.NewDB(master, readReplica1, readReplica2, readReplica3)
	SetMaxOpenConns(mydatabase, 10)

	stats := mydatabase.Master.Stats()
	assert.Equal(t, stats.MaxOpenConnections, 10, "Failed")
}

//TestSetMaxIdleConns - testing TestSetMaxIdleConns
func TestSetMaxIdleConns(t *testing.T) {

	mydatabase := mydb.NewDB(master, readReplica1, readReplica2, readReplica3)
	SetMaxIdleConns(mydatabase, 1)

	stats := mydatabase.Master.Stats()
	assert.Equal(t, stats.Idle, 1, "Failed")
}
