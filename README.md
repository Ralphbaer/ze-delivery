# Partner Service
[![CircleCI](https://circleci.com/gh/Ralphbaer/ze-delivery/tree/master.svg?style=svg&circle-token=0d6b9c7b5df0b9ef027d42b792782bf19393cf19)](https://circleci.com/gh/Ralphbaer/ze-delivery/tree/master)

This repo contains the source code of the Partner service.

## Requirements

| Name | Version | Notes |
|------|---------|---------|
| [golang](https://golang.org/dl/) | >= go1.15.14 | Main programming language
| [docker](https://www.docker.com/) | n/a | Used to start local environment providers (MongoDB)
| [terraform](https://www.terraform.io/) | n/a | Used to create AWS Fargate Service on AWS

## Providers

| Name | Version |
|------|---------|
| [aws](https://aws.amazon.com/pt/) | n/a |

## Installation


```sh
make setup
```

## Usage

# Start Local
```bash
make run           # Start service on port 3000 without test data (empty database)
make run-test      # Start service on port 3000 with test data (json](files/pdvs.json))
```

# Testing

```bash
make test          # Run all unit tests
```

# Deploy
The service uses a Circle pipeline that is triggered when a code is pushed to master branch.

The master has no PR approval criteria to facilitate the tests for the Ze Delivery team.


For testing purposes, a make command responsible for deploying the service in production is available below:

```bash
make deploy          # Using terraform, deploy the current code to production
```


## Documentation

Visit [this link](http://localhost:3000/partner/docs) for documentation.
 