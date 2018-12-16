GOOS=linux GOARM=6 GOARCH=arm go build
scp unicorn-hat pi@192.168.1.36:go/bin/unicorn-hat
ssh -t pi@192.168.1.36 "go/bin/unicorn-hat"