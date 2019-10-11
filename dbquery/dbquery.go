/*
Package dbquery - A package created to implement db query, transaction and exec related activites
*/
package dbquery

import (
	"context"
	"database/sql"

	"github.com/nagendra547/go-db-loadbalancer/dbadmin"
	"github.com/nagendra547/go-db-loadbalancer/log"
	"github.com/nagendra547/go-db-loadbalancer/mydb"
)

/*
Query - Implementing error and health check for this API.
Similar code can be also used for other APIs too. Implementing this as prototype, same strategy can be used for others.
Multi-Threading also implemented using go subroutine
*/
func Query(db *mydb.DB, query string, args ...interface{}) (*sql.Rows, error) {
	var rows *sql.Rows
	var err error
	operationDone := make(chan bool)
	go func() {
		log.Info("Query readReplica DB")
		readReplica := dbadmin.ReadReplicaRoundRobin(db)
		log.Info("ReadReplica DB is ", readReplica)
		rows, err = readReplica.Query(query, args...)
		if err != nil {
			log.Error("Error while executing query")
		}

		operationDone <- true
	}()
	<-operationDone

	return rows, err
}

//QueryContext - User can implement this API same as Query API
func QueryContext(db *mydb.DB, ctx context.Context, query string, args ...interface{}) (*sql.Rows, error) {
	return dbadmin.ReadReplicaRoundRobin(db).QueryContext(ctx, query, args...)
}

func QueryRow(db *mydb.DB, query string, args ...interface{}) *sql.Row {
	return dbadmin.ReadReplicaRoundRobin(db).QueryRow(query, args...)
}

func QueryRowContext(db *mydb.DB, ctx context.Context, query string, args ...interface{}) *sql.Row {
	return dbadmin.ReadReplicaRoundRobin(db).QueryRowContext(ctx, query, args...)
}

func Begin(db *mydb.DB) (*sql.Tx, error) {
	return db.Master.Begin()
}

func BeginTx(db *mydb.DB, ctx context.Context, opts *sql.TxOptions) (*sql.Tx, error) {
	return db.Master.BeginTx(ctx, opts)
}

func Exec(db *mydb.DB, query string, args ...interface{}) (sql.Result, error) {
	return db.Master.Exec(query, args...)
}

func ExecContext(db *mydb.DB, ctx context.Context, query string, args ...interface{}) (sql.Result, error) {
	return db.Master.ExecContext(ctx, query, args...)
}

func Prepare(db *mydb.DB, query string) (*sql.Stmt, error) {
	return db.Master.Prepare(query)
}

func PrepareContext(db *mydb.DB, ctx context.Context, query string) (*sql.Stmt, error) {
	return db.Master.PrepareContext(ctx, query)
}
