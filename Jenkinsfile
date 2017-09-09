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
   withEnv(["GOPATH=${WORKSPACE}", "PATH+GO=${root}/bin", "GOBIN=${WORKSPACE}/bin"]){
      sh "go test -v -coverprofile=coverage.out -covermode count > tests.out"

      // convert tests result
      sh "go install github.com/tebeka/go2xunit"
      sh "'${WORKSPACE}/bin/go2xunit' < tests.out -output tests.xml"
      junit "tests.xml"

      // convert coverage
      sh "go install github.com/t-yuki/gocover-cobertura"
      sh "'${WORKSPACE}/bin/gocover-cobertura' < coverage.out > coverage.xml"
      step([$class: 'CoberturaPublisher', coberturaReportFile: 'coverage.xml'])

    }
  }
  stage ('Archive') {
    archiveArtifacts '**/tests.out, **/tests.xml, **/coverage.out, **/coverage.xml'
  }
}
