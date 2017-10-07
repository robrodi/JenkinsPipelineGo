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


## Creating a Jobs DSL
* Jobs DSL plugin
* Authorize Project Plugin [ to prevent [script auth pain](https://github.com/jenkinsci/job-dsl-plugin/wiki/Script-Security)]
  * `Mange Jenkins` -> `Configure Global Security` -> `Access Control For Builds` -> `Run as User who triggered build`
* Groovy Plugin
* Create a new freeform job.
* SCM Git, Pointed at this repo
* Add Build Step -> Process Job DSLs.
* Look on Filesystem
* DSL Scripts [`JobsFile`](JobsFile)
* Save
* Click `Authorization` on the left nav of the project page
* Check `Configure Build Authorization` and Select `Run as User who Triggered Build`
* Run it
* see that a job called `PipelineFromDSL` was created, run that and hopefully it works and made a beautiful pipeline build!

## Docker
This doesn't work right if you're using my docker instracutions above.  
* Make sure docker is installed and running
* add \Program Files\Git\usr\bin to path so jenkins can find its binaries.
* 
