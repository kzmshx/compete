GO_MAIN    := main.go
GO_VERSION :=
GO         := go$(GO_VERSION)

PROBLEM_URL          := "https://atcoder.jp/contests/abc191/tasks/abc191_b"
TESTCASE_DIR := ../testcases

$(TESTCASE_DIR):
	mkdir -p $(TESTCASE_DIR)
	oj d $(PROBLEM_URL) -d $(TESTCASE_DIR)

go.mod:
	$(GO) mod init

go.sum:
	$(GO) mod tidy

.PHONY: init
init: $(TESTCASE_DIR) go.mod go.sum

.PHONY: test
test: init
	oj t -c "$(GO) run $(GO_MAIN)" -d $(TESTCASE_DIR) -t 2 -j 10

.PHONY: fmt
fmt: test
	godead $(GO_MAIN)
	$(GO) mod tidy
	$(GO) fmt

.PHONY: submit
submit: fmt
	oj s $(PROBLEM_URL) $(GO_MAIN) -y

.PHONY: commit
commit: fmt
	git add .
	git commit -m "$(PROBLEM_URL)"
	git push
