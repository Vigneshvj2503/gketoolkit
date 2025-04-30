package cli

import (
	"bufio"
	"bytes"
	_ "embed"
	"io/ioutil"
	"log"
	"os"
	"text/template"

	"github.com/shihanng/tfvar/pkg/tfvar"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

type ModuleFile struct {
	Name     string
	Source   string
	Varstmts []string
}

//go:embed resources/module.tmpl
var r string

//go:embed resources/backend.tmpl
var backend string

//go:embed resources/input.tfvars
var inputtemplate string

var tfvarCmd = &cobra.Command{
	Use:   "tfvar",
	Short: "Generate tfvars based on user input",
	Run:   Tfvar,
}

func init() {
	rootCmd.AddCommand(tfvarCmd)
	//tfvarCmd.Flags().BoolP("includeasm", "a", false, "Include ASM")
	tfvarCmd.Flags().StringArrayP("files", "f", []string{}, `Put deployment tfvars(i.e lab or commercial, lab is default) before user generated tfvars`)
}

func Tfvar(cmd *cobra.Command, args []string) {
	fromFiles, err := cmd.Flags().GetStringArray("files")
	if err != nil {
		log.Fatalf("Error reading cli flags %s", err)
	}
	unparseds := make(map[string]tfvar.UnparsedVariableValue)

	for _, fv := range fromFiles {
		if err := tfvar.CollectFromFile(fv, unparseds); err != nil {
			log.Fatalf("Error processing user input tfvars %s", err)
		}
	}

	vars := defaultTfvar()
	vars, err = tfvar.ParseValues(unparseds, vars)
	if err != nil {
		log.Fatalf("Error processing user input tfvars %s", err)
	}

	f, err := os.Create("terraform.tfvars")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	w := bufio.NewWriter(f)
	tfvar.WriteAsTFVars(w, vars)
	w.Flush()
}

func defaultTfvar() []tfvar.Variable {
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

	mods := viper.GetStringMapStringSlice("modules")

	uniquevars := make(map[string]tfvar.Variable)
	for step, folders := range mods {
		var modvarnames []string
		var variables []tfvar.Variable
		os.Mkdir(step, 0700)
		for _, folder := range folders {
			names, vals := scan(path, folder)
			modvarnames = append(modvarnames, names...)
			variables = append(variables, vals...)
			createModuleFile(step, folder, names)
		}
		modvarnames = removeDuplicateStr(modvarnames)
		createVarFile(step, modvarnames)
		createBackendFile(step)

		for _, v := range variables {
			uniquevars[v.Name] = v
		}
	}
	addOverlayFolder()

	var basicTfvars []tfvar.Variable
	for _, value := range uniquevars {
		basicTfvars = append(basicTfvars, value)
	}

	return basicTfvars
}

func addOverlayFolder() {
	if _, err := os.Stat(overlayfolder); os.IsNotExist(err) {
		log.Println("overlays folder does not exist, return")
	} else {
		os.Mkdir(overlayfolder+"/files", 0700)
	}
}

func removeDuplicateStr(strSlice []string) []string {
	allKeys := make(map[string]bool)
	list := []string{}
	for _, item := range strSlice {
		if _, value := allKeys[item]; !value {
			allKeys[item] = true
			list = append(list, item)
		}
	}
	return list
}

func createModuleFile(step string, folder string, varnames []string) {
	f, err := os.Create(step + "/" + folder + ".tf")
	if err != nil {
		panic(err)
	}
	defer f.Close()
	temp := template.Must(template.New("t1").Parse(r))

	var stmts []string
	for _, singleVar := range varnames {
		stmts = append(stmts, singleVar+" = var."+singleVar)
	}

	module := ModuleFile{Name: "\"" + folder + "\"", Source: "\"../../../../modules/" + folder + "\"", Varstmts: stmts}
	w := bufio.NewWriter(f)
	temp.Execute(w, module)
	w.Flush()

}

func createVarFile(step string, modvarnames []string) {
	f, err := os.Create(step + "/variables.tf")
	if err != nil {
		panic(err)
	}
	defer f.Close()
	for _, name := range modvarnames {
		_, err := f.WriteString("variable " + name + " {}\n")
		if err != nil {
			panic(err)
		}
	}
}

func createBackendFile(step string) {
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

	bucket_record := viper.GetString("bucket")

	type Backend struct {
		Bucket string
	}

	bucket := Backend{
		Bucket: bucket_record,
	}

	f, err := os.Create(step + "/backend.tf")
	if err != nil {
		panic(err)
	}
	defer f.Close()
	temp := template.Must(template.New("t1").Parse(backend))
	temp.Execute(f, bucket)
	//_, err = f.WriteString(backend)
	if err != nil {
		panic(err)
	}

}

func scan(path string, folder string) ([]string, []tfvar.Variable) {
	var varNames []string
	vars, err := tfvar.Load(path + folder)

	if err != nil {
		log.Fatalf("error reading config file: %s", err)
	}
	//      fmt.Println()
	//      fmt.Println("// *********** from module " + folder)
	//      tfvar.WriteAsTFVars(os.Stdout, vars)

	for _, variable := range vars {
		varNames = append(varNames, variable.Name)
	}
	return varNames, vars
}
