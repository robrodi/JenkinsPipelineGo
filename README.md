# JenkinsPipelineGo
Playing with Jenkins Pipelines &amp; Golang

## Jenkins Prereqs
* Install Jenkins' go plugin.  Manage Jenkins -> Global Tool Configuration -> Add Go -> Name = `Go 1.9`
* Warnings plugin
* Cobertura plugin

## Steps
* Use default scm
* default `go compile`
* `golin`t & `go vet`, both plumbed to [warnings](https://wiki.jenkins.io/display/JENKINS/Warnings+Plugin)
* Go Test w/ coverage [converted and plumbed](https://github.com/t-yuki/gocover-cobertura) to [Cobertura](https://wiki.jenkins.io/display/JENKINS/Cobertura+Plugin)
