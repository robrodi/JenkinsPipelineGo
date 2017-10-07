#!/usr/bin/env groovy
// https://www.cloudbees.com/blog/top-10-best-practices-jenkins-pipeline-plugin
agent {
        docker { image 'golang:alpine' }
    }
node() {
  def root = tool name: 'Go 1.9', type: 'go'
    
  stage('Preparation') {
    checkout scm
  }
  stage ('Compile') {
    sh "${root}/bin/go build"
  }
  stage ('Static Analysis'){
    withEnv(["GOPATH=${WORKSPACE}", "PATH+GO=${root}/bin:${WORKSPACE}/bin", "GOBIN=${WORKSPACE}/bin"]){
      sh "go get github.com/golang/lint/golint"

      try{
        sh "golint ."
        sh "go vet ."
      } catch (err){
        sh "echo static analyis failed.  See report"
      }
      
      warnings canComputeNew: true, canResolveRelativePaths: true, categoriesPattern: '', consoleParsers: [[parserName: 'Go Vet'], [parserName: 'Go Lint']], defaultEncoding: '', excludePattern: '', healthy: '', includePattern: '', messagesPattern: '', unHealthy: ''
    }
  }
  stage ('Test') {
   withEnv(["GOPATH=${WORKSPACE}", "PATH+GO=${root}/bin:${WORKSPACE}/bin", "GOBIN=${WORKSPACE}/bin"]){
      sh "go test -v -coverprofile=coverage.out -covermode count > tests.out"

      // convert tests result
      sh "go get github.com/tebeka/go2xunit"
      sh "go2xunit < tests.out -output tests.xml"
      junit "tests.xml"

      // convert coverage
      sh "go get github.com/t-yuki/gocover-cobertura"
      sh "gocover-cobertura < coverage.out > coverage.xml"

      step([$class: 'CoberturaPublisher', coberturaReportFile: 'coverage.xml'])
    }
  }
  stage ('Archive') {
    archiveArtifacts '**/tests.out, **/tests.xml, **/coverage.out, **/coverage.xml, **/coverage2.xml'
  }
}
