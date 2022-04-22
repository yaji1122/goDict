pipeline {
    environment {
      DOCKER = credentials('docker-hub')
    }
  agent any
  stages {
// Building your Test Images
    stage('BUILD') {

      steps {
            sh 'docker build -f Dockerfile \
            -t go-dictionary:trunk .'
          }

      post {
        failure {
            echo 'This build has failed. See logs for details.'
        }
      }
    }
// Performing Software Tests
//     stage('TEST') {
//       parallel {
//         stage('Mocha Tests') {
//           steps {
//             sh 'docker run --name nodeapp-dev --network="bridge" -d \
//             -p 9000:9000 nodeapp-dev:trunk'
//             sh 'docker run --name test-image -v $PWD:/JUnit --network="bridge" \
//             --link=nodeapp-dev -d -p 9001:9000 \
//             test-image:latest'
//           }
//         }
//         stage('Quality Tests') {
//           steps {
//             sh 'docker login --username $DOCKER_USR --password $DOCKER_PSW'
//             sh 'docker tag nodeapp-dev:trunk <DockerHub Username>/nodeapp-dev:latest'
//             sh 'docker push <DockerHub Username>/nodeapp-dev:latest'
//           }
//         }
//       }
//       post {
//         success {
//             echo 'Build succeeded.'
//         }
//         unstable {
//             echo 'This build returned an unstable status.'
//         }
//         failure {
//             echo 'This build has failed. See logs for details.'
//         }
//       }
//     }
// Deploying your Software
    stage('DEPLOY') {
          when {
           branch 'main'  //only run these steps on the master branch
          }
            steps {
                    retry(3) {
                        timeout(time:10, unit: 'MINUTES') {
                            sh 'docker tag go-dictionary:trunk yajilin/go-dictionary:latest'
                            sh 'docker push yajilin/go-dictionary:latest'
                            sh 'docker save yajilin/go-dictionary:latest | gzip > go-dictionary.tar.gz'
                        }
                    }
            }
            post {
                failure {
                    sh 'docker stop go-dictionary'
                    sh 'docker system prune -f'
                    deleteDir()
                }
            }
    }
// JUnit reports and artifacts saving
//     stage('REPORTS') {
//       steps {
//         junit 'reports.xml'
//         archiveArtifacts(artifacts: 'reports.xml', allowEmptyArchive: true)
//         archiveArtifacts(artifacts: 'nodeapp-prod-golden.tar.gz', allowEmptyArchive: true)
//       }
//     }
// Doing containers clean-up to avoid conflicts in future builds
//     stage('CLEAN-UP') {
//       steps {
//         sh 'docker stop go-dictionary'
//         sh 'docker system prune -f'
//         deleteDir()
//       }
//     }
  }
}