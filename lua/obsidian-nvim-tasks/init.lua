-- register a neovim function to print hello and attach  to command named :Hello

local api = vim.api


local function hello()
  api.nvim_out_write("Hello, Neovim!\n")
end

api.nvim_create_user_command("Hello", hello, {})

