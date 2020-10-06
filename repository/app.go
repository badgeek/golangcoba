package appRepository

import (
	"context"
	"log"

	"github.com/volatiletech/null/v8"
	"github.com/volatiletech/sqlboiler/v4/boil"
	appstate "manticore.id/golangcoba/appstate"
	"manticore.id/golangcoba/models"
)

var ctx context.Context = context.Background()

//GetVersions Get from db
func GetVersions() (models.AppVersionSlice, error) {
	versions, dbErr := models.AppVersions().All(ctx, appstate.GetConnection())
	if dbErr != nil {
		log.Fatal(dbErr.Error())
		return nil, dbErr
	}
	return versions, dbErr
}

//IsVersionExist Get from db
func IsVersionExist(id string) (bool, error) {
	verID := null.StringFrom(id)
	queryModel := models.AppVersionWhere.Name.EQ(verID)
	result, dbErr := models.AppVersions(queryModel).Count(ctx, appstate.GetConnection())
	if dbErr != nil {
		log.Fatal(dbErr.Error())
		return false, dbErr
	}
	return result > 0, nil
}

//CreateVersion Get from db
func CreateVersion(newVersion string) error {
	var newEntry models.AppVersion
	newEntry.Name = null.StringFrom(newVersion)
	err := newEntry.Insert(ctx, appstate.GetConnection(), boil.Infer())
	// result, dbErr := models.FindAppVersion()
	return err
}
