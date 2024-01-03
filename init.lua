local lsp_configurations = require('lspconfig.configs')

if not lsp_configurations.languagetool_lsp then
  lsp_configurations.languagetool_lsp = {
    default_config = {
      name = 'languagetool-lsp',
      cmd = {'go',  'run',  'cmd/lt-lsp/main.go'},
      filetypes = {'markdown'},
      root_dir = require('lspconfig.util').root_pattern('go.sum')
    }
  }
end


local lsp_on_attach = function(client, bufnr)
  local opts = {buffer = bufnr, remap = false}
  
  print("hhellllo")

  vim.keymap.set("n", "gd", function() vim.lsp.buf.definition() end, opts)
  vim.keymap.set("n", "K", function() vim.lsp.buf.hover() end, opts)
  vim.keymap.set("n", "<leader>vws", function() vim.lsp.buf.workspace_symbol() end, opts)
  vim.keymap.set("n", "<leader>vd", function() vim.diagnostic.open_float() end, opts)
  vim.keymap.set("n", "[d", function() vim.diagnostic.goto_next() end, opts)
  vim.keymap.set("n", "]d", function() vim.diagnostic.goto_prev() end, opts)
  vim.keymap.set("n", "<leader>vca", function() vim.lsp.buf.code_action() end, opts)
  vim.keymap.set("n", "<leader>vrr", function() vim.lsp.buf.references() end, opts)
  vim.keymap.set("n", "<leader>vrn", function() vim.lsp.buf.rename() end, opts)
  vim.keymap.set("i", "<C-h>", function() vim.lsp.buf.signature_help() end, opts)
end



require('lspconfig').languagetool_lsp.setup({
    on_attach = lsp_on_attach 
})
