# gobuild-git
build program version with git information

not support windows

#Usage

add variables to your package main file

* default names:
>version name: \_version_\
>branch name: \_branch_\
>commitId name: \_commitId_\
>buildTime name: \_buildTime_\

run command: 

`gobuild-git start -m main.go -i main`