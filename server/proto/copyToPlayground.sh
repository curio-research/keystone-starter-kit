# used for porting over repos from one to another

# Define your source (folder A) and destination (folder B) paths
source_folder="../proto"
destination_folder="../../../keystone-internal-playground/proto"

# Check if the source folder exists
if [ ! -d "$source_folder" ]; then
    echo "Source folder does not exist: $source_folder"
    exit 1
fi

# Check if the destination folder exists; if not, create it
if [ ! -d "$destination_folder" ]; then
    mkdir -p "$destination_folder"
fi

# Delete all existing contents in folder B
rm -rf "$destination_folder"/*

# Copy the contents of folder A to folder B
cp -r "$source_folder"/* "$destination_folder"


echo "Contents of proto folder copied to playground! 🏖"

find $destination_folder -type f -exec sed -i '' 's/pb_base "github.com\/curio-research\/keystone\/game\/proto\/output\/pb.base"/pb_base "playground\/proto\/output\/pb.base"/g' {} \;
