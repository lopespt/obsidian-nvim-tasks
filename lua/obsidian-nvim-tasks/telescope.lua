local pickers = require "telescope.pickers"
local finders = require "telescope.finders"
local conf = require("telescope.config").values
local utils = require "obsidian-nvim-tasks.utils"

local M = {}

local function formatDate(icon, date)
  -- receive 2024-01-01T00:00:00Z date time and returns only the date
  outdt = date:sub(1, 10)
  if outdt == '' then
    return ""
  end
  return icon .. outdt
end

local function createTaskFromLine(line)
  local task = vim.fn.json_decode(line)
  return {
    value = task,
    ordinal = task.description,
    display = formatDate("📅", task.dueDate or "") ..
        formatDate("⏳", task.scheduledDate or "") .. "[" .. task.status .. "] " .. task.description,
    filename = task.context.filename,
    lnum = task.context.lnum,
  }
end

M.AllObsidianTasks = function(opts)
  opts = opts or {}
  pickers.new(opts, {
    prompt_title = "Tasks",
    finder = finders.new_oneshot_job({ utils.get_executable_cli(), "-v", "/Users/gwachs/obsidian/gwachs/" }, {
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
    finder = finders.new_oneshot_job(
      { utils.get_executable_cli(), "-v", "/Users/gwachs/obsidian/gwachs/", "-ns", "x", "--SortAnyDate" }, {
        entry_maker = function(line)
          return createTaskFromLine(line)
        end
      }),
    sorter = conf.generic_sorter(opts),
  }):find()
end

-- to execute the function

M.AllNotDoneTasks()
--M.AllObsidianTasks()

return M
