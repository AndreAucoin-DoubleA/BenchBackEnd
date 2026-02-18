package db

import (
	"log"
	"time"

	gocql "github.com/apache/cassandra-gocql-driver/v2"
)

func ConnectToCassandra(host string, port int, keyspace string) *gocql.Session {
	var session *gocql.Session
	var err error

	// 1️⃣ Connect WITHOUT keyspace to create it
	for {
		cluster := gocql.NewCluster(host)
		cluster.Port = port
		cluster.Consistency = gocql.Quorum

		session, err = cluster.CreateSession()
		if err != nil {
			log.Println("Waiting for Cassandra container to be ready...")
			time.Sleep(3 * time.Second)
			continue
		}
		break
	}
	log.Println("Connected to Cassandra (no keyspace)")

	// Create keyspace
	err = session.Query(`
		CREATE KEYSPACE IF NOT EXISTS ` + keyspace + `
		WITH replication = {
			'class': 'SimpleStrategy',
			'replication_factor': 1
		};
	`).Exec()
	if err != nil {
		log.Fatal("Failed creating keyspace:", err)
	}
	log.Println("Keyspace created or already exists")
	session.Close() // close temporary session

	// 2️⃣ Connect WITH keyspace to create tables
	for {
		cluster := gocql.NewCluster(host)
		cluster.Port = port
		cluster.Keyspace = keyspace
		cluster.Consistency = gocql.Quorum

		session, err = cluster.CreateSession()
		if err != nil {
			log.Println("Waiting for keyspace to be ready...")
			time.Sleep(3 * time.Second)
			continue
		}
		break
	}
	log.Println("Connected to Cassandra with keyspace:", keyspace)

	err = session.Query(`
		CREATE TABLE IF NOT EXISTS user_by_email (
			email TEXT,
			id UUID,
			password_hash TEXT,
			created_at TEXT,
			PRIMARY KEY(email)
		);
	`).Exec()
	if err != nil {
		log.Fatal("Failed creating table:", err)
	}

	// Create tables
	err = session.Query(`
		CREATE TABLE IF NOT EXISTS wiki_url_stats (
			stat_date TEXT,
			url TEXT,
			count counter,
			PRIMARY KEY (stat_date, url)
		);
	`).Exec()
	if err != nil {
		log.Fatal("Failed creating table:", err)
	}

	err = session.Query(`
		CREATE TABLE IF NOT EXISTS wiki_users_stats (
			stat_date TEXT,
    		username TEXT,
    		PRIMARY KEY (stat_date, username)
		);
	`).Exec()
	if err != nil {
		log.Fatal("Failed creating table:", err)
	}

	err = session.Query(`
		CREATE TABLE IF NOT EXISTS wiki_total_stats (
			stat_date TEXT,
			total_changes counter,
			num_bots counter,
			num_non_bots counter,
			PRIMARY KEY (stat_date)
		);
	`).Exec()
	if err != nil {
		log.Fatal("Failed creating table:", err)
	}
	log.Println("Cassandra tables initialized")

	return session
}
