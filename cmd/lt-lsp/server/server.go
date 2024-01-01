package server

import (
	"context"
	"fmt"

	"go.lsp.dev/protocol"
	"go.uber.org/zap"
)

type Server struct {
	log    *zap.Logger
	client protocol.Client
}

// CodeAction implements protocol.Server.
func (Server) CodeAction(ctx context.Context, params *protocol.CodeActionParams) (result []protocol.CodeAction, err error) {
	return []protocol.CodeAction{}, nil
}

// CodeLens implements protocol.Server.
func (Server) CodeLens(ctx context.Context, params *protocol.CodeLensParams) (result []protocol.CodeLens, err error) {
	return []protocol.CodeLens{}, nil
}

// CodeLensRefresh implements protocol.Server.
func (Server) CodeLensRefresh(ctx context.Context) (err error) {
	return nil
}

// CodeLensResolve implements protocol.Server.
func (Server) CodeLensResolve(ctx context.Context, params *protocol.CodeLens) (result *protocol.CodeLens, err error) {
	return &protocol.CodeLens{}, nil
}

// ColorPresentation implements protocol.Server.
func (Server) ColorPresentation(ctx context.Context, params *protocol.ColorPresentationParams) (result []protocol.ColorPresentation, err error) {
	return []protocol.ColorPresentation{}, nil
}

// Completion implements protocol.Server.
func (Server) Completion(ctx context.Context, params *protocol.CompletionParams) (result *protocol.CompletionList, err error) {
	return &protocol.CompletionList{}, nil
}

// CompletionResolve implements protocol.Server.
func (Server) CompletionResolve(ctx context.Context, params *protocol.CompletionItem) (result *protocol.CompletionItem, err error) {
	return &protocol.CompletionItem{}, nil
}

// Declaration implements protocol.Server.
func (Server) Declaration(ctx context.Context, params *protocol.DeclarationParams) (result []protocol.Location, err error) {
	return []protocol.Location{}, nil
}

// Definition implements protocol.Server.
func (Server) Definition(ctx context.Context, params *protocol.DefinitionParams) (result []protocol.Location, err error) {
	return nil, nil
}

// DidChange implements protocol.Server.
func (s Server) DidChange(ctx context.Context, params *protocol.DidChangeTextDocumentParams) (err error) {
	s.log.Debug(fmt.Sprintf("%+v", params))
	s.log.Debug(fmt.Sprintf("%+v", s.client))
	s.client.PublishDiagnostics(ctx, &protocol.PublishDiagnosticsParams{
		URI: params.TextDocument.URI,
		Diagnostics: []protocol.Diagnostic{
			{
				Range: protocol.Range{
					Start: protocol.Position{},
					End:   protocol.Position{},
				},
				Message: "Das ist ein Test",
			},
		},
	})
	return nil
}

// DidChangeConfiguration implements protocol.Server.
func (s Server) DidChangeConfiguration(ctx context.Context, params *protocol.DidChangeConfigurationParams) (err error) {
	s.log.Debug(fmt.Sprintf("%+v", params))
	return nil
}

// DidChangeWatchedFiles implements protocol.Server.
func (s Server) DidChangeWatchedFiles(ctx context.Context, params *protocol.DidChangeWatchedFilesParams) (err error) {
	s.log.Debug(fmt.Sprintf("%+v", params))
	return nil
}

// DidChangeWorkspaceFolders implements protocol.Server.
func (s Server) DidChangeWorkspaceFolders(ctx context.Context, params *protocol.DidChangeWorkspaceFoldersParams) (err error) {
	s.log.Debug(fmt.Sprintf("%+v", params))
	return nil
}

// DidClose implements protocol.Server.
func (Server) DidClose(ctx context.Context, params *protocol.DidCloseTextDocumentParams) (err error) {
	return nil
}

// DidCreateFiles implements protocol.Server.
func (Server) DidCreateFiles(ctx context.Context, params *protocol.CreateFilesParams) (err error) {
	return nil
}

// DidDeleteFiles implements protocol.Server.
func (Server) DidDeleteFiles(ctx context.Context, params *protocol.DeleteFilesParams) (err error) {
	return nil
}

// DidOpen implements protocol.Server.
func (s Server) DidOpen(ctx context.Context, params *protocol.DidOpenTextDocumentParams) (err error) {
	s.log.Debug(params.TextDocument.Text)

	return nil
}

// DidRenameFiles implements protocol.Server.
func (Server) DidRenameFiles(ctx context.Context, params *protocol.RenameFilesParams) (err error) {
	return nil
}

