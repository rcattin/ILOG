// french below

# Lezgo

lezgo is a command-line-interface written in Go during our ILOG project

## Prerequisites

- Go 

## Installation
1. `git clone https://github.com/rcattin/ILOG.git `
2. `cd lezgo/src`
3. `go build`
4. Make a symlink in /usr/bin (or /usr/local/bin) directory 
```bash
cd /usr/bin
sudo cp -s /opt/toolname/tool.sh /usr/bin/lezgo
```
5. Enjoy !

## Commands
### search
Finds every directory matching the given name in the working directory and prints their size

```bash
lezgo search -d dirName
```
This will recursively search for any directory named `dirName` in the working directory and sud-directories, then prints their size

```bash
lezgo search -d dirName -p pathName
```
This will recursively search for any directory named `dirName` in the `pathName` and its sud-directories, then prints their size

Other possible flags :
 - -v to print every directory searched
 - -h to print some help 

### replicates


## Class diagram

![class diagram](https://user-images.githubusercontent.com/81159061/156588750-e13982c1-e06a-4961-b5c8-361d25ff761f.png)


