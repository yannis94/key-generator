# key-generator
Generate a random key or a passphrase. You can customize them as you wish (number of words, letters, characters or digits). 
 
This is a Go training and learning project.

# Getting started
Go version 1.20 is required
## Install
```bash
git clone github.com/yannis94/key-generator
go mod install
```

## Build
```bash
go build

# use inside of directory
./key-generator

# global usage
mv ./key-generator /usr/local/go/bin/<name-you-want>
<name-you-want> -password -l 15 -n 4
```

## Flag
If you use no flag, the program pass in manual config (ask for choices to generate key or passphrase). But you can pass some flags to generate instantly.

### Password generation
* **-password** : Generate a password
* **-l** : Configure password's total letter
* **-n** : Configure password's total number
* **-c** : Configure password's total special character

### Passphrase generation
* **-phrase** : Generate a passphrase
* **-w** : Configure passphrase's number of word


## Examples
```bash
# Password generation
key-generator -password -l 10 -n 5 -c 2 # will generate something like > W2Dc2w"Oynp3u2%0Q

# Passphrase generation
key-generator -phrase -w 12 # will generate something like > cattle-stem-teeny-suspect-mind-spade-ocean-verdant-glass-curious-political-cub

# Manual config
key-generator # will ask you choices to generate either a passphrase or a password (with configs)
```
