package creatuser

package autocert

import (
	"net/http"

	"github.com/spf13/cobra"

	"crypto/tls"
	"fmt"

	"golang.org/x/crypto/acme/autocert"
)

var (
	configYml string
	apiCheck  bool
	StartCmd  = &cobra.Command{
		Use:          "creatuser",
		Short:        "create user",
		Example:      "easy-admin createuser -c config/settings.yml",
		SilenceUsage: true,
		PreRun: func(cmd *cobra.Command, args []string) {
			setup()
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			return run()
		},
	}
)

func init() {

}

func setup() {

}

func run() error {
	return nil
}
