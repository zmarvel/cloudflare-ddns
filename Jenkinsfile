pipeline {
    agent any
    stages {
        stage('Git checkout') {
            steps {
                checkout([
                    $class: 'GitSCM',
                    branches: [[name: '*/main']],
                    extensions: [[$class: 'CleanBeforeCheckout']],
                    userRemoteConfigs: [[url: 'http://git.zackmarvel.com/zack/cloudflare-ddns.git']]])
            }
        }
        stage('Build Docker image') {
            steps {
                sh 'docker build -t cloudflare-ddns .'
            }
        }
        stage('Push docker image to Docker Hub') {
            steps {
                sh 'docker push cloudflare-ddns'
            }
        }
    }
}
