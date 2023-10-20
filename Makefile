.PHONY: pack

NOW = $(shell date -u '+%Y%m%d%I%M%S')

APP?=xfd-backent
TARGET_DIR=./output
SERVER_BIN = ${TARGET_DIR}/${APP}
COMMIT_ID=`git rev-parse --short HEAD`
#STAGE?=test
#STAGE_PROD=prod
#STAGE_TEST=test
OUTPUT_ZIP?=${APP}.tar.gz

prepare:
	mkdir -p ${TARGET_DIR}
	mkdir -p ${TARGET_DIR}/logs
	mkdir -p ${TARGET_DIR}/config
	mkdir -p ${TARGET_DIR}/file_nft
	echo "-------------------version:${COMMIT_ID}---------------------------\n"
	echo "-------------------version:${STAGE}---------------------------\n"
	echo ${COMMIT_ID} > ${TARGET_DIR}/version
	cp start.sh ${TARGET_DIR}
	cp supervisor.conf ${TARGET_DIR}
	cp -r stages ${TARGET_DIR}
	cp -r resource ${TARGET_DIR}
#ifeq ($(STAGE), $(STAGE_PROD))
#	cp stages/as-prod/config.toml ${TARGET_DIR}/config
#else
#	cp stages/as-test/config.toml ${TARGET_DIR}/config
#endif

build: prepare
	export GO111MODULE=on && export GOPROXY=https://goproxy.cn && go mod tidy && go build -ldflags "-w -s" -o $(SERVER_BIN) main.go

pack: clean build
	tar czvf ${OUTPUT_ZIP} -C output . && mv ${OUTPUT_ZIP} ${TARGET_DIR} 

clean:
	if [ -d "${TARGET_DIR}" ]; then rm -rf ${TARGET_DIR}; fi
