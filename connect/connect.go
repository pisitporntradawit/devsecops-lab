package connect

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/joho/godotenv"
)

func Connect() (*pgxpool.Pool, error) {
	err := godotenv.Load(".env")
	if err != nil {
		log.Println("No .env file found")
	}
	dsn := os.Getenv("DATABASE_POSTGRES")

	conn, err := pgxpool.New(context.Background(), dsn)
	if err != nil {
		log.Fatal("Unable to connect:", err)
	}

	err = conn.Ping(context.Background())
	if err != nil {
		log.Fatal("Database not responding:", err)
	}

	fmt.Println("✅ Connected to PostgreSQL!")

	return conn, nil
}
