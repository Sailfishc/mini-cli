package templates

import (
	"fmt"

	"github.com/biezhi/moe"
	"github.com/mkideal/cli"
	"github.com/mkideal/pkg/debug"
)

// BaseConfig base config
type BaseConfig struct {
	Name            string `cli:"-"`
	PackageName     string
	Version         string
	RenderType      string
	DBType          string
	BuildTool       string
	BladeVersion    string
	TplDependency   string
	MySQLDependency string
}

type maker func(*cli.Context, *BaseConfig) error

var templatesMap = make(map[string]maker)

// register tempalteMap
func register(tool string, fn maker) bool {
	if _, ok := templatesMap[tool]; ok {
		debug.Panicf("repeat register template %s", tool)
	}
	templatesMap[tool] = fn
	return true
}

// New new application
// 工厂类型约定使用New
func New(ctx *cli.Context, cfg BaseConfig) error {
	clr := ctx.Color()
	moe := moe.New(clr.Bold("creating project, please wait...")).Spinner("dots3").Color(moe.Green).Start()

	fn, ok := templatesMap[cfg.BuildTool]
	if !ok {
		return fmt.Errorf("unsupported template type %s", clr.Yellow(cfg.BuildTool))
	}
	err := fn(ctx, &cfg)
	moe.Stop()
	if err == nil {
		fmt.Printf("application %s create successful!\n", cfg.Name)
	}
	return err
}
