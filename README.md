# golang-aws-lamda

[![Test](https://github.com/taaaaakahiro/golang-aws-lambda-colly-discordgo/actions/workflows/test.yml/badge.svg)](https://github.com/taaaaakahiro/golang-aws-lambda-colly-discordgo/actions/workflows/test.yml)

## Set Up
```
$ make build
$ cd terrform
$ terraform applu -auto-approve
```

## Invoke from local
```
$ aws lambda invoke --function-name golang-lambda-terraform-example out --log-type Tail
```

