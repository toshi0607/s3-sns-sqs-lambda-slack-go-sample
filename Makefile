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
		--s3-bucket stack-bucket-for-s3-sns-sqs-lambda-slack-go-sample \
		--output-template-file sam.yml
	sam deploy \
		--template-file sam.yml \
		--stack-name stack-s3-sns-sqs-lambda-slack-go-sample \
		--capabilities CAPABILITY_IAM \
		--parameter-overrides \
		  WebhookURL=$(WEBHOOK_URL) \
		  Channel=$(CHANNEL) \
		  UserName=$(USER_NAME) \
		  Icon=$(ICON)

delete:
	aws s3 rm s3://sqs-sns-lambda-sample --recursive
	aws cloudformation delete-stack --stack-name stack-s3-sns-sqs-lambda-slack-go-sample
	aws s3 rm s3://stack-bucket-for-s3-sns-sqs-lambda-slack-go-sample --recursive
	aws s3 rb s3://stack-bucket-for-s3-sns-sqs-lambda-slack-go-sample
