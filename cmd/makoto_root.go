package cmd

import (
	"log"
	"os"
	"path"

	homedir "github.com/atrox/homedir"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	Makoto = &cobra.Command{
		Use:   "Makoto",
		Short: "Makoto is a karel package configurator",
		Long: `
	   	  MAKOTO A KAREL PACKAGE CONFIGURATOR
	   	  ===================================

Makoto is a Package Configurator for Karel files.
The idea to this tool was given using package config in linux.
The package configure it is easy to build your code because
the tool handels the setting for the compiler flags.

Beacause FANUC KAREL has a limitation on the number of lines in
the code I do have the programing stype to seperate all KAREL
information in dedecated files. To get the information together
this tool helps me.

AUTHOR:
	Anton Feldmann <anton.feldmann@gmail.com>

\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\`,
	}

	confFile, filepath string
)

func init() {
	cobra.OnInitialize(initConfig)

	Makoto.PersistentFlags().StringVar(&confFile, "config",
		"",
		"config file (default $HOME/.config/makoto/makoto.json)")

	Makoto.AddCommand(version)
	Makoto.AddCommand(kpc_cmd)
	Makoto.AddCommand(all)
}

func Execute() {
	Makoto.Execute()
}

func initConfig() {
	if confFile != "" {
		viper.SetConfigFile(confFile)
	} else {

		home, err := homedir.Dir()
		if err != nil {
			log.Fatal(err)
		}

		filepath = path.Join(home, ".config", "makoto")
		confFile = "makoto.json"

		viper.AddConfigPath(filepath)
		viper.SetConfigName("makoto")
	}

	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		log.Println(err)
	} else {
		if err := viper.Unmarshal(&rfg); err != nil {
			log.Fatal("unable to decode into the makoto configuration structure, %v", err)
		}
	}

	if len(rfg.RootDir) > 0 {

		rfg.init(getEnvOrDefault("KPC_PATH", filepath))
	}
	rfg.save(filepath, confFile)

}

func getEnvOrDefault(name, or string) string {
	if value, ok := os.LookupEnv(name); ok {
		return value
	}
	return or
}
