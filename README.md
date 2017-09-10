# JenkinsPipelineGo
Playing with Jenkins Pipelines &amp; Golang

## Steps
* Use default scm
* default `go compile`
* `golin`t & `go vet`, both plumbed to [warnings](https://wiki.jenkins.io/display/JENKINS/Warnings+Plugin)
* Go Test w/ coverage [converted and plumbed](https://github.com/t-yuki/gocover-cobertura) to [Cobertura](https://wiki.jenkins.io/display/JENKINS/Cobertura+Plugin)
