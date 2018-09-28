#!groovy
pipeline {
  environment {
    BUILD_SCRIPTS_GIT="https://github.com/Davitron/go-micro-services/"
    BUILD_HOME='/var/lib/jenkins/workspace'
    AWS_ACCESS_KEY_ID     = credentials('AWS_ACCESS_KEY')
    AWS_SECRET_ACCESS_KEY = credentials('AWS_SECRET_KEY')
  }

  agent any
  

  tools { go 'Go 1.11'}

  stages {
    stage('Checkout: Code') {
      steps {
        sh "mkdir -p $WORKSPACE/repo;\
        git config --global user.email 'segunmatthews@outlook.com';\
        git config --global user.name 'davitron';\
        git config --global push.default simple;\
        git clone $BUILD_SCRIPTS_GIT "
      }
    }

    stage("set kops store") {
      steps {
        sh "sudo kubectl create -R -f k8s/"
      }
    }
  }
  
  
  post {
    always {
      cleanWs()
    }
  }
}