package cmd

import (
	"os"
	"path"
	"strings"

	"github.com/ArkjuniorK/gofiber-boilerplate/internal/core"
	"github.com/dave/jennifer/jen"
	"github.com/spf13/cobra"
)

var genCmd = &cobra.Command{
	Use:   "gen",
	Short: "Generator for model and pkg/module",
}

var modelGenCmd = &cobra.Command{
	Use:   "model",
	Short: "Generate model and placed in internal/model",
	Run:   generateModel,
}

var pkgGenCmd = &cobra.Command{
	Use:   "pkg",
	Short: "Generate new pkg or module and placed in internal/pkg",
	Run:   generatePkg,
}

func generateModel(cmd *cobra.Command, args []string) {
	logger := core.NewLogger().GetCore()

	name, err := cmd.Flags().GetString("name")
	if err != nil {
		logger.Error("Unable to get the name of the model", "error", err)
		return
	}

	if len(name) == 0 {
		logger.Warn("Name of the model is required")
		return
	}

	logger.Info("Generate model", "name", name)

	wd, err := os.Getwd()
	if err != nil {
		logger.Error("Unable to get current working directory", "error", err)
	}

	var modelName string
	names := strings.Split(name, "_")
	for _, name := range names {
		modelName += strings.ToUpper(name[0:1]) + name[1:]
	}

	modelFile := path.Join(wd, "internal", "model", name+".go")
	model := jen.NewFilePath("model")
	model.Type().Id(modelName).Struct()

	if err = model.Save(modelFile); err != nil {
		logger.Error("Unable to generate the model", "error", err)
		return
	}

	logger.Info("Success generate model", "name", name)
}

func generatePkg(cmd *cobra.Command, args []string) {
	var (
		logger   = core.NewLogger().GetCore()
		corePath = "github.com/ArkjuniorK/gofiber-boilerplate/internal/core"
	)

	name, err := cmd.Flags().GetString("name")
	if err != nil {
		logger.Error("Unable to get the name of the pkg", "error", err)
		return
	}

	if len(name) == 0 {
		logger.Error("Name of the pkg is required")
		return
	}

	logger.Info("Generate pkg/module", "name", name)

	wd, err := os.Getwd()
	if err != nil {
		logger.Error("Unable to get current working directory", "error", err)
		return
	}

	pkgDir := path.Join(wd, "internal", "pkg", name)

	if err = os.Mkdir(pkgDir, os.ModePerm); err != nil {
		logger.Error("Unable to create the pkg", "error", err)
		return
	}

	logger.Info("Generate pkg's schema")

	schFile := path.Join(pkgDir, "schema.go")
	sch := jen.NewFilePath(name)

	if err = sch.Save(schFile); err != nil {
		logger.Error("Unable to generate the pkg schema", "error", err)
		return
	}

	logger.Info("Generate pkg's repository")

	rpoFile := path.Join(pkgDir, "repository.go")
	rpo := jen.NewFilePath(name)
	rpo.Type().Id("repository").Struct(
		jen.Id("db").Op("*").Qual(corePath, "Database"),
	)

	if err = rpo.Save(rpoFile); err != nil {
		logger.Error("Unable to generate the pkg repository", "error", err)
		return
	}

	logger.Info("Generate pkg's service")

	svcFile := path.Join(pkgDir, "service.go")
	svc := jen.NewFilePath(name)
	svc.Type().Id("service").Struct(
		jen.Id("repo").Op("*").Id("repository"),
		jen.Id("logger").Op("*").Qual(corePath, "Logger"),
	)

	if err = svc.Save(svcFile); err != nil {
		logger.Error("Unable to generate the pkg service", "error", err)
		return
	}

	logger.Info("Generate pkg's handler")

	hdlFile := path.Join(pkgDir, "handler.go")
	hdl := jen.NewFilePath(name)
	hdl.Type().Id("handler").Struct(
		jen.Id("service").Op("*").Id("service"),
		jen.Id("repository").Op("*").Id("repository"),
	)

	if err = hdl.Save(hdlFile); err != nil {
		logger.Error("Unable to generate the pkg handler", "error", err)
		return
	}

	logger.Info("Generate pkg's router")

	rtrFile := path.Join(pkgDir, "router.go")
	rtr := jen.NewFilePath(name)
	rtr.Type().Id("router").Struct(
		jen.Id("api").Op("*").Qual(corePath, "API"),
		jen.Id("pubsub").Op("*").Qual(corePath, "PubSub"),
		jen.Id("logger").Op("*").Qual(corePath, "Logger"),

		jen.Id("handler").Op("*").Id("handler"),
	).Line().Func().Params(jen.Id("r").Op("*").Id("router")).Id("Serve").Params().Block()

	if err = rtr.Save(rtrFile); err != nil {
		logger.Error("Unable to generate the pkg router", "error", err)
		return
	}

	logger.Info("Generate pkg's main")

	mainFile := path.Join(pkgDir, name+".go")
	main := jen.NewFilePath(name)
	main.Func().Id("New").
		Params(
			jen.Id("db").Op("*").Qual(corePath, "Database"),
			jen.Id("api").Op("*").Qual(corePath, "API"),
			jen.Id("lg").Op("*").Qual(corePath, "Logger"),
			jen.Id("ps").Op("*").Qual(corePath, "PubSub"),
		).Block(
		jen.Var().Defs(
			jen.Id("rpo").Op("=").Op("&").Id("repository").Values(jen.Id("db")),
			jen.Id("svc").Op("=").Op("&").Id("service").Values(jen.Id("rpo"), jen.Id("lg")),
			jen.Id("hdl").Op("=").Op("&").Id("handler").Values(jen.Id("svc"), jen.Id("rpo")),
			jen.Id("rtr").Op("=").Op("&").Id("router").Values(jen.Id("api"), jen.Id("ps"), jen.Id("lg"), jen.Id("hdl")),
		).Line().Id("rtr").Dot("Serve").Call(),
	)

	if err = main.Save(mainFile); err != nil {
		logger.Error("Unable to generate the pkg main", "error", err)
		return
	}

	logger.Info("Success generate pkg/module", "name", name)
}
