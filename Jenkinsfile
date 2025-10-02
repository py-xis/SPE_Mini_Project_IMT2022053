pipeline {
  agent any

  environment {
    REGISTRY = 'docker.io'
    APP_NAME = 'scicalc'
    DOCKERHUB_USER = 'pyxis2004' // Redeclared so that it accessible in the email also
    // macOS-friendly PATH: Homebrew Go + Docker Desktop CLI
    PATH = "/opt/homebrew/bin:/usr/local/go/bin:/usr/local/bin:/Applications/Docker.app/Contents/Resources/bin:/usr/bin:/bin:${env.PATH}"
  }

  options {
    disableConcurrentBuilds()
    buildDiscarder(logRotator(numToKeepStr: '20'))
    timestamps()
  }

  triggers {
    githubPush()
  }

  stages {

    stage('Checkout') {
      steps {
        checkout scm
        script {
          env.GIT_SHORT = sh(returnStdout: true, script: 'git rev-parse --short HEAD').trim()
        }
      }
    }

    stage('Run tests (Go)') {
      steps {
        sh '''
          which go || true
          go version
          go mod download
          go test ./... -v
        '''
      }
    }

    stage('Build image (local tag)') {
      steps {
        sh '''
          which docker || true
          docker --version
          docker build -t ${APP_NAME}:local-${GIT_SHORT} .
        '''
      }
    }

    stage('Login & Push image') {
      steps {
        withCredentials([usernamePassword(credentialsId: 'dockerhub-creds',
                                          usernameVariable: 'DOCKERHUB_USER',
                                          passwordVariable: 'DOCKERHUB_PASS')]) {
          sh '''
            echo "$DOCKERHUB_PASS" | docker login -u "$DOCKERHUB_USER" --password-stdin

            docker tag ${APP_NAME}:local-${GIT_SHORT} ${REGISTRY}/${DOCKERHUB_USER}/${APP_NAME}:latest
            docker tag ${APP_NAME}:local-${GIT_SHORT} ${REGISTRY}/${DOCKERHUB_USER}/${APP_NAME}:${GIT_SHORT}

            docker push ${REGISTRY}/${DOCKERHUB_USER}/${APP_NAME}:latest
            docker push ${REGISTRY}/${DOCKERHUB_USER}/${APP_NAME}:${GIT_SHORT}

            docker logout
          '''
        }
      }
    }

    stage('Deploy locally (Ansible)') {
      steps {
        withCredentials([usernamePassword(credentialsId: 'dockerhub-creds',
                                          usernameVariable: 'DOCKERHUB_USER',
                                          passwordVariable: 'DOCKERHUB_PASS')]) {
          sh '''
            ansible-galaxy collection install community.docker --force >/dev/null 2>&1 || true

            cd ansible
            ansible-playbook -i inventory.ini deploy.yml \
              -e registry="${REGISTRY}" \
              -e dockerhub_username="${DOCKERHUB_USER}" \
              -e app_name="${APP_NAME}" \
              -e deploy_tag="${GIT_SHORT}"
          '''
        }
      }
    }
  }

  post {
    success {
      emailext(
        to: 'Pranav.Kulkarni@iiitb.ac.in',
        subject: "SUCCESS: ${env.JOB_NAME} #${env.BUILD_NUMBER} (${env.GIT_SHORT})",
        mimeType: 'text/html',
        attachLog: true,
        body: """
          <h3>Build Succeeded</h3>
          <table border="0" cellpadding="4">
            <tr><td><b>Job</b></td><td>${env.JOB_NAME}</td></tr>
            <tr><td><b>Build</b></td><td>#${env.BUILD_NUMBER}</td></tr>
            <tr><td><b>Commit</b></td><td>${env.GIT_SHORT}</td></tr>
            <tr><td><b>Image</b></td><td>${env.REGISTRY}/${env.DOCKERHUB_USER}/${env.APP_NAME}:${env.GIT_SHORT}</td></tr>
          </table>
          <p>Console log: <a href="${env.BUILD_URL}console">${env.BUILD_URL}console</a></p>
        """
      )
    }
    failure {
      emailext(
        to: 'Pranav.Kulkarni@iiitb.ac.in',
        subject: "FAILED: ${env.JOB_NAME} #${env.BUILD_NUMBER}",
        mimeType: 'text/html',
        attachLog: true,
        body: """
          <h3>Build Failed</h3>
          <p>Check the console log:</p>
          <p><a href="${env.BUILD_URL}console">${env.BUILD_URL}console</a></p>
        """
      )
    }
  }
}