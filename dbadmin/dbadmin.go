/*
Package dbadmin - A package created to implement db admin related activites
*/
package dbadmin

import (
	"database/sql"
	"time"

	"github.com/nagendra547/go-db-loadbalancer/health"
	"github.com/nagendra547/go-db-loadbalancer/log"
	"github.com/nagendra547/go-db-loadbalancer/mydb"
)

/*ReadReplicaRoundRobin - Get a read replica in round robin fashion. Implemented multi threading using go subroutine
Other methods can be also implemented in similar fashion
*/
func ReadReplicaRoundRobin(db *mydb.DB) *sql.DB {
	log.Info("Checking ReadReplicaRoundRobin")

	// Check how many replicas are actually available.
	// If Ping not working then no need to count in available replicas

	var availableReplicas []interface{}
	var index int
	operationDone := make(chan bool)
	go func() {
		db.Count++
		for i := range db.Readreplicas {
			temp := db.Readreplicas[i]

			if err := health.PingReadreplicas(temp); err == nil {
				availableReplicas = append(availableReplicas, temp)
			}
		}
		log.Info("Available Replicas", len(availableReplicas))
		index = db.Count % len(availableReplicas)
		operationDone <- true
	}()
	<-operationDone
	return availableReplicas[index].(*sql.DB)
}

// SetConnMaxLifetime - Setting Connection Max Life time
func SetConnMaxLifetime(db *mydb.DB, d time.Duration) {
	log.Info("Setting SetConnMaxLifetime")
	db.Master.SetConnMaxLifetime(d)
	for i := range db.Readreplicas {
		db.Readreplicas[i].(*sql.DB).SetConnMaxLifetime(d)
	}
}

// SetMaxIdleConns - Setting Max Idle connections
func SetMaxIdleConns(db *mydb.DB, n int) {
	log.Info("Setting SetMaxIdleConns")
	if err := health.PingMaster(db); err == nil {
		db.Master.SetMaxIdleConns(n)
	}

	for i := range db.Readreplicas {
		r1 := db.Readreplicas[i]
		if err := health.PingReadreplicas(r1); err == nil {
			r1.(*sql.DB).SetMaxIdleConns(n)
		} else {
			log.Error(r1, "is down.", "SetMaxIdleConns has been ignored")
		}

	}
}

// SetMaxOpenConns - Setting Max Open connections
func SetMaxOpenConns(db *mydb.DB, n int) {
	log.Info("Setting SetMaxOpenConns")
	db.Master.SetMaxOpenConns(n)
	for i := range db.Readreplicas {
		db.Readreplicas[i].(*sql.DB).SetMaxOpenConns(n)
	}
}

// Close -
func Close(db *mydb.DB) error {
	log.Info("Closing Master DB")
	error := db.Master.Close()
	if error != nil {
		log.Error(db.Master, "is down")
		return error
	}
	for i := range db.Readreplicas {
		log.Info("Closing Read Replicas")
		error := db.Readreplicas[i].(*sql.DB).Close()
		log.Error(db.Readreplicas[i], "is down")
		return error
	}
	return nil
}
