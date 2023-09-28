package repository_test

import (
	"errors"
	"fmt"
	"service-routes/internal/domain/entity"
	"service-routes/internal/domain/repository"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func Test_RepositoryResource_RegisterResource(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("failed to open sqlmock database: %s", err)
	}
	defer db.Close()

	gormDB, err := gorm.Open(mysql.New(mysql.Config{
		Conn:                      db,
		SkipInitializeWithVersion: true,
	}), &gorm.Config{})
	gormDB.Logger.LogMode(logger.Info)
	if err != nil {
		t.Fatalf("failed to open gorm database: %s", err)
	}
	createdAt := time.Now()
	activity := entity.Resource{
		Model: entity.Model{
			ID:        0x1,
			CreatedAt: createdAt,
			UpdatedAt: createdAt,
		},
	}

	repo := repository.NewRepositoryResource(gormDB)

	t.Run("Success", func(t *testing.T) {
		mock.ExpectBegin()

		mock.ExpectExec("INSERT INTO `resources` (.+) VALUES (.+)").WithArgs(
			sqlmock.AnyArg(),
			sqlmock.AnyArg(),
			sqlmock.AnyArg(),
			sqlmock.AnyArg(),
			sqlmock.AnyArg(),
			sqlmock.AnyArg(),
			sqlmock.AnyArg(),
		).WillReturnResult(sqlmock.NewResult(1, 1))
		mock.ExpectCommit()

		result, err := repo.AddResource(activity)
		if err := mock.ExpectationsWereMet(); err != nil {
			fmt.Printf("Unfulfilled expectations: %s \n", err)
		}
		assert.NoError(t, err)
		assert.Equal(t, &activity, result)
	})

	t.Run("Error", func(t *testing.T) {
		mock.ExpectBegin()
		mock.ExpectQuery("INSERT INTO `resources` (.+) VALUES (.+)").WithArgs(
			sqlmock.AnyArg(),
			sqlmock.AnyArg(),
			sqlmock.AnyArg(),
			sqlmock.AnyArg(),
			sqlmock.AnyArg(),
			sqlmock.AnyArg(),
			sqlmock.AnyArg(),
		).WillReturnError(errors.New("failed to insert user"))
		mock.ExpectRollback()

		_, err := repo.AddResource(activity)
		assert.Error(t, err)
	})
}
