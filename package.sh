#!/usr/bin/env bash
 
# Package as xx.tar.gz or xx.zip
# Usage: ./package.sh -a amd64 -p linux -v v1.6.0
 
set -o errexit
set -o nounset
set -o pipefail

eval $(go env)

BINARY_NAME=''
MAIN_FILE=""

VERSION=''
GIT_COMMIT_ID=''

INPUT_OS=()
INPUT_ARCH=()
DEFAULT_OS=${GOHOSTOS}
DEFAULT_ARCH=${GOHOSTARCH}

SUPPORT_OS=(linux darwin windows)
SUPPORT_ARCH=(386 amd64)

LDFLAGS=''
INCLUDE_FILE=()
PACKAGE_DIR=''
BUILD_DIR=''

git_latest_tag() {
    local COMMIT_ID=""
    local TAG_NAME=""
    COMMIT_ID=`git rev-list --tags --max-count=1`
    TAG_NAME=`git describe --tags "${COMMIT_ID}"`
 
    echo ${TAG_NAME}
}

git_latest_commit() {
    echo "$(git rev-parse --short HEAD)"
}

print_message() {
    echo "$1"
}

print_message_and_exit() {
    if [[ -n $1 ]]; then
        print_message "$1"
    fi
    exit 1
}

set_os_arch() {
    if [[ ${#INPUT_OS[@]} = 0 ]];then
        INPUT_OS=("${DEFAULT_OS}")
    fi
 
    if [[ ${#INPUT_ARCH[@]} = 0 ]];then
        INPUT_ARCH=("${DEFAULT_ARCH}")
    fi
 
    for OS in "${INPUT_OS[@]}"; do
        if [[  ! "${SUPPORT_OS[*]}" =~ ${OS} ]]; then
            print_message_and_exit "不支持的系统${OS}"
        fi
    done
 
    for ARCH in "${INPUT_ARCH[@]}";do
        if [[ ! "${SUPPORT_ARCH[*]}" =~ ${ARCH} ]]; then
            print_message_and_exit "不支持的CPU架构${ARCH}"
        fi
    done
}

init() {
    set_os_arch
 
    if [[ -z "${VERSION}" ]];then
        VERSION=`git_latest_tag`
    fi
    GIT_COMMIT_ID=`git_latest_commit`
    LDFLAGS="-w -X 'main.AppVersion=${VERSION}' -X 'main.BuildDate=`date '+%Y-%m-%d %H:%M:%S'`' -X 'main.GitCommit=${GIT_COMMIT_ID}'"
 
    PACKAGE_DIR=${BINARY_NAME}-package
    BUILD_DIR=${BINARY_NAME}-build
 
    if [[ -d ${BUILD_DIR} ]];then
        rm -rf ${BUILD_DIR}
    fi
    if [[ -d ${PACKAGE_DIR} ]];then
        rm -rf ${PACKAGE_DIR}
    fi
 
    mkdir -p ${BUILD_DIR}
    mkdir -p ${PACKAGE_DIR}
}

build() {
    local FILENAME=''
    for OS in "${INPUT_OS[@]}";do
        for ARCH in "${INPUT_ARCH[@]}";do
            if [[ "${OS}" = "windows"  ]];then
                FILENAME=${BINARY_NAME}.exe
            else
                FILENAME=${BINARY_NAME}
            fi
            env CGO_ENABLED=0 GOOS=${OS} GOARCH=${ARCH} go build -ldflags "${LDFLAGS}" -o ${BUILD_DIR}/${BINARY_NAME}-${OS}-${ARCH}/${FILENAME} ${MAIN_FILE}
        done
    done
}

package_binary() {
    cd ${BUILD_DIR}
 
    for OS in "${INPUT_OS[@]}";do
        for ARCH in "${INPUT_ARCH[@]}";do
        package_file ${BINARY_NAME}-${OS}-${ARCH}
        if [[ "${OS}" = "windows" ]];then
            zip -rq ../${PACKAGE_DIR}/${BINARY_NAME}-${VERSION}-${OS}-${ARCH}.zip ${BINARY_NAME}-${OS}-${ARCH}
        else
            tar czf ../${PACKAGE_DIR}/${BINARY_NAME}-${VERSION}-${OS}-${ARCH}.tar.gz ${BINARY_NAME}-${OS}-${ARCH}
        fi
        done
    done
 
    cd ${OLDPWD}
}
 
# 打包文件
package_file() {
    if [[ "${#INCLUDE_FILE[@]}" = "0" ]];then
        return
    fi
    for item in "${INCLUDE_FILE[@]}"; do
            cp -r ../${item} $1
    done
}
 
# 清理
clean() {
    if [[ -d ${BUILD_DIR} ]];then
        rm -rf ${BUILD_DIR}
    fi
}
 
# 运行
run() {
    init
    build
    package_binary
    clean
}

package_gocron() {
    BINARY_NAME='gocron'
    MAIN_FILE="./cmd/gocron/gocron.go"
    INCLUDE_FILE=()


    run
}

package_gocron_node() {
    BINARY_NAME='gocron-node'
    MAIN_FILE="./cmd/node/node.go"
    INCLUDE_FILE=()

    run
}
 
# p 平台 linux darwin windows
# a 架构 386 amd64
# v 版本号  默认取git最新tag
while getopts "p:a:v:" OPT;
do
    case ${OPT} in
    p) IPS=',' read -r -a INPUT_OS <<< "${OPTARG}"
    ;;
    a) IPS=',' read -r -a INPUT_ARCH <<< "${OPTARG}"
    ;;
    v) VERSION=$OPTARG
    ;;
    *)
    ;;
    esac
done
 
package_gocron
package_gocron_node

