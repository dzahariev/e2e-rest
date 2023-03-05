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
        stage("build") {
            steps {
                echo 'Build step started ...'
                sh 'make build'
            }
        }
        stage("clean") {
            steps {
                echo 'Clean step started ...'
                sh 'make clean'
            }
        }
    }
}
