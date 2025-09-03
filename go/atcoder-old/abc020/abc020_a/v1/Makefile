GO_VERSION :=
GO         := go$(GO_VERSION)

PROBLEM_URL  := "https://atcoder.jp/contests/abc020/tasks/abc020_a"
MAIN_FILE    := main.go
TESTCASE_DIR := ../testcases

$(TESTCASE_DIR):
	mkdir -p $@ && oj d $(PROBLEM_URL) -d $@

.PHONY: init test fmt submit commit

init: $(TESTCASE_DIR)

test: init
	oj t -c "$(GO) run $(MAIN_FILE)" -d $(TESTCASE_DIR) -t 2 -j 10

fmt: test
	godead $(MAIN_FILE) && $(GO) fmt

submit: fmt
	oj s $(PROBLEM_URL) $(MAIN_FILE) -y

commit: fmt
	git add . && git commit -m "$(PROBLEM_URL)" && git push
