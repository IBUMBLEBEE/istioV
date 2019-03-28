package bootstrap

var Args args

type args struct {
	Host           string
	Port           int
	InCluster      bool
	ConfigFile     string
	Namespace      string
	IstioNamespace string
}
