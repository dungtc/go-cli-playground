# Create a simple cli (command line interface) tool from scratch in Golang

A cli interacts with computer program like go, docker, k8s, npm, vice versa.

Example:
```
$ go get -u URL
$ docker run -t image
$ kubectl get pods
$ npm install --save package
```

A common structure
```
App Subcommands --Flags Args
```

**App** has many subcommands
**Subcommands** has own flags and arguments

# Build
```
$ go install github.com/dungtc/go-cli-playground
```
It will install your binary into $PATH


# Usage

### Help
```
-help
	Usage:
		go-cli-playground [command]
	Available Commands:
		add         Add a new task
		count       Count total tasks
		list        List of tasks
```

### Add task command
```
$ go-cli-playground add Make a cookie
```


### List task command
```
$ go-cli-playground list
1. Do some thing
2. Make a cookie
```

### Count task command
```
$ go-cli-playground count
Count: 2
```