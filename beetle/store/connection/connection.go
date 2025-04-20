package connection

import (
	"fmt"
	"os"

	"github.com/rs/zerolog"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"github.com/dark-vinci/stripchat/beetle/utils"
)

const packageName = "stripchat.store.connection"

type Store struct {
	Connection *gorm.DB
	Log        *zerolog.Logger
}

func NewStore(z zerolog.Logger, e *utils.Environment) *Store {
	log := z.With().Str(utils.PackageStrHelper, packageName).Logger()

	db, err := gorm.Open(
		postgres.Open(
			fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=disable TimeZone=Africa/Lagos",
				e.DbHost,
				e.DbPort,
				e.DbUsername,
				e.DbName,
				e.DbPassword,
			),
		),
		&gorm.Config{},
	)

	if err != nil {
		log.Fatal().Err(err).Msg("failed to connect to master databases")
		panic(err)
	}

	return &Store{
		Connection: db,
		Log:        &z,
	}
}

//// GetConnection helper for tests/mock
//func GetConnection(t *testing.T) (sqlmock.Sqlmock, *Store) {
//	var (
//		mock sqlmock.Sqlmock
//		db   *gorm.DB
//		err  error
//	)
//
//	db, mock, err = gorm_sqlmock.New(gorm_sqlmock.Config{
//		Config:     &gorm.Config{},
//		DriverName: "postgres",
//		DSN:        "mock",
//	})
//
//	require.NoError(t, err)
//
//	return mock, NewFromDB(db)
//}

// NewFromDB created a new storage with just the database reference passed in
func NewFromDB(db *gorm.DB) *Store {
	logger := zerolog.New(os.Stderr).With().Timestamp().Logger()
	storeLog := logger.With().Str(utils.PackageStrHelper, packageName).Logger()

	return &Store{
		Connection: db,
		Log:        &storeLog,
	}
}

func (db *Store) Close() {
	m, _ := db.Connection.DB()
	err := m.Close()

	if err != nil {
		return
	}
}
