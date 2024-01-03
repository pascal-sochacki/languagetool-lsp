package main

import (
	"context"
	"os"
	"os/signal"

	"github.com/pascal-sochacki/languagetool-lsp/internal/server"
	"github.com/pascal-sochacki/languagetool-lsp/pkg/languagetool"
	"go.lsp.dev/jsonrpc2"
	"go.lsp.dev/protocol"
	"go.uber.org/zap"
)

func NewLogger() (*zap.Logger, error) {
	cfg := zap.NewProductionConfig()
	cfg.OutputPaths = []string{
		"./logs.log",
	}
	cfg.Level = zap.NewAtomicLevel()
	cfg.Level.SetLevel(zap.DebugLevel)
	return cfg.Build()
}

type stdReaderWriterCloser struct {
	log *zap.Logger
}

func (s stdReaderWriterCloser) Read(p []byte) (int, error) {
	n, err := os.Stdin.Read(p)
	return n, err
}

func (s stdReaderWriterCloser) Write(p []byte) (int, error) {
	n, err := os.Stdout.Write(p)
	return n, err
}

func (s stdReaderWriterCloser) Close() error {
	if err := os.Stdin.Close(); err != nil {
		return err
	}
	if err := os.Stdout.Close(); err != nil {
		return err
	}
	return nil
}

func main() {
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

	stream := jsonrpc2.NewStream(stdReaderWriterCloser{log: log})
	server, serverInit := server.NewServer(log, *languagetool.NewClient())

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
}
