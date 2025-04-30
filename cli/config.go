package cli

const version = "v23.4.01-1"

const configFile = ".gke.yml"

var installsequence = []string{"apis", "cluster", "ipmasq", "acm", "asm", "overlays", "newoverlay", "upload"}
var deletesequence = []string{"asm", "acm", "apis", "ipmasq", "upload", "cluster", "newoverlay"}
var cleanupsequence = []string{"overlays"}

//	"proxy":    {"paas-bootstrap"} is moved out
var modules = map[string][]string{
	"cluster":    {"ot-cluster-gke"},
	"ipmasq":     {"paas-ipmasq"},
	"asm":        {"ot-managed-asm"},
	"secretcsi":  {"ot-secretstore"},
	"newoverlay": {"ot-kustomize", "ot-asm-acm", "ot-gitlab-agents", "ot-trident", "ot-prometheus", "ot-backupyaml", "ot-twistlock", "ot-newrelic", "ot-asm-policy"},
	"overlays":   {"ot-velero"},
	"upload":     {"gitlab-utils"},
	"acm":        {"ot-acm"},
	"apis":       {"google-apis"},
}

var overlaysReplace = []string{"ot-kustomize", "ot-nginx", "ot-trident", "ot-backupyaml", "ot-prometheus", "ot-twistlock", "ot-newrelic"}

const path = "../../../modules/"
const overlayfolder = "newoverlay"
const clusterfolder = "cluster"

func getInstallSequence() []string {
	return installsequence
}

func getDeleteSequence() []string {
	return deletesequence
}

func getCleanupSequence() []string {
	return cleanupsequence
}

func getModules() map[string][]string {
	return modules
}
