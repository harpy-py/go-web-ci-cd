pipeline{
    agent none
    
    environment {
        IMAGE_TAG = "build-${env.BUILD_NUMBER}"
        DOCKER_USER = "USER"
    }
    
    stages{
        stage('Clone Repository'){
            agent{ label 'master' }
            steps {
                git 'https://github.com/harpy-py/go-web-ci-cd.git'
            }
        }

        stage('Login to Docker Hub') {
          steps {
            withCredentials([usernamePassword(credentialsId: 'dockerhub-creds', usernameVariable: 'DOCKER_USER', passwordVariable: 'DOCKER_PASS')]) {
              bat 'docker login -u %DOCKER_USER% -p %DOCKER_PASS%'
            }
        }
    }
        stage('Docker build image'){
            agent { label 'master' }
            steps {
                bat """
                set IMAGE_TAG=%IMAGE_TAG%
                set DOCKER_USER=%DOCKER_USER%
                docker-compose build
                """
            }
        }
        stage('List Docker images'){
            agent { label 'master' }
            steps{
                bat 'docker images'
            }
        }
        stage('Push Images to hub'){
            agent { label 'master' }
            steps{
                bat """
                docker push %DOCKER_USER%/go-web-app-web:%IMAGE_TAG%
                """
            }
        }
        stage('k8s Deployment'){
            agent { label 'master' }
            steps {
                script {
                    
                    def imageTag = "build-${BUILD_NUMBER}"

                powershell """
                    (Get-Content k8s\\kustomization.yml) -replace 'newTag: .*', 'newTag: ${imageTag}' | Set-Content k8s\\kustomization.yml
                """

                powershell 'kubectl apply -k k8s/'
        }
    }
        }
    }
}
