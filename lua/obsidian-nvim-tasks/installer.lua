-- compile program in ../../cli directory with go build

local job = require('plenary.job')
local api = vim.api
local utils = require("obsidian-nvim-tasks.utils")
local logger = require("obsidian-nvim-tasks.logger")

local M = {}

function M.Install()
  local jj = job:new({
    cwd = utils.get_package_path() .. '/cli',
    command = 'make',
    on_stderr = function(j, err)
      vim.schedule(function()
        logger.Error(err)
      end)
    end,
    on_exit = function(j, return_val)
      vim.schedule(function()
        if return_val ~= 0 then
          logger.Error("Cli tool installation failed")
        end
      end)
    end
  })
  jj:start()
end

return M
