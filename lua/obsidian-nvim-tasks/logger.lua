local M = {}

local api = vim.api

function M.Info(msg)
  api.nvim_out_write("ObsidianNvimTasks: " .. msg .. "\n")
end

function M.Error(msg)
  api.nvim_err_write("ObsidianNvimTasks: " .. msg .. "\n")
end

return M
