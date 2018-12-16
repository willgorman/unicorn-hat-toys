all:
	GOOS=linux GOARM=6 GOARCH=arm go build

.PHONY: clean
clean:
	rm ./unicorn-hat

unicorn-hat:
	GOOS=linux GOARM=6 GOARCH=arm go build

.PHONY: sync
sync: unicorn-hat
	scp unicorn-hat pi@192.168.1.36:go/bin/unicorn-hat

.PHONY: run
run: sync
	ssh -t pi@192.168.1.36 "go/bin/unicorn-hat"