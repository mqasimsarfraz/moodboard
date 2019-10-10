#!/bin/bash

# make sure docker is installed
which docker > /dev/null 2>&1
if [[ $? -ne 0 ]]; then
  echo "==> Error: Please install docker..."
  exit 1
fi

# make sure GIPHY_API_KEY is set
if [[ -z $GIPHY_API_KEY ]]; then
  echo "==> Error: Please set 'GIPHY_API_KEY'"
  echo "==> Error: For example using:"
  echo "==> Error: export GIPHY_API_KEY=12345678"
  exit 1
fi

# build the image
docker rmi -f moodboard-dev || true
docker build --no-cache -t moodboard-dev .

# clean the old container and start
docker rm -f moodboard-dev || true
docker run --rm -d -p 80:3080 --name moodboard-dev -e GIPHY_API_KEY=${GIPHY_API_KEY} moodboard-dev

echo "==> INFO: Moodboard running at http://localhost"
