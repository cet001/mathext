all : install

clean :
	@echo ">>> Cleaning and initializing intmath project <<<"
	@go clean
	@gofmt -w .
	@go get github.com/stretchr/testify

test : clean
	@echo ">>> Running unit tests <<<"
	@go test ./ ./ints

test-coverage : clean
	@echo ">>> Running unit tests and calculating code coverage <<<"
	@go test ./ ./ints -cover

install : test
	@echo ">>> Building and installing intmath <<<"
	@go install
	@echo ">>> intmath installed successfully! <<<"
	@echo ""
