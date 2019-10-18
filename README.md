# Moodboard
This tool will help show your team mood using a giphy GIF.
## Installation
```
docker pull smqasims/moodboard
docker run --rm -d -p 80:3080 --name moodboard smqasims/moodboard
```
if you run above on `example.com` then moodboard will be available at http://example.com

## API
The moodboard starts with a  default mood of `hello world` but you can update it using following:
```
curl -XPUT localhost/mood/<MOOD>
```
example:
```
curl -XPUT localhost/mood/happy
```
or use the [client](https://github.com/MQasimSarfraz/moodboard/blob/master/mood.sh).

### Rate limiting
By default moodboard will use giphy public beta key to talk to giphy API which is rate limited to `1000` requests per day. To avoid hitting rate limit you can get your [personal token](https://support.giphy.com/hc/en-us/articles/360020283431-Request-A-GIPHY-API-Key) and setting it using `GIPHY_API_KEY` env variable.
