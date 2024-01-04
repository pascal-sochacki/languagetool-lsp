package server

import (
	"context"
	"fmt"
	"testing"

	"github.com/pascal-sochacki/languagetool-lsp/pkg/languagetool"
	"go.lsp.dev/protocol"
	"go.uber.org/zap"
)

type MockServer struct {
	result *languagetool.CheckResult
}

func (m *MockServer) CheckText(ctx context.Context, text string, language string) (languagetool.CheckResult, error) {
	return *m.result, nil
}

func (m *MockServer) setCheckResult(result languagetool.CheckResult) {
	m.result = &result
}

type ClientRecorder struct {
	Diagostics []protocol.PublishDiagnosticsParams
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
func (c *ClientRecorder) PublishDiagnostics(ctx context.Context, params *protocol.PublishDiagnosticsParams) (err error) {
	fmt.Printf("%+v", *params)
	c.Diagostics = append(c.Diagostics, *params)
	return nil
}

func (c ClientRecorder) getDiagostics() []protocol.PublishDiagnosticsParams {
	return c.Diagostics
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
	text   string
	answer languagetool.CheckResult
	expect []protocol.Diagnostic
}

func TestDidChange(t *testing.T) {
	tests := []TestDidChangeTest{
		{
			text: "Apffelstaft\n",
			answer: languagetool.CheckResult{
				Matches: []languagetool.Match{
					{
						Message: "Möglicher Tippfehler gefunden.",
						Offset:  0,
						Length:  10,
					},
				},
			},
			expect: []protocol.Diagnostic{
				{
					Message: "Möglicher Tippfehler gefunden.",
					Range: protocol.Range{
						Start: protocol.Position{Character: 0},
						End:   protocol.Position{Character: 10},
					},
				},
			},
		},
		{
			text: "\nApffelstaft",
			answer: languagetool.CheckResult{
				Matches: []languagetool.Match{
					{
						Message: "Möglicher Tippfehler gefunden.",
						Offset:  0,
						Length:  10,
					},
				},
			},
			expect: []protocol.Diagnostic{
				{
					Message: "Möglicher Tippfehler gefunden.",
					Range: protocol.Range{
						Start: protocol.Position{Character: 0, Line: 1},
						End:   protocol.Position{Character: 10, Line: 1},
					},
				},
			},
		},
	}

	for _, test := range tests {
		mock := &MockServer{}

		recorder := &ClientRecorder{}
		server, init := NewServer(zap.NewNop(), mock)
		init(recorder)
		params := protocol.DidChangeTextDocumentParams{}
		params.ContentChanges = []protocol.TextDocumentContentChangeEvent{
			{

				Text: test.text,
			},
		}
		fmt.Printf("%+v", test.answer)

		mock.setCheckResult(test.answer)
		server.DidChange(context.Background(), &params)

		if len(recorder.getDiagostics()) != len(test.expect) {
			t.Fatalf("wrong length of diagostics want: %d, got: %d", len(test.expect), len(recorder.getDiagostics()))
		}
	}
}

type TestCodeActionTest struct {
	text   string
	answer languagetool.CheckResult
}

func TestCodeAction(t *testing.T) {

	tests := []TestCodeActionTest{
		{
			text: "Das ist ein Text mit Fehlern",
			answer: languagetool.CheckResult{
				Matches: []languagetool.Match{
					{
						Message: "Der die das",
						Offset:  0,
						Length:  3,
					},
				},
			},
		},
	}

	for _, test := range tests {
		mock := &MockServer{}

		recorder := &ClientRecorder{}
		server, init := NewServer(zap.NewNop(), mock)
		init(recorder)
		params := protocol.DidChangeTextDocumentParams{}
		params.ContentChanges = []protocol.TextDocumentContentChangeEvent{
			{

				Text: test.text,
			},
		}
		fmt.Printf("%+v", test.answer)

		mock.setCheckResult(test.answer)
		server.DidChange(context.Background(), &params)
		actions, _ := server.CodeAction(context.Background(), &protocol.CodeActionParams{
			Range: protocol.Range{
				Start: protocol.Position{},
				End:   protocol.Position{Character: 3},
			},
		})

		if 1 != len(actions) {
			t.Fatalf("wrong length of actions want: %d, got: %d", 1, len(recorder.getDiagostics()))
		}

	}
}
