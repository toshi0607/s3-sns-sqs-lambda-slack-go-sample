build: build-write-ext build-write-file-name build-notifier

build-write-ext:
	GOARCH=amd64 GOOS=linux go build -o artifact/write_ext ./handlers/write_ext

build-write-file-name:
	GOARCH=amd64 GOOS=linux go build -o artifact/write_file_name ./handlers/write_file_name

build-notifier:
	GOARCH=amd64 GOOS=linux go build -o artifact/notifier ./handlers/notifier

deploy: build
	sam package \
		--template-file template.yml \
		--s3-bucket s3-sns-sqs-lambda-slack-go-sample \
		--output-template-file sam.yml
	sam deploy \
		--template-file sam.yml \
		--stack-name stack-s3-sns-sqs-lambda-slack-go-sample \
		--capabilities CAPABILITY_IAM
