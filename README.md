# Moodboard
This tool will help show your team mood using a giphy GF.
## Installation
```
docker run --rm -d -e GIPHY_API_KEY=${API_KEY} -p 80:3080 --name moodboard smqasims/moodboard
```
if you run above on `example.com` then moodboard will be available at http://example.com 

## API
By default moodbaord will update after `30s` with a default mood of `hello` but you can update it using following:
```
curl -XPUT localhost/mood/<MOOD>
```
example:
```
curl -XPUT localhost/mood/happy
```