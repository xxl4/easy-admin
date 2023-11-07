package creatuser

import (
	"fmt"

	"github.com/nicelizhi/easy-admin-core/config/source/file"
	"github.com/nicelizhi/easy-admin-core/sdk"
	"github.com/nicelizhi/easy-admin-core/sdk/config"
	"github.com/nicelizhi/easy-admin-core/sdk/service"
	"github.com/nicelizhi/easy-admin/app/admin/models"
	"github.com/nicelizhi/easy-admin/common/database"
	"github.com/spf13/cobra"
)

type SysUser struct {
	service.Service
}

var (
	configYml string
	host      string
	StartCmd  = &cobra.Command{
		Use:     "creatuser",
		Short:   "create user",
		Example: "easy-admin createuser -c config/settings.yml",
		RunE: func(cmd *cobra.Command, args []string) error {
			return run()
		},
	}
)

func init() {
	StartCmd.PersistentFlags().StringVarP(&configYml, "config", "c", "config/settings.yml", "Start server with provided configuration file")
}

func run() error {
	config.Setup(
		file.NewSource(file.WithPath(configYml)),
		initDB,
	)
	if host == "" {
		host = "*"
	}
	db := sdk.Runtime.GetDbByKey(host)

	var model *models.SysUser

	var username = "admin"
	//var defaultpass = "admin"

	row := db.Debug().Where("username = ?", username).First(&model)

	if row.Error != nil {
		fmt.Println(row.Error)
		return nil
	}

	fmt.Println(row.RowsAffected)

	return nil
}

func initDB() {
	database.Setup()
}
