-- compile program in ../../cli directory with go build

local job = require('plenary.job')
local api = vim.api
local utils = require("obsidian-nvim-tasks.utils")

local M = {
  Install = function()
    local jj = job:new({
      cwd = utils.get_package_path() .. '/cli',
      command = 'go',
      args = { 'build', '.' },
      on_stderr = function(j, err)
        -- execute in main runtime
        vim.schedule(function()
          api.nvim_err_writeln(err)
        end)
      end,
      on_exit = function(j, return_val)
        vim.schedule(function()
          --api.nvim_err_writeln(vim.inspect(j))
        end)
      end
    })
    jj:start()
    jj:join()
  end,
}

return M
