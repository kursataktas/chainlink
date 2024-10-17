package main

import (
	"context"
	"crypto/tls"

	"github.com/spf13/cobra"
	"google.golang.org/grpc/credentials"

	"github.com/smartcontractkit/chainlink/integration-tests/deployment"
	"github.com/smartcontractkit/chainlink/integration-tests/deployment/devenv"
	"github.com/smartcontractkit/chainlink/v2/core/logger"
)

var (
	rootCmd       = &cobra.Command{Use: ""}
	lggr          logger.Logger
	cribConfig    string
	cribEnv       *deployment.Environment
	cribEnvConfig devenv.EnvironmentConfig
	noOfChains    int
	privateKeys   []string
)

func init() {
	var closeLggr func() error
	lggr, closeLggr = logger.NewLogger()
	defer func() {
		err := closeLggr()
		if err != nil {
			panic(err)
		}
	}()
	rootCmd.PersistentFlags().IntVarP(&noOfChains, "no-of-chains", "n", 1, "number of chains to deploy")
	err := rootCmd.MarkPersistentFlagRequired("no-of-chains")
	if err != nil {
		lggr.Fatalw("no number of chains provided", "err", err)
	}
	for i := 0; i < noOfChains; i++ {
		var privateKey string
		rootCmd.PersistentFlags().StringVarP(&privateKey, "private-key", "p", "", "private key for chain")
		err = rootCmd.MarkPersistentFlagRequired("private-key")
		if err != nil {
			lggr.Fatalw("no private key provided for ", "chain index", i, "err", err)
		}
		privateKeys = append(privateKeys, privateKey)
	}
	rootCmd.PersistentFlags().StringVarP(&cribConfig, "crib-config", "c", "", "CRIB environment configuration file")
	err = rootCmd.MarkPersistentFlagRequired("crib-config")
	if err != nil {
		lggr.Fatalw("no chain configuration file provided", "err", err)
	}
	cribEnvConfig, err = devenv.LoadEnvironmentConfig(cribConfig)
	if err != nil {
		lggr.Fatalw("failed to load environment configuration", "err", err)
	}
	if !cribEnvConfig.JDConfig.IsEmpty() {
		cribEnvConfig.JDConfig.Creds = credentials.NewTLS(&tls.Config{
			MinVersion: tls.VersionTLS12,
		})
	}
	cribEnv, _, err = devenv.NewEnvironment(context.Background(), lggr, cribEnvConfig)
	if err != nil {
		lggr.Fatalw("failed to create environment", "err", err)
	}

	rootCmd.AddCommand(ccipHomeDeploy)
}

func main() {
	err := rootCmd.Execute()
	if err != nil {
		lggr.Fatalw("Error executing command", "err", err)
	}
}