// DidSave implements protocol.Server.
func (s Server) DidSave(ctx context.Context, params *protocol.DidSaveTextDocumentParams) (err error) {
	s.log.Debug(fmt.Sprintf("%+v", params))
	return nil
}

// DocumentColor implements protocol.Server.
func (Server) DocumentColor(ctx context.Context, params *protocol.DocumentColorParams) (result []protocol.ColorInformation, err error) {
	return nil, nil
}

// DocumentHighlight implements protocol.Server.
func (Server) DocumentHighlight(ctx context.Context, params *protocol.DocumentHighlightParams) (result []protocol.DocumentHighlight, err error) {
	return nil, nil
}

// DocumentLink implements protocol.Server.
func (Server) DocumentLink(ctx context.Context, params *protocol.DocumentLinkParams) (result []protocol.DocumentLink, err error) {
	return nil, nil
}

// DocumentLinkResolve implements protocol.Server.
func (Server) DocumentLinkResolve(ctx context.Context, params *protocol.DocumentLink) (result *protocol.DocumentLink, err error) {
	return nil, nil
}

// DocumentSymbol implements protocol.Server.
func (Server) DocumentSymbol(ctx context.Context, params *protocol.DocumentSymbolParams) (result []interface{}, err error) {
	return nil, nil
}

// ExecuteCommand implements protocol.Server.
func (Server) ExecuteCommand(ctx context.Context, params *protocol.ExecuteCommandParams) (result interface{}, err error) {
	return nil, nil
}

// Exit implements protocol.Server.
func (Server) Exit(ctx context.Context) (err error) {
	return nil
}

// FoldingRanges implements protocol.Server.
func (Server) FoldingRanges(ctx context.Context, params *protocol.FoldingRangeParams) (result []protocol.FoldingRange, err error) {
	return nil, nil
}

// Formatting implements protocol.Server.
func (Server) Formatting(ctx context.Context, params *protocol.DocumentFormattingParams) (result []protocol.TextEdit, err error) {
	return nil, nil
}

// Hover implements protocol.Server.
func (Server) Hover(ctx context.Context, params *protocol.HoverParams) (result *protocol.Hover, err error) {
	return nil, nil
}

// Implementation implements protocol.Server.
func (Server) Implementation(ctx context.Context, params *protocol.ImplementationParams) (result []protocol.Location, err error) {
	return nil, nil
}

// IncomingCalls implements protocol.Server.
func (Server) IncomingCalls(ctx context.Context, params *protocol.CallHierarchyIncomingCallsParams) (result []protocol.CallHierarchyIncomingCall, err error) {
	return nil, nil
}

// Initialize implements protocol.Server.
func (s Server) Initialize(ctx context.Context, params *protocol.InitializeParams) (result *protocol.InitializeResult, err error) {
	s.log.Debug("called Initialize")

	result = &protocol.InitializeResult{}

	result.ServerInfo = &protocol.ServerInfo{}
	result.ServerInfo.Name = "languagetool-lsp"
	result.ServerInfo.Version = "0.0.1"
	result.Capabilities = protocol.ServerCapabilities{}
	result.Capabilities.TextDocumentSync = 1
	result.Capabilities.CodeActionProvider = protocol.CodeActionOptions{ResolveProvider: true}
	return result, nil
}

// Initialized implements protocol.Server.
func (s Server) Initialized(ctx context.Context, params *protocol.InitializedParams) (err error) {
	s.log.Debug("called Initialized")
	return nil
}

// LinkedEditingRange implements protocol.Server.
func (Server) LinkedEditingRange(ctx context.Context, params *protocol.LinkedEditingRangeParams) (result *protocol.LinkedEditingRanges, err error) {
	return nil, nil
}

// LogTrace implements protocol.Server.
func (Server) LogTrace(ctx context.Context, params *protocol.LogTraceParams) (err error) {
	return nil
}

// Moniker implements protocol.Server.
func (Server) Moniker(ctx context.Context, params *protocol.MonikerParams) (result []protocol.Moniker, err error) {
	return nil, nil
}

// OnTypeFormatting implements protocol.Server.
func (Server) OnTypeFormatting(ctx context.Context, params *protocol.DocumentOnTypeFormattingParams) (result []protocol.TextEdit, err error) {
	return nil, nil
}

// OutgoingCalls implements protocol.Server.
func (Server) OutgoingCalls(ctx context.Context, params *protocol.CallHierarchyOutgoingCallsParams) (result []protocol.CallHierarchyOutgoingCall, err error) {
	return nil, nil
}

// PrepareCallHierarchy implements protocol.Server.
func (Server) PrepareCallHierarchy(ctx context.Context, params *protocol.CallHierarchyPrepareParams) (result []protocol.CallHierarchyItem, err error) {
	return nil, nil
}

