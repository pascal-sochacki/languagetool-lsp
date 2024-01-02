package server

import (
	"context"
	"testing"

	"github.com/pascal-sochacki/languagetool-lsp/pkg/languagetool"
	"go.lsp.dev/protocol"
	"go.uber.org/zap"
)

type MockServer struct {
}

// CheckText implements languagetool.LanguagetoolApi.
func (MockServer) CheckText(ctx context.Context, text string, language string) (languagetool.CheckResult, error) {
	return languagetool.CheckResult{}, nil
}

type ClientRecorder struct {
}

// ApplyEdit implements protocol.Client.
func (ClientRecorder) ApplyEdit(ctx context.Context, params *protocol.ApplyWorkspaceEditParams) (result bool, err error) {
	panic("unimplemented")
}

// Configuration implements protocol.Client.
func (ClientRecorder) Configuration(ctx context.Context, params *protocol.ConfigurationParams) (result []interface{}, err error) {
	panic("unimplemented")
}

// LogMessage implements protocol.Client.
func (ClientRecorder) LogMessage(ctx context.Context, params *protocol.LogMessageParams) (err error) {
	panic("unimplemented")
}

// Progress implements protocol.Client.
func (ClientRecorder) Progress(ctx context.Context, params *protocol.ProgressParams) (err error) {
	panic("unimplemented")
}

// PublishDiagnostics implements protocol.Client.
func (ClientRecorder) PublishDiagnostics(ctx context.Context, params *protocol.PublishDiagnosticsParams) (err error) {
	return nil
}

// RegisterCapability implements protocol.Client.
func (ClientRecorder) RegisterCapability(ctx context.Context, params *protocol.RegistrationParams) (err error) {
	panic("unimplemented")
}

// ShowMessage implements protocol.Client.
func (ClientRecorder) ShowMessage(ctx context.Context, params *protocol.ShowMessageParams) (err error) {
	panic("unimplemented")
}

// ShowMessageRequest implements protocol.Client.
func (ClientRecorder) ShowMessageRequest(ctx context.Context, params *protocol.ShowMessageRequestParams) (result *protocol.MessageActionItem, err error) {
	panic("unimplemented")
}

// Telemetry implements protocol.Client.
func (ClientRecorder) Telemetry(ctx context.Context, params interface{}) (err error) {
	panic("unimplemented")
}

// UnregisterCapability implements protocol.Client.
func (ClientRecorder) UnregisterCapability(ctx context.Context, params *protocol.UnregistrationParams) (err error) {
	panic("unimplemented")
}

// WorkDoneProgressCreate implements protocol.Client.
func (ClientRecorder) WorkDoneProgressCreate(ctx context.Context, params *protocol.WorkDoneProgressCreateParams) (err error) {
	panic("unimplemented")
}

// WorkspaceFolders implements protocol.Client.
func (ClientRecorder) WorkspaceFolders(ctx context.Context) (result []protocol.WorkspaceFolder, err error) {
	panic("unimplemented")
}

type TestDidChangeTest struct {
	text string
}

func TestDidChange(t *testing.T) {

	mock := MockServer{}
	recorder := ClientRecorder{}
	server, init := NewServer(zap.NewNop(), mock)
	init(recorder)

	tests := []TestDidChangeTest{
		{
			text: "hello world",
		},
	}

	for _, test := range tests {

		params := protocol.DidChangeTextDocumentParams{}
		params.ContentChanges = []protocol.TextDocumentContentChangeEvent{
			{

				Text: test.text,
			},
		}

		server.DidChange(context.Background(), &params)
	}
}
