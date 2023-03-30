// Copyright 2019 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     https://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Sample database-sql demonstrates connecting to a Cloud SQL instance.
// The application is a Go version of the "Tabs vs Spaces"
// web app presented at Google Cloud Next 2019 as seen in this video:
// https://www.youtube.com/watch?v=qVgzP3PsXFw&t=1833s
package cloudsql

import (
	"database/sql"
	"io/ioutil"
	"log"
	"os"
	"strings"
	"sync"
	"time"
)

var (
	//indexTmpl = template.Must(template.New("index").Parse(indexHTML))
	db   *sql.DB
	once sync.Once
)

// GetDB lazily instantiates a database connection pool. Users of Cloud Run or
// Cloud Functions may wish to skip this lazy instantiation and connect as soon
// as the function is loaded. This is primarily to help testing.
func GetDB() *sql.DB {
	once.Do(func() {
		db = mustConnect()
	})
	return db
}

func readFile(path string) (*string, error) {
	content, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}
	file := string(content)
	return &file, nil
}

// migrateDB creates the table if it does not already exist.
func migrateDB(db *sql.DB) error {
	sqls, err := readFile("./dao/cloudsql/create_table.sql")
	if err != nil {
		return err
	}
	sqlArr := strings.Split(*sqls, ";")
	for _, sqlStr := range sqlArr {
		sqlStr = strings.TrimSpace(sqlStr)
		if sqlStr == "" {
			continue
		}
		_, err = db.Exec(sqlStr)
		if err != nil {
			return err
		}
	}
	return nil
}

// mustConnect creates a connection to the database based on environment
// variables. Setting one of INSTANCE_HOST, INSTANCE_UNIX_SOCKET, or
// INSTANCE_CONNECTION_NAME will establish a connection using a TCP socket, a
// Unix socket, or a connector respectively.
func mustConnect() *sql.DB {
	var (
		db  *sql.DB
		err error
	)

	// Use a TCP socket when INSTANCE_HOST (e.g., 127.0.0.1) is defined
	if os.Getenv("INSTANCE_HOST") != "" {
		db, err = connectTCPSocket()
		if err != nil {
			log.Fatalf("connectTCPSocket: unable to connect: %s", err)
		}
	}
	// Use a Unix socket when INSTANCE_UNIX_SOCKET (e.g., /cloudsql/proj:region:instance) is defined.
	if os.Getenv("INSTANCE_UNIX_SOCKET") != "" {
		db, err = connectUnixSocket()
		if err != nil {
			log.Fatalf("connectUnixSocket: unable to connect: %s", err)
		}
	}

	// Use the connector when INSTANCE_CONNECTION_NAME (proj:region:instance) is defined.
	if os.Getenv("INSTANCE_CONNECTION_NAME") != "" {
		if os.Getenv("DB_USER") == "" && os.Getenv("DB_IAM_USER") == "" {
			log.Fatal("Warning: One of DB_USER or DB_IAM_USER must be defined")
		}
		// Use IAM Authentication (recommended) if DB_IAM_USER is set
		if os.Getenv("DB_IAM_USER") != "" {
			db, err = connectWithConnectorIAMAuthN()
		} else {
			db, err = connectWithConnector()
		}
		if err != nil {
			log.Fatalf("connectConnector: unable to connect: %s", err)
		}
	}

	if db == nil {
		log.Fatal("Missing database connection type. Please define one of INSTANCE_HOST, INSTANCE_UNIX_SOCKET, or INSTANCE_CONNECTION_NAME")
	}

	//if err := migrateDB(db); err != nil {
	//	log.Fatalf("unable to create table: %s", err)
	//}

	return db
}

// configureConnectionPool sets database connection pool properties.
// For more information, see https://golang.org/pkg/database/sql
func configureConnectionPool(db *sql.DB) {
	// [START cloud_sql_mysql_databasesql_limit]
	// Set maximum number of connections in idle connection pool.
	db.SetMaxIdleConns(5)

	// Set maximum number of open connections to the database.
	db.SetMaxOpenConns(7)
	// [END cloud_sql_mysql_databasesql_limit]

	// [START cloud_sql_mysql_databasesql_lifetime]
	// Set Maximum time (in seconds) that a connection can remain open.
	db.SetConnMaxLifetime(1800 * time.Second)
	// [END cloud_sql_mysql_databasesql_lifetime]

	// [START cloud_sql_mysql_databasesql_backoff]
	// database/sql does not support specifying backoff
	// [END cloud_sql_mysql_databasesql_backoff]
	// [START cloud_sql_mysql_databasesql_timeout]
	// The database/sql package currently doesn't offer any functionality to
	// configure connection timeout.
	// [END cloud_sql_mysql_databasesql_timeout]
}