// PrepareRename implements protocol.Server.
func (Server) PrepareRename(ctx context.Context, params *protocol.PrepareRenameParams) (result *protocol.Range, err error) {
	return nil, nil
}

// RangeFormatting implements protocol.Server.
func (Server) RangeFormatting(ctx context.Context, params *protocol.DocumentRangeFormattingParams) (result []protocol.TextEdit, err error) {
	return nil, nil
}

// References implements protocol.Server.
func (Server) References(ctx context.Context, params *protocol.ReferenceParams) (result []protocol.Location, err error) {
	return nil, nil
}

// Rename implements protocol.Server.
func (Server) Rename(ctx context.Context, params *protocol.RenameParams) (result *protocol.WorkspaceEdit, err error) {
	return nil, nil
}

// Request implements protocol.Server.
func (s Server) Request(ctx context.Context, method string, params interface{}) (result interface{}, err error) {
	s.log.Debug(fmt.Sprintf("%+v", params))
	return nil, nil
}

// SemanticTokensFull implements protocol.Server.
func (Server) SemanticTokensFull(ctx context.Context, params *protocol.SemanticTokensParams) (result *protocol.SemanticTokens, err error) {
	return nil, nil
}

// SemanticTokensFullDelta implements protocol.Server.
func (Server) SemanticTokensFullDelta(ctx context.Context, params *protocol.SemanticTokensDeltaParams) (result interface{}, err error) {
	return nil, nil
}

// SemanticTokensRange implements protocol.Server.
func (Server) SemanticTokensRange(ctx context.Context, params *protocol.SemanticTokensRangeParams) (result *protocol.SemanticTokens, err error) {
	return nil, nil
}

// SemanticTokensRefresh implements protocol.Server.
func (Server) SemanticTokensRefresh(ctx context.Context) (err error) {
	return nil
}

// SetTrace implements protocol.Server.
func (Server) SetTrace(ctx context.Context, params *protocol.SetTraceParams) (err error) {
	return nil
}

// ShowDocument implements protocol.Server.
func (Server) ShowDocument(ctx context.Context, params *protocol.ShowDocumentParams) (result *protocol.ShowDocumentResult, err error) {
	return nil, nil
}

// Shutdown implements protocol.Server.
func (Server) Shutdown(ctx context.Context) (err error) {
	return nil
}

// SignatureHelp implements protocol.Server.
func (Server) SignatureHelp(ctx context.Context, params *protocol.SignatureHelpParams) (result *protocol.SignatureHelp, err error) {
	return &protocol.SignatureHelp{}, nil
}

// Symbols implements protocol.Server.
func (Server) Symbols(ctx context.Context, params *protocol.WorkspaceSymbolParams) (result []protocol.SymbolInformation, err error) {
	return []protocol.SymbolInformation{}, nil
}

// TypeDefinition implements protocol.Server.
func (Server) TypeDefinition(ctx context.Context, params *protocol.TypeDefinitionParams) (result []protocol.Location, err error) {
	return []protocol.Location{}, nil
}

// WillCreateFiles implements protocol.Server.
func (Server) WillCreateFiles(ctx context.Context, params *protocol.CreateFilesParams) (result *protocol.WorkspaceEdit, err error) {
	return &protocol.WorkspaceEdit{}, nil
}

// WillDeleteFiles implements protocol.Server.
func (Server) WillDeleteFiles(ctx context.Context, params *protocol.DeleteFilesParams) (result *protocol.WorkspaceEdit, err error) {
	return &protocol.WorkspaceEdit{}, nil
}

// WillRenameFiles implements protocol.Server.
func (Server) WillRenameFiles(ctx context.Context, params *protocol.RenameFilesParams) (result *protocol.WorkspaceEdit, err error) {
	return &protocol.WorkspaceEdit{}, nil
}

// WillSave implements protocol.Server.
func (Server) WillSave(ctx context.Context, params *protocol.WillSaveTextDocumentParams) (err error) {
	return nil
}

// WillSaveWaitUntil implements protocol.Server.
func (Server) WillSaveWaitUntil(ctx context.Context, params *protocol.WillSaveTextDocumentParams) (result []protocol.TextEdit, err error) {
	return []protocol.TextEdit{}, nil
}

// WorkDoneProgressCancel implements protocol.Server.
func (Server) WorkDoneProgressCancel(ctx context.Context, params *protocol.WorkDoneProgressCancelParams) (err error) {
	return nil
}

func NewServer(log *zap.Logger) (*Server, func(protocol.Client)) {
	a := &Server{
		log: log,
	}
	b := func(client protocol.Client) {
		a.client = client

	}
	return a, b
}
