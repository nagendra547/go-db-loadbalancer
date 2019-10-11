package health

import (
	"context"
	"database/sql"

	"github.com/nagendra547/go-db-loadbalancer/log"
	"github.com/nagendra547/go-db-loadbalancer/mydb"
)

// PingMaster -
func PingMaster(db *mydb.DB) error {
	log.Info("Checking Master")
	if err := db.Master.Ping(); err != nil {
		log.Error("Master is down")
		return err
	}
	return nil
}

// PingReadreplicas -
func PingReadreplicas(db interface{}) error {
	log.Info("Checking Read Replicas")
	if err := db.(*sql.DB).Ping(); err != nil {
		log.Error(db, "is down")
		return (err)

	}
	return nil
}

//PingContext -  health check for context
func PingContext(ctx context.Context, db *mydb.DB) error {
	log.Info("Checking Master Context")
	if err := db.Master.PingContext(ctx); err != nil {
		log.Error("Master DB is down")
		return err
	}

	for i := range db.Readreplicas {
		log.Info("Checking ReadReplicas Context")
		if err := db.Readreplicas[i].(*sql.DB).PingContext(ctx); err != nil {
			log.Error(db.Readreplicas[i], "is down")
			return err
		}
	}

	return nil
}
