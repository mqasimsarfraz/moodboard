#!/bin/bash
set -e

# make sure docker is installed
which docker > /dev/null 2>&1
if [[ $? -ne 0 ]]; then
  echo "==> Error: Please install docker..."
  exit 1
fi

# build the image
docker rmi -f moodboard-dev || true
docker build --no-cache -t moodboard-dev .

# clean the old container and start
docker rm -f moodboard-dev || true
docker run --rm -d -p 80:3080 --name moodboard-dev moodboard-dev

echo "==> INFO: Moodboard running at http://localhost"
