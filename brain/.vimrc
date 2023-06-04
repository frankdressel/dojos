set encoding=utf-8
set nocompatible
filetype off

set rtp+=~/.vim/bundle/Vundle.vim/
call vundle#begin()

" let Vundle manage Vundle
" required!
Plugin 'VundleVim/Vundle.vim'
Plugin 'vim-airline/vim-airline'
Plugin 'scrooloose/nerdtree'
Plugin 'jnurmine/Zenburn'
Plugin 'scrooloose/syntastic'
Plugin 'kien/ctrlp.vim'
Plugin 'gabrielelana/vim-markdown'
Plugin 'Valloric/YouCompleteMe'
Plugin 'Raimondi/delimitMate'
Plugin 'SirVer/ultisnips'
Plugin 'honza/vim-snippets'
Plugin 'preservim/tagbar'
Plugin 'udalov/kotlin-vim'
call vundle#end()

filetype plugin indent on
autocmd FileType yaml setl indentkeys-=<:>

syntax on

" Powerline setup
set guifont=DejaVu\ Sans\ Mono\ for\ Powerline\ 9
set laststatus=2

map <F2> :NERDTreeToggle<CR>

set autochdir

set tabstop=4
set shiftwidth=4
set expandtab

set t_Co=256
" Taken from: https://sanctum.geek.nz/arabesque/gracefully-degrading-vimrc/
silent! colors zenburn

set number

" Ctrl-P
nmap ; :CtrlPBuffer<CR>

" Syntastic
set statusline+=%#warningmsg#
set statusline+=%{SyntasticStatuslineFlag()}
set statusline+=%*

let g:syntastic_always_populate_loc_list = 1
let g:syntastic_auto_loc_list = 1
let g:syntastic_check_on_open = 1
let g:syntastic_check_on_wq = 0

" Taken from: https://robots.thoughtbot.com/vim-splits-move-faster-and-more-naturally
" * Ctrl+j/k lets you navigate the splits
nmap <C-J> <C-W><C-J>
nmap <C-K> <C-W><C-K>
nmap <C-L> <C-W><C-L>
nmap <C-H> <C-W><C-H>
" * Split is right and below
set splitbelow
set splitright

" Taken from: http://statico.github.io/vim.html
nmap j gj
nmap k gk

" Tagbar
nmap <F8> :TagbarToggle<CR>

set incsearch
set ignorecase
set smartcase
set hlsearch
nmap \q :nohlsearch<CR>

let g:ycm_language_server = [
  \   { 'name': 'kotlin',
  \     'filetypes': [ 'kotlin' ],
  \     'cmdline': [ '/progs/kotlin-language-server/server/build/install/server/bin/kotlin-language-server' ],
  \   },
  \ ]
