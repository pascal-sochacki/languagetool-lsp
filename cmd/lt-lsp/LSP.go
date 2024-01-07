package ltlsp

import (
	"context"
	"os"
	"os/signal"

	"github.com/pascal-sochacki/languagetool-lsp/internal"
	"github.com/pascal-sochacki/languagetool-lsp/internal/server"
	"github.com/pascal-sochacki/languagetool-lsp/pkg/languagetool"
	"github.com/spf13/cobra"
	"go.lsp.dev/jsonrpc2"
	"go.lsp.dev/protocol"
	"go.uber.org/zap"
)

func init() {
	RootCmd.AddCommand(LSPRun)
}

var LSPRun = &cobra.Command{
	Use:   "run",
	Short: "start the lsp",
	Run: func(cmd *cobra.Command, args []string) {
		log, err := NewLogger()
		if err != nil {
			os.Exit(1)
		}

		defer func() {
			_ = log.Sync()
		}()
		defer func() {
			if r := recover(); r != nil {
				log.Fatal("handled panic", zap.Any("recovered", r))
			}
		}()

		ctx := context.Background()
		ctx, cancel := context.WithCancel(ctx)
		ctx = context.WithValue(ctx, "logger", log)
		signalChan := make(chan os.Signal, 1)
		signal.Notify(signalChan, os.Interrupt)
		defer func() {
			signal.Stop(signalChan)
			cancel()
		}()

		go func() {
			select {
			case <-signalChan:
				cancel()
			case <-ctx.Done():
			}
			<-signalChan // Second signal, hard exit.
			os.Exit(2)
		}()
		log.Info("starting...")

		stream := jsonrpc2.NewStream(internal.StdReaderWriterCloser{Log: log})
		server, serverInit := server.NewServer(log, *languagetool.NewClient(log))

		_, conn, client := protocol.NewServer(ctx, server, stream, log)
		defer conn.Close()

		serverInit(client)

		select {
		case <-ctx.Done():
			log.Info("context closed")
		case <-conn.Done():
			log.Info("conn closed")
		}
		return

	},
}

func NewLogger() (*zap.Logger, error) {
	cfg := zap.NewProductionConfig()
	cfg.OutputPaths = []string{
		"./logs.log",
	}
	cfg.Level = zap.NewAtomicLevel()
	cfg.Level.SetLevel(zap.DebugLevel)
	return cfg.Build()
}
