local api = vim.api
local finder = require("obsidian-nvim-tasks.telescope")

local M = {}
function M.RegisterDefaultKeybingins()
  api.nvim_set_keymap('n', '<leader>nt', '',
    { noremap = true, silent = true, callback = function() finder.AllNotDoneTasks() end })
end

return M
