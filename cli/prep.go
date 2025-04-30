package cli

import (
	"bytes"
	"io/ioutil"
	"log"
	"os"

	//"github.com/hashicorp/terraform-config-inspect/tfconfig"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var prepCmd = &cobra.Command{
	Use:   "prep",
	Short: "prepare the cluster configuration",
	Run:   Prep,
}

func init() {
	rootCmd.AddCommand(prepCmd)
	prepCmd.Flags().BoolP("includeasm", "a", false, "Include ASM")
	prepCmd.Flags().BoolP("addcsi", "s", false, "Add Secrets Store Csi Driver")
}

func Prep(cmd *cobra.Command, args []string) {
	configByte, ferr := ioutil.ReadFile(configFile)
	if ferr != nil {
		log.Fatalf("error reading config file: %s", ferr)
	}

	viper.SetConfigType("yaml")
	viper.SetConfigFile(configFile)
	verr := viper.ReadConfig(bytes.NewReader(configByte))
	if verr != nil {
		log.Fatalf("error parsing config file: %s", verr)
	}

	createUserInputTemplate()

	asm, _ := cmd.Flags().GetBool("includeasm")
	csi, _ := cmd.Flags().GetBool("addcsi")
	//install := viper.GetStringSlice("installsequence")
	install := getInstallSequence()
	del := getDeleteSequence()
	cleanUp := getCleanupSequence()

	mods := getModules()

	if !asm {
		install = replaceList(install, "asm")
		del = replaceList(del, "asm")
		delete(mods, "asm")
		mods["newoverlay"] = overlaysReplace
	}

	if !csi {
		install = replaceList(install, "secretcsi")
		del = replaceList(del, "secretcsi")
		delete(mods, "secretcsi")

	}
	viper.Set("installsequence", install)
	viper.Set("deletesequence", del)
	viper.Set("cleanupsequence", cleanUp)
	viper.Set("modules", mods)
	viper.WriteConfig()

}

func replaceList(source []string, origin string) []string {

	for i, item := range source {
		if item == origin {
			return append(source[:i], source[i+1:]...)
		}
	}
	return source

}

func createUserInputTemplate() {
	f, err := os.Create("user.input.template")
	if err != nil {
		panic(err)
	}
	defer f.Close()
	_, err = f.WriteString(inputtemplate)
	if err != nil {
		panic(err)
	}
}
