# fcm-send

script for sending notification message using Firebase Cloud Messaging(FCM)

## Install

```
go get github.com/sckm/fcm-send
```

## Usage
create `data.json` file
``` json
{
    "message": "example message",
    "title": "example title"
}
```

run the following:
``` bash
$ token='YOUR DEVICE REGISTRATION TOKEN'
$ server_key='YOUR SERVER KEY'
$ fcm-send -t $token -s $server_key -p data.json
```