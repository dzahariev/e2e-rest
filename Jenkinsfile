pipeline {
    // install golang 1.20.1 on Jenkins node
    agent any
    tools {
        go 'go1.20.1'
    }
    environment {
        CGO_ENABLED = 0 
        GOPATH = "${JENKINS_HOME}/jobs/${JOB_NAME}/builds/${BUILD_ID}"
    }
    stages {
        stage("format") {
            steps {
                echo 'Format step started ...'
                sh 'go fmt ./...'
            }
        }
        
        stage("build") {
            steps {
                echo 'Build step started ...'
                sh 'go build -o bin/e2e-rest'
            }
        }
        stage("clean") {
            steps {
                echo 'Clean step started ...'
                sh 'rm -fr bin'
            }
        }
    }
}
