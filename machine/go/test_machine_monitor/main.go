// A command-line application where each sub-command implements a get_* call in bot_config.py
package main

import (
	"context"
	"encoding/json"
	"flag"
	"io/fs"
	"strings"
	"time"

	"go.skia.org/infra/go/common"
	"go.skia.org/infra/go/sklog"
	"go.skia.org/infra/machine/go/configs"
	"go.skia.org/infra/machine/go/machineserver/config"
	"go.skia.org/infra/machine/go/test_machine_monitor/machine"
	"go.skia.org/infra/machine/go/test_machine_monitor/server"
	"go.skia.org/infra/machine/go/test_machine_monitor/swarming"
)

const (
	// Make the triggerInterrogation channel buffered so we don't lag responding
	// to HTTP requests from the Swarming bot.
	interrogationChannelSize = 10

	oauth2RetryDuration = time.Minute
)

// flags
var (
	configFlag        = flag.String("config", "prod.json", "The name to the configuration file, such as prod.json or test.json, as found in machine/go/configs.")
	local             = flag.Bool("local", false, "Running locally if true. As opposed to in production.")
	machineServerHost = flag.String("machine_server", "https://machines.skia.org", "A URL with the scheme and domain name of the machine hosting the machine server API.")
	metadataURL       = flag.String("metadata_url", "http://metadata:8000/computeMetadata/v1/instance/service-accounts/default/token", "The URL of the metadata server that provides service account tokens.")
	port              = flag.String("port", ":11000", "HTTP service address (e.g., 'localhost:8000' or ':8000')")
	promPort          = flag.String("prom_port", ":20000", "Metrics service address (e.g., 'localhost:10110' or ':10110')")
	pythonExe         = flag.String("python_exe", "", "Absolute path to Python.")
	startSwarming     = flag.Bool("start_swarming", false, "Start swarming_bot.zip.")
	swarmingBotZip    = flag.String("swarming_bot_zip", "", "Absolute path to where the swarming_bot.zip code should run from.")
	username          = flag.String("username", "chrome-bot", "The username of the account that accepts SSH connections.")
)

var (
	// Version can be changed via -ldflags.
	Version = "development"
)

func main() {
	var err error

	for {
		err = common.InitWith(
			"test_machine_monitor",
			common.PrometheusOpt(promPort),
			common.CloudLogging(local, "skia-public"),
		)
		if err == nil {
			break
		}
		// Keep re-trying if the error comes from the Date/Time on the machine being wrong.
		//
		// There is a longer response from the server returned in the error
		// message here of the form:
		//
		//     Response: {"error":"invalid_grant","error_description":"Invalid JWT:
		//     Token must be a short-lived token (60 minutes) and in a reasonable
		//     timeframe. Check your iat and exp values in the JWT claim."}
		//
		// But that response is controlled by the server and may change in the
		// future.
		if strings.Contains(err.Error(), "Failed to initialize: oauth2: cannot fetch token: 400 Bad Request") {
			sklog.Info("Date/Time on machine is incorrect resulting in an oauth2 error.")
			time.Sleep(oauth2RetryDuration)
		} else {
			// Some other error.
			sklog.Fatalf("Failed to initialize: %s", err)
		}
	}

	sklog.Infof("Version: %s", Version)
	var instanceConfig config.InstanceConfig
	b, err := fs.ReadFile(configs.Configs, *configFlag)
	if err != nil {
		sklog.Fatalf("Failed to read config file %q: %s", *configFlag, err)
	}
	err = json.Unmarshal(b, &instanceConfig)
	if err != nil {
		sklog.Fatal(err)
	}
	if err != nil {
		sklog.Fatalf("Failed to open config file: %q: %s", *configFlag, err)
	}

	ctx := context.Background()
	triggerInterrogationCh := make(chan bool, interrogationChannelSize)
	machineState, err := machine.New(ctx, *local, instanceConfig, Version, *startSwarming, *machineServerHost, triggerInterrogationCh)
	if err != nil {
		sklog.Fatalf("Failed to create machine: %s", err)
	}
	if err := machineState.Start(ctx); err != nil {
		sklog.Fatalf("Failed to start machine: %s", err)
	}

	sklog.Infof("Starting the server.")
	machineSwarmingServer, err := server.New(machineState, triggerInterrogationCh)
	if err != nil {
		sklog.Fatal(err)
	}
	go func() {
		sklog.Fatal(machineSwarmingServer.Start(*port))
	}()

	if *startSwarming {
		if *pythonExe == "" {
			sklog.Fatalf("Flag --python_exe is required when --start_swarming is true.")
		}
		if *swarmingBotZip == "" {
			sklog.Fatalf("Flag --swarming_bot_zip is required when --start_swarming is true.")
		}
		sklog.Infof("Starting swarming_bot.")
		bot, err := swarming.New(*pythonExe, *swarmingBotZip, *metadataURL)
		if err != nil {
			sklog.Fatal(err)
		}
		sklog.Fatal(bot.Launch(ctx))
	} else {
		select {}
	}
}
