package dbquery

import (
	"database/sql"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/nagendra547/go-db-loadbalancer/log"
	"github.com/nagendra547/go-db-loadbalancer/mydb"
	"github.com/stretchr/testify/assert"
)

var master *sql.DB
var readReplica1 *sql.DB
var readReplica2 *sql.DB
var readReplica3 *sql.DB

var mockMaster sqlmock.Sqlmock
var mock0 sqlmock.Sqlmock
var mock1 sqlmock.Sqlmock
var mock2 sqlmock.Sqlmock

/*This init is executed all time while creating a new mock instance of database
 */
func init() {
	var err0 error
	master, mockMaster, err0 = sqlmock.New()
	checkError(err0)

	readReplica1, mock0, err0 = sqlmock.New()
	checkError(err0)
	readReplica2, mock1, err0 = sqlmock.New()
	checkError(err0)
	readReplica3, mock2, err0 = sqlmock.New()
	checkError(err0)

}
func checkError(err error) {
	if err != nil {
		log.Error("an error '%s' was not expected when opening a stub database connection", err)
	}
}

//TestQuery - testing Query
// all the other testing can be also done in similar fashion
func TestQuery(t *testing.T) {

	mydatabase := mydb.NewDB(master, readReplica1, readReplica2, readReplica3)

	//mock1.ExpectBegin()
	id := int64(1)
	name := "some_name"
	timeNow := time.Now().UTC()
	newProp := "disallow"

	rows1 := sqlmock.NewRows([]string{"id", "name", "property", "created_at", "updated_at"}).
		AddRow(id, name, newProp, timeNow, timeNow)

	mock1.ExpectQuery("Select products").WithArgs("3").WillReturnRows(rows1)
	rows, err := Query(mydatabase, "Select products", "3")
	testCols, _ := rows.Columns()
	assert.Nil(t, err)
	assert.Equal(t, len(testCols), 5, "failed")
	assert.Equal(t, testCols[0], "id", "failed")

}
