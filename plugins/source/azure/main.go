package main

import (
	"github.com/cloudquery/cloudquery/plugins/source/azure/resources/plugin"
	"github.com/cloudquery/plugin-sdk/serve"
	// _ "net/http/pprof"
)

const sentryDSN = "https://be7c45692567444299f8bef3de545b86@o1396617.ingest.sentry.io/6747596"

func main() {
	// go func() {
	// 	log.Println(http.ListenAndServe("localhost:8080", nil))
	// }()
	serve.Source(plugin.Plugin(), serve.WithSourceSentryDSN(sentryDSN))
}
