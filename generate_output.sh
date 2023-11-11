PROTO_ROOT="./proto/schemas"

# Generate Go Proto ---------------------------------
GO_OUTPUT_ROOT="./server/serverpb" # relative go proto path
echo "🧪 Generating Go bindings..."

mkdir -p $GO_OUTPUT_ROOT/output
rm -rf $GO_OUTPUT_ROOT/output/*

./proto/protoc --proto_path=$PROTO_ROOT --go_out=paths=source_relative:$GO_OUTPUT_ROOT $PROTO_ROOT/*.proto

rm -r server/serverpb/output

# Generate Typescript bindings ---------------------------------
TS_OUTPUT_ROOT="./client/clientpb"
echo "🧪 Generating Typescript bindings..."

mkdir -p $TS_OUTPUT_ROOT/output
rm -rf $TS_OUTPUT_ROOT/output/*

./proto/protoc \
    --plugin="./node_modules/.bin/protoc-gen-ts_proto" \
    --ts_proto_opt=esModuleInterop=true \
    --ts_proto_out="./client/clientpb" \
    ./proto/schemas/*.proto


./proto/protoc \
		--plugin="./node_modules/.bin/protoc-gen-js" \
		--js_out="import_style=commonjs,binary:./client/clientpb" \
		./proto/schemas/*.proto

find client/clientpb/proto/schemas -type f \( -name "*.ts" -o -name "*.js" \) -exec mv {} client/clientpb/ \;
rm -r client/clientpb/output client/clientpb/proto