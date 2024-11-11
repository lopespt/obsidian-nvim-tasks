require("obsidian-nvim-tasks.init")
local installer = require("obsidian-nvim-tasks.installer")

function R(module)
  package.loaded[module] = nil
  return require(module)
end

function Reload()
  for key, value in pairs(package.loaded) do
    -- if key constains obsidian
    if string.find(key, "obsidian--nvim") then
      vim.api.nvim_out_write(key .. "\n")
      R(key)
    end
  end
end

installer.Install()
