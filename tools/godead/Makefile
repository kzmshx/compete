GO_VERSION :=
GO := go${GO_VERSION}

go.mod:
	${GO} mod init

.PHONY: tidy
tidy:
	${GO} mod tidy

.PHONY: fmt
fmt:
	${GO} fmt

.PHONY: install
install:
	${GO} install .
