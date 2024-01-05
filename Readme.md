![Example](assets/example.png "Example of the usage")

# LanguageTool LSP

This is a simple LSP which wraps LanguageTool for correction of markdown (and in the future code comments)
The LSP just calls the public available API (https://languagetool.org/http-api/). This API has some Limitations, just read the Documentation.

## Motivation

I'm suffering from a little dyslexia, and sometimes I need to write markdowns, comments and commit messages... So I built this little helper.
Building Tools for me and other Developers brings me a lot of joy!

## Information

I do pay for LanguageTool for the premium feature, so I don't know how good the free version is 

## Install

```
go install cmd/lt-lsp/lt-lsp.go
```

### Neovim

```lua
local lsp_configurations = require('lspconfig.configs')

if not lsp_configurations.languagetool_lsp then
  lsp_configurations.languagetool_lsp = {
    default_config = {
      name = 'languagetool-lsp',
      cmd = {'lt-lsp'},
      filetypes = {'markdown'},
      root_dir = require('lspconfig.util').root_pattern('*')
    }
  }
end
```



