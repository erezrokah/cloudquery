package main

import (
	"github.com/cloudquery/cloudquery/plugins/source/firestore/resources/plugin"
	"github.com/cloudquery/plugin-sdk/v4/serve"
)

const sentryDSN = "https://0f23371080a341f1984c9f6d7e928e4a@o1396617.ingest.sentry.io/4505007730917376"

func main() {
	serve.Plugin(plugin.Plugin(), serve.WithPluginSentryDSN(sentryDSN))
}
