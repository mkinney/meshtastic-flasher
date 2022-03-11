Notes to get started:
 go get -u github.com/spf13/cobra@latest
 go install github.com/spf13/cobra-cli@latest
 cd $GOPATH/src/github.com/mkinney
 mkdir meshtastic-flasher
 cd meshtastic-flasher
 go mod init meshtastic-flasher
 cobra-cli init -a "Mike Kinney"

 See https://taskfile.dev/#/installation
  brew install go-task/tap/go-task
 or
  go install github.com/go-task/task/v3/cmd/task@latest
