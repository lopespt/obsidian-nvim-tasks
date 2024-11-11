local api = vim.api
local finder = require("obsidian-nvim-tasks.telescope")

function registerKeybingins()
  api.nvim_set_keymap('n', '<leader>nt', '',
    { noremap = true, silent = true, callback = function() finder.AllNotDoneTasks() end })
end
