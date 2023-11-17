PROTO_ROOT="./schemas"

# Generate Go Proto ---------------------------------
GO_OUTPUT_ROOT="../server/serverpb" # relative go proto path
echo "🧪 Generating Go bindings..."

rm -rf $GO_OUTPUT_ROOT/*
mkdir -p $GO_OUTPUT_ROOT

./protoc --proto_path=$PROTO_ROOT --go_out=paths=source_relative:$GO_OUTPUT_ROOT $PROTO_ROOT/*.proto

# # Generate Typescript bindings ---------------------------------
TS_OUTPUT_ROOT="../client/src/clientpb"
echo "🧪 Generating Typescript bindings..."

rm -rf $TS_OUTPUT_ROOT/*
mkdir -p $TS_OUTPUT_ROOT

./protoc \
    --plugin="./node_modules/.bin/protoc-gen-ts_proto" \
    --ts_proto_opt=esModuleInterop=true \
    --ts_proto_out="../client/src/clientpb" \
    ./schemas/*.proto


./protoc \
		--plugin="./node_modules/.bin/protoc-gen-js" \
		--js_out="import_style=commonjs,binary:../client/src/clientpb" \
		./schemas/*.proto