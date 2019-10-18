#!/bin/sh

if [[ -z "$MOODBOARD_URL" ]]; then
  echo "==> Error: Please specify moodboard URL using ..."
  echo "export MOODBOARD_URL=moodboard.example.com"
  exit 1
fi

MOOD="$(node -e 'console.log(encodeURIComponent(process.argv.slice(1).join(" ")))' -- "$@")"
exec curl -XPUT "$MOODBOARD_URL/mood/$MOOD"
