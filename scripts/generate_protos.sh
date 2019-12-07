#!/bin/bash

DIR=`pwd`
rm -rf $DIR/gen
LYFT_IMAGE="lyft/protocgenerator:d53ce1490e235bf765c93b4a8cfcdd07a1325024"

docker run -u $(id -u):$(id -g) -v $DIR:/defs $LYFT_IMAGE -i ./common/ -d ./common/ -l go --go_source_relative

languages=("python" "cpp" "java")

for lang in "${languages[@]}"
do
    docker run -u $(id -u):$(id -g) -v $DIR:/defs $LYFT_IMAGE -i ./common/proto -d ./common/proto -l $lang
done

# Docs generated
# docker run -u $(id -u):$(id -g) -e REPO_BLOB_SHA=master -e PROJECT_ANNOTATION_PREFIX=flyte.plugins.tfoperator -v $DIR:/defs $LYFT_IMAGE -i ./common/proto -d ./common/proto -l protodoc

# This section is used by Travis CI to ensure that the generation step was run
if [ -n "$DELTA_CHECK" ]; then
  DIRTY=$(git status --porcelain)
  if [ -n "$DIRTY" ]; then
    echo "FAILED: proto updated without commiting generated code."
    echo "Ensure make generate has run and all changes are committed."
    DIFF=$(git diff)
    echo "diff detected: $DIFF"
    DIFF=$(git diff --name-only)
    echo "files different: $DIFF"
    exit 1
  else
    echo "SUCCESS: Generated code is up to date."
  fi
fi
