package database

import (
	"github.com/blyndusk/elastic-books/api/models"
	"github.com/sirupsen/logrus"
)

func Migrate() {
	_ = Db.AutoMigrate(&models.User{})
	logrus.Info("Migrations done !")
}
