package main

import (
	"context"
	"runtime"

	stub "github.com/banzaicloud/prometheus-jmx-exporter-operator/pkg/stub"
	sdk "github.com/coreos/operator-sdk/pkg/sdk"
	sdkVersion "github.com/coreos/operator-sdk/version"

	"github.com/sirupsen/logrus"
)

func printVersion() {
	logrus.Infof("Go Version: %s", runtime.Version())
	logrus.Infof("Go OS/Arch: %s/%s", runtime.GOOS, runtime.GOARCH)
	logrus.Infof("operator-sdk Version: %v", sdkVersion.Version)
}

func main() {
	printVersion()
	sdk.Watch("banzaicloud.com/v1alpha1", "PrometheusJmxExporter", "default", 10)
	sdk.Watch("v1", "Pod", "default", 10)
	sdk.Handle(stub.NewHandler())
	sdk.Run(context.TODO())
}
