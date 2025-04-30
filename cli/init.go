package cli

import (
	"fmt"
	"log"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var initCmd = &cobra.Command{
	Use:   "init",
	Short: "initiate the cluster configuration",
	Run:   iniCreateConfig,
}

func init() {
	rootCmd.AddCommand(initCmd)
	//var cluster string
	//initCmd.Flags().StringVarP(&cluster, "cluster", "-c", "", "specify the cluster full name")
	initCmd.Flags().BoolP("force", "f", false, "Force to re-initate")
	initCmd.Flags().StringP("cluster", "c", "", "Cluster full name")
	initCmd.Flags().StringP("bucket", "b", "terraform-state-otl-us", "The bucket for Terraform backend")
}

func iniCreateConfig(cmd *cobra.Command, args []string) {

	forced, _ := cmd.Flags().GetBool("force")
	if _, err := os.Stat(configFile); err == nil {
		if !forced {
			log.Fatalln("config file already exists. Init will remove the previous configuration. Exit")
		}
	}
	viper.SetConfigType("yaml")
	viper.SetConfigFile(configFile)
	name, _ := cmd.Flags().GetString("cluster")
	if len(name) == 0 {
		log.Fatalln("cluster name is required. Exit")
	}
	bucket, _ := cmd.Flags().GetString("bucket")
	viper.Set("clustername", name)
	viper.Set("installsequence", installsequence)
	viper.Set("deletesequence", deletesequence)
	viper.Set("modules", modules)
	viper.Set("bucket", bucket)
	viper.WriteConfig()
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Version of GKE toolkit",
	Run:   Version,
}

func init() {
	rootCmd.AddCommand(versionCmd)
}

func Version(cmd *cobra.Command, args []string) {
	fmt.Println("otgke version " + version)
}
