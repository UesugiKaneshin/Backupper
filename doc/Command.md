## Golang
go run main.go
go build -o backupper main.go
cd ~/e/Backupper/src/ & go build -o ../dist/backupper main.go

## Linux
netstat -tlnp | grep 'LISTEN' | grep -i '10000'
netstat -tlnp
nohup ./backupper &

## Git
git log --graph --pretty=oneline --abbrev-commit
git merge --no-ff -m "" feature
git merge --squash develop
