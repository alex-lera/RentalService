pipeline {
  agent any
  stages {
    stage('Get git') {
      steps {
        echo 'Getting git'
        git 'https://github.com/alex-lera/RentalService'
      }
    }

    stage('Pre Test') {
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
        sh 'docker build -t servicerental:1.2 . | grep built | cut -d " " -f3 > test.txt'
        sh 'docker tag $(cat test.txt) 192.168.171.135:5000/servicerental:1.2'
        sh 'docker push 192.168.171.135:5000/servicerental:1.2'
      }
    }

  }
  tools {
    dockerTool 'docker'
    go 'go'
  }
}