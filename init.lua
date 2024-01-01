vim.api.nvim_create_autocmd('FileType', {
    pattern = 'markdown',
    callback = function ()
        vim.lsp.start({
            name = 'languagetool-lsp',
            cmd = {'go',  'run',  'cmd/lt-lsp/main.go'},
            root_dir = vim.fs.dirname(vim.fs.find({ "main" }, {upward = true })[1]),
            filetypes = { "markdown", "text" },
        })
    end
})

