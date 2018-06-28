.PHONY: all
all: clean
	go install ./... && \
	cli generate -dir=testdata example echo && \
	go run testdata/main.go testdata/echo_command.go echo

.PHONY: clean
clean:
ifeq ($(SHELL),cmd)
	if exist "testdata" for /d %%G in ("testdata") do rd /s /q "%%G"
	if not exist "testdata" mkdir testdata
else
	rm -rf testdata
	mkdir -p testdata
endif
