let mapleader = " "

" Install vundle
" ===========================================================================begin
let shouldInstallBundles = 0

if !filereadable($HOME . "/.config/nvim/autoload/plug.vim")
  echo "~≥ Installing vim-plug \n"
  silent !curl -fLo $HOME/.config/nvim/autoload/plug.vim --create-dirs https://raw.githubusercontent.com/junegunn/vim-plug/master/plug.vim
  let shouldInstallBundles = 1
endif
" ===========================================================================end

" Plugins
" ===========================================================================begin
call plug#begin('~/.config/nvim/plugged')
Plug 'preservim/nerdtree'                                              " NerdTree
Plug 'vim-airline/vim-airline'                                         " Status bar
Plug 'neoclide/coc.nvim'                                               " Auto Completion
Plug 'preservim/tagbar'                                                " Tagbar for code navigation
Plug 'tpope/vim-commentary'                                            " Comment code
Plug 'elzr/vim-json', {'for' : 'json'}                                 " Json highlight
Plug 'pangloss/vim-javascript', { 'for': 'javascript' }                " Javascript syntax highlighting
Plug 'maxmellon/vim-jsx-pretty', { 'for': 'javascript' }               " React & Typescript syntax highlighting
Plug 'airblade/vim-gitgutter'                                          " Show git modification info 
Plug 'tpope/vim-fugitive'
Plug 'tpope/vim-rhubarb'
Plug 'junegunn/gv.vim'
Plug 'frazrepo/vim-rainbow'
Plug 'junegunn/fzf', { 'dir': '~/.fzf', 'do': './install --all' }
Plug 'junegunn/fzf.vim'
Plug 'carlitux/deoplete-ternjs', { 'do': 'yarn global add tern' }
Plug 'ternjs/tern_for_vim', { 'for': 'javascript' }
Plug 'hail2u/vim-css3-syntax', { 'for': [ 'css', 'scss', 'sass'] }
Plug 'tpope/vim-markdown', { 'for': 'markdown' }
Plug 'ap/vim-css-color'                                               " CSS Color Preview
call plug#end()
" ===========================================================================end

" Normal Settings
" ===========================================================================begin
:set number
:set autoindent
:set tabstop=4
:set shiftwidth=4
:set smarttab
:set softtabstop=4
:set history=10000
:set mouse=v
set clipboard=unnamedplus
" ===========================================================================end

" Theme
" ===========================================================================begin
" Git Diff
highlight DiffAdd      gui=none    guifg=NONE          ctermbg=22
highlight DiffChange   gui=none    guifg=NONE          ctermbg=130
highlight DiffDelete   gui=bold    guifg=#ff8080       ctermbg=88
highlight DiffText     gui=none    guifg=NONE          ctermbg=18
" ===========================================================================end

" NerdTree Config 
" ===========================================================================begin
" Start NERDTree and put the cursor back in the other window.
autocmd VimEnter * NERDTree | wincmd p
" Exit Vim if NERDTree is the only window remaining in the only tab.
autocmd BufEnter * if tabpagenr('$') == 1 && winnr('$') == 1 && exists('b:NERDTree') && b:NERDTree.isTabTree() | quit | endif

let g:NERDTreeDirArrowExpandable = '+'
let g:NERDTreeDirArrowCollapsible = '-'
" ===========================================================================end

" Vim-Rainbow Config
" ===========================================================================begin
let g:rainbow_active = 1
" ===========================================================================end

" Gitgutter Config
" ===========================================================================begin
let g:gitgutter_async=0
let g:gitgutter_enabled = 1
let g:gitgutter_map_keys = 0
let g:gitgutter_sign_added = '+'
let g:gitgutter_sign_modified = 'm'
let g:gitgutter_sign_removed = '-'
" ===========================================================================end

" Airline Config
" ===========================================================================begin
let g:airline#extensions#wordcount#enabled = 1
let g:airline#extensions#hunks#non_zero_only = 1
" ===========================================================================end

" Key Binding 
" ===========================================================================begin
nnoremap ( <Plug>(GitGutterPrevHunk)
nnoremap ) <Plug>(GitGutterNextHunk)
nnoremap <leader>n :NERDTreeFocus<CR>
map <C-[> <C-o>
map <C-]> <C-i>
nnoremap <C-d> :GitGutterDiffOrig<CR>
nnoremap <C-7> :TagbarToggle<CR>
nnoremap <C-n> :NERDTree<CR>
nnoremap <C-t> :NERDTreeToggle<CR>
" Remap for rename current word
nmap <leader>rn <Plug>(coc-rename)
" ===========================================================================end


" Notes
" ===========================================================================begin
" Switch between tabs: Next Tab: <gt> , Prior Tab: <gT>
" Rename variable, functions: <gd> , <rn> "rename" <Esc> <.>
" Change project root: <leader> <n> <C>
" ===========================================================================end
