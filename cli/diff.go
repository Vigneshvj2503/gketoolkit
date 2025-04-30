package cli

import (
	"bytes"
	"context"
	"io/ioutil"
	"log"
	"os"
	"strings"

	"github.com/hashicorp/terraform-exec/tfexec"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var diffCmd = &cobra.Command{
	Use:   "diff",
	Short: "check if cluster will be recreated with the configuration",
	Run:   terraformPlan,
}

func init() {
	rootCmd.AddCommand(diffCmd)
}

func terraformPlan(*cobra.Command, []string) {
	//viper.SetConfigFile(configFile)
	configByte, ferr := ioutil.ReadFile(configFile)
	if ferr != nil {
		log.Fatalf("error reading config file: %s", ferr)
	}

	viper.SetConfigType("yaml")

	verr := viper.ReadConfig(bytes.NewReader(configByte))
	if verr != nil {
		log.Fatalf("error parsing config file: %s", verr)
	}
	workspace := viper.GetString("clustername") + "-cluster"

	execPath := "/usr/local/bin/terraform"

	tf, err := tfexec.NewTerraform(clusterfolder, execPath)

	if err != nil {
		log.Fatalf("error running Terraform: %s", err)
	}

	tf.SetStdout(os.Stdout)
	tf.SetStderr(os.Stderr)

	log.Println("*********** Running Terraform Plan on cluster module  ***********")
	err = tf.Init(context.Background(), tfexec.Upgrade(true))
	if err != nil {
		log.Fatalf("error running Init: %s", err)
	}
	err = tf.WorkspaceNew(context.Background(), workspace)

	if err != nil {
		if strings.Contains(err.Error(), "already exists") {
			tf.WorkspaceSelect(context.Background(), workspace)
		} else {
			log.Fatalf("error creating workspace: %s", err)
		}
	}

	err = tf.Init(context.Background(), tfexec.Upgrade(true))
	if err != nil {
		log.Fatalf("error running Init: %s", err)
	}
	err = tf.WorkspaceNew(context.Background(), workspace)
	if err != nil {
		if strings.Contains(err.Error(), "already exists") {
			tf.WorkspaceSelect(context.Background(), workspace)
		} else {
			log.Fatalf("error creating workspace: %s", err)
		}
	}

	c := tfexec.VarFile("../terraform.tfvars")
	_, err = tf.Plan(context.Background(), c)

	if err != nil {
		log.Fatalf("error running Terraform plan: %s", err)
	}
	log.Println()
}
