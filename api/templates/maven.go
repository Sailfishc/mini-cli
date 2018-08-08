package templates

import (
	"os"

	"github.com/mkideal/cli"
)

var _ = register("Maven", Maven)

// Create New Application To Maven
func Maven(ctx *cli.Context, cfg *BaseConfig) error {
	appDir := "mini-cli-application"

	if err := os.Mkdir(appDir, 0755); err != nil {
		return nil
	}

	// create pom.xml
	// pomPath := appDir + "/pom.xml"
	// utils.WriteTemplate("tpl_pom", pomPath, TplPom, cfg)
	// out, err := os.Create(pomPath)

	return nil
}
