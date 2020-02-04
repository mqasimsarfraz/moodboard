<p align="left">
    <a href="https://hub.docker.com/r/smqasims/moodboard" alt="Pulls">
        <img src="https://img.shields.io/docker/pulls/smqasims/moodboard" /></a>
    <a href="https://mqasimsarfraz.github.io/" alt="Maintained">
        <img src="https://img.shields.io/maintenance/yes/2020.svg" /></a>

</p>

# Moodboard
A fun board tool to show your team mood using a giphy GIF.

[![Moodboard Demo](https://media.giphy.com/media/XDiQzh0JawsC0EDxiE/giphy.gif)](https://media.giphy.com/media/XDiQzh0JawsC0EDxiE/giphy.gif)

We use it at our workplace and we love it.

## Installation
```
docker run --rm -d -p 80:3080 --name moodboard smqasims/moodboard:v1.0.2
```
if you run above on `example.com` then moodboard will be available at http://example.com. 

You can also try out the platform specific binaries (e.g Raspberry Pi) in the [release](https://github.com/MQasimSarfraz/moodboard/releases) section of the project which start with a default port `3080` 

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

## Form
You can also use a simple form to update the mood by browsing to the path `/mood/form`.

[![Moodboard Form Demo](https://raw.githubusercontent.com/MQasimSarfraz/moodboard/master/images/form.png)](https://raw.githubusercontent.com/MQasimSarfraz/moodboard/master/images/form.png)

### Rate limiting
By default moodboard will use giphy public beta key to talk to giphy API which is rate limited to `1000` requests per day. To avoid hitting rate limit you can get your [personal token](https://support.giphy.com/hc/en-us/articles/360020283431-Request-A-GIPHY-API-Key) and setting it using `GIPHY_API_KEY` environment variable while launching your container. 