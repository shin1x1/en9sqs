# en9sqs

This tool enqueue messages to Amazon SQS. The main purpose of en9sqs is testing job workers performance.


## Usage

```
$ go get -u github.com/shin1x1/en9sqs
$ en9sqs -h
NAME:
   en9sqs - enqueue messages to Amazon SQS tool

USAGE:
   en9sqs [global options] command [command options] queue_url message

VERSION:
   0.0.1

COMMANDS:
     help, h  Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --region value, -r value       Region (default: "ap-northeast-1")
   --concurrency value, -c value  Number of concurrency workers (default: "1")
   --messages value, -n value     Number of enqueueing messages (default: "1")
   --help, -h                     show help
   --version, -v                  print the version
```


## Build

```
$ git clone this_repo
$ cd this_repo
$ make
```
