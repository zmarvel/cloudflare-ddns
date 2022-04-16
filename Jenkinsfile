pipeline {
    agent any
    stages {
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
