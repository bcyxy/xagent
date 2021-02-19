# 编译
echo "###### Build ###############################"
go build \
    -o xagent \
    -ldflags "
        -X 'main.gitCommitID=`git rev-parse HEAD`'
        -X 'main.buildTime=`date +"%Y-%m-%d %H:%M:%S"`'
        -X 'main.buildGoVersion=`go version`'
    " \
    src/main.go

# 打包
echo ""
echo "###### Package #############################"
rm -rf output
mkdir -p output/bin
mv xagent output/bin/
cp -f src/control.sh output/
cp -rf conf/ output/
