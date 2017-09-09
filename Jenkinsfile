#!/usr/bin/env groovy

node() {
  def root = tool name: 'Go 1.9', type: 'go'

  stage('Preparation') {
    checkout scm
  }
  stage ('Compile') {
    sh "${root}/bin/go build"
  }
  stage ('Test') {
   withEnv(["GOROOT=${root}", "GOPATH=${WORKSPACE}", "PATH+GO=${root}/bin", "GOBIN=${WORKSPACE}/bin"]){
      sh "go test -v > tests.out"
      sh "go install github.com/tebeka/go2xunit"
      sh "cat tests.out | '${WORKSPACE}/bin/go2xunit' -output tests.xml"
      junit "tests.xml"
    }
  }
  stage ('Archive') {
    archiveArtifacts '**/tests.out'
    archiveArtifacts '**/tests.xml'
  }
}
