#!/usr/bin/env groovy

pipeline {
  agent any
  tools { go '1.20.5' }
  environment {
      GO111MODULE = 'on'
      CGO_ENABLED = 0 
      GOPATH = "${JENKINS_HOME}/jobs/${JOB_NAME}/builds/${BUILD_ID}"
  }

  stages {
    stage('Build') {
        steps {
          // Output will be something like "go version go1.19 darwin/arm64"
          sh 'go version'
          sh 'go get -u golang.org/x/lint/golint'
          sh 'go mod download'
        }
    }
    
    stage('Test') {
        steps {
            sh 'go vet .'
            echo 'Running linting'
            sh 'golint .'
            echo 'Running test'
            sh 'go test -v'
        }
    }

    stage("Deploy back end") {
        steps {
            echo 'Compiling gateway'
            sh 'go build -o gateway cmd/gateway/main.go'
            
            echo 'Compiling auth_service'
            sh 'go build -o auth_service cmd/auth_service/main.go'
            
            echo 'Compiling user_service'
            sh 'go build -o user_service cmd/user_service/main.go'
            
            echo 'Compiling consumer_service'
            sh 'go build -o consumer_service cmd/consumer_service/main.go'
            
        }
    }
  }
}