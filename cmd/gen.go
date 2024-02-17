package cmd

import (
	"os"
	"path"
	"strings"

	"github.com/ArkjuniorK/gofiber-boilerplate/internal/core"
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

	fn := path.Join(wd, "internal", "model", name+".go")
	fl, err := os.Create(fn)
	if err != nil {
		logger.Error("Unable create the model", "error", err)
	}

	modelName := strings.ToUpper(name[0:1]) + name[1:]

	payload := `package model

type ` + modelName + ` struct {}
	`

	_, err = fl.WriteString(payload)
	if err != nil {
		logger.Error("Unable to write the model", "error", err)
	}

	logger.Info("Success generate model", "name", name)
}

func generatePkg(cmd *cobra.Command, args []string) {}
