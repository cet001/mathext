all : install

clean :
	@echo ">>> Cleaning and initializing mathext project <<<"
	@go clean
	@gofmt -w .

test : clean
	@echo ">>> Running unit tests <<<"
	@go test ./ ./ints ./vectors

test-coverage : clean
	@echo ">>> Running unit tests and calculating code coverage <<<"
	@go test ./ ./ints ./vectors -cover

install : test
	@echo ">>> Building and installing mathext <<<"
	@go install
	@echo ">>> mathext installed successfully! <<<"
	@echo ""
