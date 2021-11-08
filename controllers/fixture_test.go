package controllers

import (
	"better-console-backend/domain/member"
	"better-console-backend/domain/organization"
	"better-console-backend/domain/rbac"
	"better-console-backend/domain/site"
	"better-console-backend/domain/webhook"
	"fmt"
	"github.com/go-testfixtures/testfixtures/v3"
	_ "github.com/mattn/go-sqlite3"
)

type DatabaseFixture struct {
}

func (DatabaseFixture) setUpDefault() {
	fmt.Println("Set up database test fixture")
	gormDB.AutoMigrate(&member.MemberEntity{}, &site.SettingEntity{}, &rbac.PermissionEntity{}, &rbac.RoleEntity{},
		&organization.OrganizationEntity{}, &webhook.WebHookEntity{}, &webhook.WebHookMessageEntity{})

	sqlDB, err := gormDB.DB()
	if err != nil {
		panic(err)
	}

	fixtures, err := testfixtures.New(
		testfixtures.Database(sqlDB),                      // You database connection
		testfixtures.Dialect("sqlite"),                    // Available: "postgresql", "timescaledb", "mysql", "mariadb", "sqlite" and "sqlserver"
		testfixtures.Directory("../testdata/db_fixtures"), // the directory containing the YAML files
		testfixtures.DangerousSkipTestDatabaseCheck(),
	)

	if err != nil {
		panic(err)
	}

	if err := fixtures.Load(); err != nil {
		panic(err)
	}
	fmt.Println("End of database test fixture")
}
