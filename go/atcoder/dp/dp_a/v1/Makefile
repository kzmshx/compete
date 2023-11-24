SHELL := /bin/bash

GO_VERSION   := 1.20.6
URL          := "https://atcoder.jp/contests/dp/tasks/dp_a"
MAIN         := main.go
TESTCASE_DIR := ../testcases

go.mod:
	go${GO_VERSION} mod init

$(TESTCASE_DIR):
	mkdir -p ../testcases
	oj d ${URL} -d ${TESTCASE_DIR}

.PHONY: init
init: go.mod $(TESTCASE_DIR)

.PHONY: tidy
tidy:
	go${GO_VERSION} mod tidy

.PHONY: test
test: init tidy
	oj t -c "go${GO_VERSION} run ${MAIN}" -d ${TESTCASE_DIR} -j 10

.PHONY: pretty
pretty: test
	godead ${MAIN} && go${GO_VERSION} fmt

.PHONY: submit
submit: pretty
	oj s ${URL} ${MAIN} -y
