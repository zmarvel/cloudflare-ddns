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
                sh 'docker build -t zpmarvel/cloudflare-ddns .'
            }
        }

        stage('Push docker image to Docker Hub') {
            environment {
                DOCKERHUB_CREDENTIALS=credentials('b13e0e36-a886-4b25-b8a9-53cd8a2f9e58')
            }
            steps {
                sh 'echo $DOCKERHUB_CREDENTIALS_PSW | docker login -u $DOCKERHUB_CREDENTIALS_USR --password-stdin'
                sh 'docker push zpmarvel/cloudflare-ddns'
            }
        }
    }
}
