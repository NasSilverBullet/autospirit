# autospirit

## cli tool to insert timestamps on TeamSpirit automatically

### Installing(requires Go 1.11.4)
```shell
$ brew install phantomjs
$ export GO111MODULE=auto
$ go get -u github.com/NasSilverBullet/autospirit/cmd/autospirit
```

### Setting(Please take care of security vulnerability!)
```
$ cd $GOPATH/src/github.com/NasSilverBullet/autospirit/config/
$ cp config.example.json config.json
$ vi config.json // please update config.json
```

### Usage
```shell
$ autospirit [command]
```

### Available Commands
```shell
bye         Leave the office
hello       Going to work
help        Help about any command
``` 

### Flags
```shell
 -h, --help   help for autospirit
```
