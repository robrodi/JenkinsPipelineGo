#!/usr/bin/env groovy

node() {
  def root = tool name: 'Go 1.9', type: 'go'
  environment {
       GOROOT = "${root}"
       PATH = "${root}/bin:$PATH"
       GOPATH = "${WORKSPACE}"
   }
  stage('Preparation') {
    checkout scm
  }
  stage ('Compile')
  {
    sh "echo $path"
    sh "${root}/bin/go build"
  }
}
