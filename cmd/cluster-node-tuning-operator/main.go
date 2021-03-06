package main

import (
	"flag"
	"os"
	"runtime"

	"k8s.io/klog"

	tunedv1 "github.com/openshift/cluster-node-tuning-operator/pkg/apis/tuned/v1"
	"github.com/openshift/cluster-node-tuning-operator/pkg/operator"
	"github.com/openshift/cluster-node-tuning-operator/pkg/signals"
	"github.com/openshift/cluster-node-tuning-operator/version"
)

var (
	boolVersion = flag.Bool("version", false, "show program version and exit")
	boolLocal   = flag.Bool("local", false, "local run outside a pod")
)

func printVersion() {
	klog.Infof("Go Version: %s", runtime.Version())
	klog.Infof("Go OS/Arch: %s/%s", runtime.GOOS, runtime.GOARCH)
	klog.Infof("%s Version: %s", tunedv1.TunedClusterOperatorResourceName, version.Version)
}

func main() {
	klog.InitFlags(nil)
	flag.Parse()

	printVersion()

	if *boolVersion {
		os.Exit(0)
	}

	stopCh := signals.SetupSignalHandler()

	controller, err := operator.NewController()
	if err != nil {
		klog.Fatal(err)
	}

	err = controller.Run(stopCh)
	if err != nil {
		klog.Fatalf("error running controller: %s", err.Error())
	}
}
