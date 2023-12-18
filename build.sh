idl_dir=${IDL_DIR:=../service-idl}
files=$(find $idl_dir -type f -name '*.proto')
if [ $? != 0 ];then
  exit 1
fi
for file in $files
do
  goctl rpc protoc $file --go_out=../ --go-grpc_out=../ --zrpc_out=../ --style=goZero
  if [ $? != 0 ];then
    exit 2
  fi
done