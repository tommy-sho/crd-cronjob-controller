module github.com/tommy-sho

go 1.12

require (
	github.com/go-logr/logr v0.1.0
	github.com/onsi/ginkgo v1.8.0
	github.com/onsi/gomega v1.5.0
	github.com/revel/cron v0.21.0
	github.com/tommy-sho/crd-cronjob-controller v0.0.0-20190707054404-5499ade0d68c
	golang.org/x/net v0.0.0-20190613194153-d28f0bde5980
	k8s.io/api v0.0.0-20190409021203-6e4e0e4f393b
	k8s.io/apimachinery v0.0.0-20190703205208-4cfb76a8bf76
	k8s.io/client-go v11.0.1-0.20190409021438-1a26190bd76a+incompatible
	sigs.k8s.io/controller-runtime v0.2.0-beta.2
)
