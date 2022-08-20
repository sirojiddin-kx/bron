package testing

import (
	"fmt"
	"os"
	"testing"

	_ "github.com/jackc/pgx/v4/stdlib"
	"github.com/jmoiron/sqlx"
	"github.com/sirojiddin-kx/bron/config"
	"github.com/sirojiddin-kx/bron/storage"
)

var (
	TestStorage storage.StorageI
	companyId   = "b7ab5845-70e9-4632-84eb-80b59319cf8d"
)

func TestMain(t *testing.M) {
	var err error
	cfg := config.Load()

	conStr := fmt.Sprintf("host=%s port=%v user=%s password=%s dbname=%s sslmode=%s", cfg.PostgresHost,
		cfg.PostgresPort, cfg.PostgresUser, cfg.PostgresPassword, cfg.PostgresDatabase, "disable")

	db, err := sqlx.Connect("pgx", conStr)
	fmt.Println("connected successfully")
	if err != nil {
		panic(err)
	}

	TestStorage = storage.NewStoragePG(db)

	os.Exit(t.Run())
}
