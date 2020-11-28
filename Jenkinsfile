node {
    try{
        notifyBuild('STARTED')
        bitbucketStatusNotify(buildState: 'INPROGRESS')
        
        ws("${JENKINS_HOME}/jobs/${JOB_NAME}/builds/${BUILD_ID}/") {
            withEnv(["GOPATH=${JENKINS_HOME}/jobs/${JOB_NAME}/builds/${BUILD_ID}"]) {
                env.PATH="${GOPATH}/bin:$PATH"
                
                stage('Checkout'){
                    echo 'Checking out SCM'
                    checkout scm
                }
                
                stage('Pre Test'){
                    echo 'Pulling Dependencies'
            
                    sh 'go version'
                    sh 'go get -u github.com/golang/dep/cmd/dep'
                    sh 'go get -u github.com/golang/lint/golint'
                    sh 'go get github.com/tebeka/go2xunit'
                    
                    //or -update
                    sh 'cd ${GOPATH}/src/cmd/project/ && dep ensure' 
                }
        
                stage('Test'){
                    
                    //List all our project files with 'go list ./... | grep -v /vendor/ | grep -v github.com | grep -v golang.org'
                    //Push our project files relative to ./src
                    sh 'cd $GOPATH && go list ./... | grep -v /vendor/ | grep -v github.com | grep -v golang.org > projectPaths'
                    
                    //Print them with 'awk '$0="./src/"$0' projectPaths' in order to get full relative path to $GOPATH
                    def paths = sh returnStdout: true, script: """awk '\$0="./src/"\$0' projectPaths"""
                    
                    echo 'Vetting'

                    sh """cd $GOPATH && go tool vet ${paths}"""

                    echo 'Linting'
                    sh """cd $GOPATH && golint ${paths}"""
                    
                    echo 'Testing'
                    sh """cd $GOPATH && go test -race -cover ${paths}"""
                }
            
                stage('Build'){
                    echo 'Building Executable'
                
                    //Produced binary is $GOPATH/src/cmd/project/project
                    sh """cd $GOPATH/src/cmd/project/ && go build -ldflags '-s'"""
                }
                
                stage('BitBucket Publish'){
                
                    //Find out commit hash
                    sh 'git rev-parse HEAD > commit'
                    def commit = readFile('commit').trim()
                
                    //Find out current branch
                    sh 'git name-rev --name-only HEAD > GIT_BRANCH'
                    def branch = readFile('GIT_BRANCH').trim()
                    
                    //strip off repo-name/origin/ (optional)
                    branch = branch.substring(branch.lastIndexOf('/') + 1)
                
                    def archive = "${GOPATH}/project-${branch}-${commit}.tar.gz"

                    echo "Building Archive ${archive}"
                    
                    sh """tar -cvzf ${archive} $GOPATH/src/cmd/project/project"""

                    echo "Uploading ${archive} to BitBucket Downloads"
                    withCredentials([string(credentialsId: 'bb-upload-key', variable: 'KEY')]) { 
                        sh """curl -s -u 'user:${KEY}' -X POST 'downloads-page-url' --form files=@'${archive}' --fail"""
                    }
                }
            }
        }
    }catch (e) {
        // If there was an exception thrown, the build failed
        currentBuild.result = "FAILED"
        
        bitbucketStatusNotify(buildState: 'FAILED')
    } finally {
        // Success or failure, always send notifications
        notifyBuild(currentBuild.result)
        
        def bs = currentBuild.result ?: 'SUCCESSFUL'
        if(bs == 'SUCCESSFUL'){
            bitbucketStatusNotify(buildState: 'SUCCESSFUL')
        }
    }
}

def notifyBuild(String buildStatus = 'STARTED') {
  // build status of null means successful
  buildStatus =  buildStatus ?: 'SUCCESSFUL'

  // Default values
  def colorName = 'RED'
  def colorCode = '#FF0000'
  def subject = "${buildStatus}: Job '${env.JOB_NAME} [${env.BUILD_NUMBER}]'"
  def summary = "${subject} <${env.BUILD_URL}|Job URL> - <${env.BUILD_URL}/console|Console Output>"

  // Override default values based on build status
  if (buildStatus == 'STARTED') {
    color = 'YELLOW'
    colorCode = '#FFFF00'
  } else if (buildStatus == 'SUCCESSFUL') {
    color = 'GREEN'
    colorCode = '#00FF00'
  } else {
    color = 'RED'
    colorCode = '#FF0000'
  }

  // Send notifications
  slackSend (color: colorCode, message: summary)
}
