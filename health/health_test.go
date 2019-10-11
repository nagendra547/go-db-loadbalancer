package health

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

//TestPingMaster - testing ping for master
func TestPingMaster(t *testing.T) {

	mydatabase := mydb.NewDB(master, readReplica1, readReplica2, readReplica3)
	myerr := PingMaster(mydatabase)
	assert.Nil(t, myerr)
}

//TestPingMaster - testing ping for Replica
func TestPingReadReplica(t *testing.T) {
	mydatabase := mydb.NewDB(master, readReplica1, readReplica2, readReplica3)

	for i := range mydatabase.Readreplicas {
		temp := mydatabase.Readreplicas[i]
		err := PingReadreplicas(temp)
		assert.Nil(t, err)
	}

}
