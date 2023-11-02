package cmd

import (
	"errors"
	"fmt"
	"os"

	"github.com/nicelizhi/easy-admin-core/sdk/pkg"

	"github.com/nicelizhi/easy-admin/cmd/app"
	"github.com/nicelizhi/easy-admin/common/global"

	"github.com/spf13/cobra"

	"github.com/nicelizhi/easy-admin/cmd/migrate"
	"github.com/nicelizhi/easy-admin/cmd/server"
	"github.com/nicelizhi/easy-admin/cmd/version"
)

var rootCmd = &cobra.Command{
	Use:          "easy-admin",
	Short:        "easy-admin",
	SilenceUsage: true,
	Long:         `easy-admin`,
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) < 1 {
			tip()
			return errors.New(pkg.Red("requires at least one arg"))
		}
		return nil
	},
	PersistentPreRunE: func(*cobra.Command, []string) error { return nil },
	Run: func(cmd *cobra.Command, args []string) {
		tip()
	},
}

func tip() {
	usageStr := `欢迎使用 ` + pkg.Green(`easy-admin `+global.Version) + ` 可以使用 ` + pkg.Red(`-h`) + ` 查看命令`
	usageStr1 := `也可以参考 ` + global.HomePage + ` /guide/command 的相关内容`
	fmt.Printf("%s\n", usageStr)
	fmt.Printf("%s\n", usageStr1)
}

func init() {
	rootCmd.AddCommand(server.StartCmd)
	rootCmd.AddCommand(migrate.StartCmd)
	rootCmd.AddCommand(version.StartCmd)
	rootCmd.AddCommand(app.StartCmd)
}

// Execute : apply commands
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(-1)
	}
}
