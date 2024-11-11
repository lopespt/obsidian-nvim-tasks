local M = {}

local fn = vim.fn

function M.get_package_path()
  -- Path to this source file, removing the leading '@'
  local source = string.sub(debug.getinfo(1, "S").source, 2)

  -- Path to the package root
  return fn.fnamemodify(source, ":p:h:h:h")
end

function M.get_executable_cli()
  return M.get_package_path() .. '/cli/obsidian-tasks.git'
end

return M
