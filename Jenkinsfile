pipeline {
    agent any
    stages {
        stage('Build') {
            steps {
                sh 'docker build -t go-devops-app ./app'
            }
        }
        stage('Test') {
            steps {
                sh 'curl http://localhost:5000/health'
            }
        }
        stage('Deploy') {
            steps {
                sh 'docker-compose up -d'
            }
        }
    }
}
