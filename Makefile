STACK_NAME := stack-s3-sns-sqs-lambda-slack-go-sample
TEMPLATE_FILE := template.yml
SAM_FILE := sam.yml

build: build-write-ext build-write-file-name build-notifier
.PHONY: build

build-write-ext:
	GOARCH=amd64 GOOS=linux go build -o artifact/write_ext ./handlers/write_ext
.PHONY: build-write-ext

build-write-file-name:
	GOARCH=amd64 GOOS=linux go build -o artifact/write_file_name ./handlers/write_file_name
.PHONY: build-write-file-name

build-notifier:
	GOARCH=amd64 GOOS=linux go build -o artifact/notifier ./handlers/notifier
.PHONY: build-notifier

deploy: build
	sam package \
		--template-file $(TEMPLATE_FILE) \
		--s3-bucket $(STACK_BUCKET) \
		--output-template-file $(SAM_FILE)
	sam deploy \
		--template-file $(SAM_FILE) \
		--stack-name $(STACK_NAME) \
		--capabilities CAPABILITY_IAM \
		--parameter-overrides \
		  WebhookURL=$(WEBHOOK_URL) \
		  Channel=$(CHANNEL) \
		  UserName=$(USER_NAME) \
		  Icon=$(ICON) \
		  FileBucket=$(FILE_BUCKET)
.PHONY: deploy

delete:
	aws s3 rm "s3://$(FILE_BUCKET)" --recursive
	aws cloudformation delete-stack --stack-name $(STACK_NAME)
	aws s3 rm "s3://$(STACK_BUCKET)" --recursive
	aws s3 rb "s3://$(STACK_BUCKET)"
.PHONY: delete

test:
	go test ./...
.PHONY: test
