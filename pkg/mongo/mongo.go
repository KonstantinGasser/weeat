package mongo

import (
	"context"
	"fmt"
	"time"

	"github.com/pkg/errors"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

// Conn represents a connection to
// a mongoDB instance
type Conn struct {
	user, password string
	host           string
	port           int
	timeout        time.Duration
	c              *mongo.Client
}

func New(host string, port int, timeout time.Duration, opts ...func(*Conn)) *Conn {
	conn := &Conn{
		host:    host,
		port:    port,
		timeout: timeout,
	}

	for _, opt := range opts {
		opt(conn)
	}

	return conn
}

func WithAuth(username, password string) func(*Conn) {
	return func(c *Conn) {
		c.user, c.password = username, password
	}
}

// Connect opens a mongoDB connection with the Conn host and port
// properties. It times-out after the set Conn.timeout
func (conn *Conn) Connect() error {

	ctx, cancel := context.WithTimeout(context.Background(), conn.timeout)
	defer cancel()

	var err error
	conn.c, err = mongo.Connect(ctx, options.Client().ApplyURI(conn.uri()))
	if err != nil {
		return errors.Wrap(err, "connecting to mongodb caused an issue")
	}

	if err := conn.Ping(); err != nil {
		return err
	}

	return nil
}

func (conn *Conn) Ping() error {
	ctx, cancel := context.WithTimeout(context.Background(), conn.timeout)
	defer cancel()

	return errors.Wrap(conn.c.Ping(ctx, readpref.Primary()), "ping to mongodb failed")
}

func (conn *Conn) uri() string {
	if len(conn.user) > 0 && len(conn.password) > 0 {
		return fmt.Sprintf("mongodb://%s:%s@%s:%d", conn.user, conn.password, conn.host, conn.port)
	}
	return fmt.Sprintf("mongodb://%s:%d", conn.host, conn.port)
}
