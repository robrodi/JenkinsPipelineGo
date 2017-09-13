# JenkinsPipelineGo
Playing with Jenkins Pipelines &amp; Golang

## Jenkins Prereqs
* `docker pull jenkins`
*  `docker run -d -p [your port here]:8080 -v [your host path here]:/var/jenkins_home:z -t jenkins`
* Install Jenkins' go plugin.  Manage Jenkins -> Global Tool Configuration -> Add Go -> Name = `Go 1.9`
* Warnings plugin
* Cobertura plugin

## Setting up the job
* Create a new pipeline job
* Use github to this repo.
* Pipeline -> Pipeline script from SCM
* SCM Git, pointed at this repo!
* Script file should be `Jenkinsfile`

## Pipeline Steps
* Use default scm
* default `go compile`
* `golint` & `go vet`, both plumbed to [warnings](https://wiki.jenkins.io/display/JENKINS/Warnings+Plugin)
* Go Test w/ coverage [converted and plumbed](https://github.com/t-yuki/gocover-cobertura) to [Cobertura](https://wiki.jenkins.io/display/JENKINS/Cobertura+Plugin)
