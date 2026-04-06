# bash completion for aict

_aict() {
    local cur prev words cword
    _init_completion || return

    local tools="ls cat grep find stat wc diff file head tail du df realpath basename dirname pwd sort uniq cut tr env system ps checksums"

    if [[ $cword -eq 1 ]]; then
        COMPREPLY=($(compgen -W "$tools" -- "$cur"))
        return
    fi

    case "${words[1]}" in
        ls)
            _filedir
            ;;
        grep|find)
            _filedir -d
            ;;
        cat|head|tail|wc)
            _filedir
            ;;
        stat)
            _filedir
            ;;
        diff)
            _filedir
            ;;
        *)
            ;;
    esac
} && complete -F _aict aict
