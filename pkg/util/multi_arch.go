package util

import (
	"fmt"
	"runtime"

	"github.com/sirupsen/logrus"
)

// ResolveMultiArchNamespaceFor returns the namespace name based on the os architecture
func ResolveMultiArchNamespaceFor(namespace string) string {
	arch := runtime.GOARCH
	if arch == "amd64" {
		return namespace
	}
	ret := fmt.Sprintf("%s-%s", namespace, arch)
	logrus.Infof("Resolved multi-arch namespace for %s to %s", namespace, ret)
	return ret
}
