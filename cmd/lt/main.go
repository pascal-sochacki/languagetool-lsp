package main

import (
	"context"
	"fmt"
	"os"

	"github.com/pascal-sochacki/languagetool-lsp/pkg/languagetool"
	"go.uber.org/zap"
)

func main() {
	client := languagetool.NewClient(zap.NewNop())
	result, err := client.CheckText(context.Background(), "Das ist ein Langer Text das hat fehler hat!", "auto")
	if err != nil {
		fmt.Sprintln(err.Error())
		os.Exit(2)
	}
	fmt.Printf("%+v", result)

}
