# zsh completion for aict

autoload -U compinit
compinit

local -a tools
tools=(
  'ls:Directory listings'
  'cat:File contents'
  'grep:Pattern search'
  'find:Filesystem search'
  'stat:File metadata'
  'wc:Line/word count'
  'diff:File comparison'
  'file:Type detection'
  'head:First lines'
  'tail:Last lines'
  'du:Disk usage'
  'df:Filesystem stats'
  'realpath:Resolve path'
  'basename:Extract filename'
  'dirname:Extract directory'
  'pwd:Working directory'
  'sort:Sort lines'
  'uniq:Remove duplicates'
  'cut:Extract columns'
  'tr:Translate characters'
  'env:Environment variables'
  'system:System information'
  'ps:Process list'
  'checksums:Hash computation'
)

aict_tools() {
    local -a commands
    local cmd
    for cmd in $tools; do
        commands+=("${cmd%%:*}")
    done
    _describe 'command' commands
}

_aict() {
    case $CURRENT in
        2)
            aict_tools
            ;;
        *)
            _files
            ;;
    esac
}

compdef _aict aict
