package utilstests

import (
	"database/sql"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/bperezgo/admin_franchise/config"
	"github.com/ory/dockertest/v3"
	"github.com/ory/dockertest/v3/docker"
)

type Docker struct {
	Pool       *dockertest.Pool
	Host, Port string
}

func UpDocker() Docker {
	pool, err := dockertest.NewPool("")
	if err != nil {
		log.Fatalf("Could not construct pool: %s", err)
	}

	err = pool.Client.Ping()
	if err != nil {
		log.Fatalf("Could not connect to Docker: %s", err)
	}

	c := config.GetConfig()

	// pulls an image, creates a container based on it and runs it
	resource, err := pool.RunWithOptions(&dockertest.RunOptions{
		Repository: "postgres",
		Tag:        "latest",
		Env: []string{
			fmt.Sprintf("POSTGRES_PASSWORD=%s", c.POSTGRES_PASSWORD),
			fmt.Sprintf("POSTGRES_USER=%s", c.POSTGRES_USERNAME),
			fmt.Sprintf("POSTGRES_DB=%s", c.POSTGRES_DATABASE),
			"listen_addresses = '*'",
		},
	}, func(config *docker.HostConfig) {
		// set AutoRemove to true so that stopped container goes away by itself
		config.AutoRemove = true
		config.RestartPolicy = docker.RestartPolicy{Name: "no"}
	})
	if err != nil {
		log.Fatalf("Could not start resource: %s", err)
	}

	hostAndPort := resource.GetHostPort("5432/tcp")
	databaseUrl := fmt.Sprintf("postgres://%s:%s@%s/%s?sslmode=disable",
		c.POSTGRES_USERNAME,
		c.POSTGRES_PASSWORD,
		hostAndPort,
		c.POSTGRES_DATABASE,
	)

	log.Println("Connecting to database on url: ", databaseUrl)

	if err := resource.Expire(120); err != nil { // Tell docker to hard kill the container in 120 seconds
		log.Fatalf("Could not set expiration: %s", err)
	}

	// exponential backoff-retry, because the application in the container might not be ready to accept connections yet
	pool.MaxWait = 120 * time.Second
	if err = pool.Retry(func() error {
		db, err := sql.Open("postgres", databaseUrl)
		if err != nil {
			return err
		}
		return db.Ping()
	}); err != nil {
		log.Fatalf("Could not connect to docker: %s", err)
	}

	hostAndPortSeparated := strings.Split(hostAndPort, ":")

	host := hostAndPortSeparated[0]
	port := hostAndPortSeparated[1]

	c.ChangePostgresPort(port)

	return Docker{
		Pool: pool,
		Host: host,
		Port: port,
	}
}
