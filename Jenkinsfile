
pipeline {
    agent any

    stages {
        stage('pull code') {
            steps {
               git branch: 'main', url: 'git://github.com/czc-beijing/goInstance.git'
            }
        }
        stage('build') {
            steps {
               sh '''echo "开始构建"
               echo "结束构建"'''
            }
        }
        stage('repldy') {
            steps {
                sshPublisher(publishers: [sshPublisherDesc(configName: 'work', transfers: [sshTransfer(cleanRemote: false, excludes: '', execCommand: 'echo "success"', execTimeout: 120000, flatten: false, makeEmptyDirs: false, noDefaultExcludes: false, patternSeparator: '[, ]+', remoteDirectory: '/goInstance', remoteDirectorySDF: false, removePrefix: '', sourceFiles: '**')], usePromotionTimestamp: false, useWorkspaceInPromotion: false, verbose: false)])
            }
        }
    }
}

