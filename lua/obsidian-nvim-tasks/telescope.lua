local pickers = require "telescope.pickers"
local finders = require "telescope.finders"
local conf = require("telescope.config").values
local utils = require "obsidian-nvim-tasks.utils"

local M = {}

local function createTaskFromLine(line)
  local task = vim.fn.json_decode(line)
  return {
    value = task,
    ordinal = task.description,
    display = task.status .. " - " .. task.description,
    filename = task.context.filename,
    lnum = task.context.lnum,
  }
end

M.AllObsidianTasks = function(opts)
  opts = opts or {}
  pickers.new(opts, {
    prompt_title = "Tasks",
    finder = finders.new_oneshot_job({ utils.get_executable_cli() }, {
      entry_maker = function(line)
        --convert json to table
        return createTaskFromLine(line)
      end
    }),
    sorter = conf.generic_sorter(opts),
  }):find()
end


M.AllNotDoneTasks = function(opts)
  opts = opts or {}
  pickers.new(opts, {
    prompt_title = "Tasks: Not Done",
    finder = finders.new_oneshot_job({ utils.get_executable_cli() }, {
      entry_maker = function(line)
        return createTaskFromLine(line)
      end
    }),
    sorter = conf.generic_sorter(opts),
  }):find()
end

-- to execute the function

M.AllObsidianTasks()

return M
