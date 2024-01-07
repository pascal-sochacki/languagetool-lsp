package internal

import (
	"os"

	"go.uber.org/zap"
)

type StdReaderWriterCloser struct {
	Log *zap.Logger
}

func (s StdReaderWriterCloser) Read(p []byte) (int, error) {
	n, err := os.Stdin.Read(p)
	return n, err
}

func (s StdReaderWriterCloser) Write(p []byte) (int, error) {
	n, err := os.Stdout.Write(p)
	return n, err
}

func (s StdReaderWriterCloser) Close() error {
	if err := os.Stdin.Close(); err != nil {
		return err
	}
	if err := os.Stdout.Close(); err != nil {
		return err
	}
	return nil
}
