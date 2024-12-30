FROM fedora:42

RUN dnf install -y golang neovim git curl

RUN git clone https://github.com/NvChad/starter ~/.config/nvim && nvim
