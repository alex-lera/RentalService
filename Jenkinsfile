pipeline {
  agent any
  stages {
    stage('Get git') {
      steps {
        echo 'Getting git'
        git 'https://github.com/alex-lera/RentalService'
      }
    }

    stage('Get dependencies') {
      steps {
        echo 'Installing dependencies'
        sh 'go get "github.com/gorilla/mux"'
        sh 'go get "github.com/go-sql-driver/mysql"'
        sh 'go get "gopkg.in/yaml.v2"'
      }
    }

    stage('Build-go') {
      steps {
        echo 'Compiling and building'
        sh 'go build rentalcar.go'
      }
    }

    stage('Build') {
      steps {
        sh 'docker build -t servicerental:latest .'
        sh 'docker tag servicerental:latest 192.168.171.135:5000/servicerental:latest'
        sh 'docker push 192.168.171.135:5000/servicerental:latest'
      }
    }

  }
  tools {
    dockerTool 'docker'
    go 'go'
  }
}