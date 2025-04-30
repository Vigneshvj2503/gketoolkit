package cli

import (
	"bytes"
	"context"
	"io"
	"io/ioutil"
	"log"
	"os"
	"strings"
	"time"

	"github.com/hashicorp/terraform-exec/tfexec"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "delete cluster",
	Run:   deleteCluster,
}

func init() {
	rootCmd.AddCommand(deleteCmd)
}

func deleteCluster(*cobra.Command, []string) {
	configFile := ".gke.yml"

	configByte, ferr := ioutil.ReadFile(configFile)
	if ferr != nil {
		log.Fatalf("error reading config file: %s", ferr)
	}

	viper.SetConfigType("yaml")

	verr := viper.ReadConfig(bytes.NewReader(configByte))
	if verr != nil {
		log.Fatalf("error parsing config file: %s", verr)
	}
	clusterName := viper.GetString("clustername")
	directories := viper.GetStringSlice("deletesequence")

	logFile := "otgke" + time.Now().Format("20060102150405") + ".log"
	f, err := os.Create(logFile)
	if err != nil {
		log.Fatalf("error creating log file: %s", err)
	}
	_, err = f.WriteString("Otgke version " + version + " \n")
	if err != nil {
		log.Fatalf("error writing to log file: %s", err)
	}
	f.Sync()
	f.Close()

	for _, dir := range directories {
		tDestroy(dir, clusterName+"-"+dir, logFile)
	}

}

func tDestroy(workingDir string, workspace string, logFile string) {
	execPath := "/usr/local/bin/terraform"

	tf, err := tfexec.NewTerraform(workingDir, execPath)

	if err != nil {
		log.Fatalf("error running Terraform: %s", err)
	}

	f, err := os.OpenFile(logFile, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)

	if err != nil {
		log.Fatalf("error opening file: %v", err)
	}
	defer f.Close()
	log.SetOutput(f)
	wrt := io.MultiWriter(os.Stdout, f)
	tf.SetStdout(wrt)
	tf.SetStderr(wrt)
	log.Println("***********executing in directory ", pwd(), " ***********")
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
	err = tf.Destroy(context.Background(), c)

	if err != nil {
		log.Fatalf("error running Destroy: %s", err)
	}
	log.Println()

}
