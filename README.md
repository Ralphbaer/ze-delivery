# Partner Service
[![CircleCI](https://circleci.com/gh/Ralphbaer/ze-delivery.svg?style=svg&circle-token=37cdd1fd1e89719b206adf5ae503bc83073b9c3a)](https://circleci.com/gh/Ralphbaer/ze-delivery/?branch=master)

The badge above shows the status of the last pipeline. If everything goes well, the badge will be "green" indicating success or "red" (failed) indicating some error in the build.

This repo contains the source code of the Partner service.

## Architecture

![alt text](./hexagonal-macro.png "Title")

## Requirements

| Name | Version | Notes | Mandatory
|------|---------|---------|---------|
| [golang](https://golang.org/dl/) | >= go1.15.14 | Main programming language | true
| [docker](https://www.docker.com/) | n/a | Used to start local environment providers (MongoDB) | true
| [aws-cli](https://aws.amazon.com/pt/cli/) | v2 | Used to create all AWS Enviroment (Just in case you want to know) | false
| [sh/bash] | depending on OS. Anyway, you should be able do execute any .sh file | Used to lint checks, test processes and some console interface customizations | true
| [make](https://www.gnu.org/software/make/) | depending on OS. Anyway, you should be able do execute make commands to run the project, tests and localenvironment | n/a | true

## Providers

| Name | Version | Notes
|------|---------|---------|
| [aws](https://aws.amazon.com/pt/) | n/a | All the infraestructure are on AWS
| [mongodb](https://www.mongodb.com/) | any stable version | If you want, you can use any mongoDB client to access the local database created | true

# Usage

### Start Local
Inside /partner, you can run:
```bash
make run           # Start service on port 3000 without test data (empty database)
make run-test      # Start service on port 3000 with test data (json](files/pdvs.json))
```

# Testing

```bash
make test                 # Run all unit tests and integration test
make unit-test            # Run all unit tests
make integration-test     # Run integration test
```
## Documentation

Visit [this link](http://ze-delivery-microservices-elb-6682139.us-east-2.elb.amazonaws.com/partner/docs#overview) for API documentation. If you want to access the docs locally, just change the host in the url to localhost:3000. Something like: http://localhost:3000/partner/docs

# Deploy
The service uses a Circle pipeline that is triggered when a code is pushed to master branch.
The master has no PR approval criteria to facilitate the tests for the Ze Delivery team.
For testing deploy purposes, you guys can change anything in the code or just change the signal.id file inside /partner, commit and push the change to master and wait the circle-ci to deploy the code. :)



 