SHELL := /bin/bash

URL          := "https://atcoder.jp/contests/math-and-algorithm/tasks/math_and_algorithm_h"
MAIN         := main.go
TESTCASE_DIR := ../testcases

go.mod:
	go mod init

$(TESTCASE_DIR):
	mkdir -p ../testcases
	oj d ${URL} -d ${TESTCASE_DIR}

.PHONY: init
init: go.mod $(TESTCASE_DIR)

.PHONY: test
test: init
	oj t -c "go run ${MAIN}" -d ${TESTCASE_DIR} -j 10

.PHONY: pretty
pretty: test
	godead ${MAIN}
	go fmt

.PHONY: submit
submit: pretty
	oj s ${URL} ${MAIN} -y
