package model

type Config struct {
	Port          string
	Stream        string
	CassandraPort int
	CassandraHost string
	KeyspaceKey   string
	JWTSecret     string
}
