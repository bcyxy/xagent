# 编译
echo "###### Build ###############################"
go build \
-ldflags "
    -X 'main.gitCommitID=`git rev-parse HEAD`'
    -X 'main.buildTime=`date +"%Y-%m-%d %H:%M:%S"`'
    -X 'main.buildGoVersion=`go version`'
"

# 打包
echo ""
echo "###### Package #############################"
rm -rf output
mkdir -p output/bin
mv xagent output/bin
cp -rf conf/ output/
