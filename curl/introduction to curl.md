# Introduction to cURL
![curl image](https://curl.se/logo/curl-logo.svg)

## Agenda
- What is cURL ?
- Simple example GET
- Simple example POST with JSON body
- Simple example authentication with header
- Simple example download resource

## What is cURL ?
> cURL is abbreviation for "Client URL". cURL is command line tool and library for transferring data with URLs.

## GET
checkout
```sh
git checkout 1881d984cfca7f303fe08799d017f4299f9db78a
```

curl
```sh
curl '127.0.0.1'
```

## POST with JSON body
checkout
```sh
git checkout dccf79d8c434f0fbd311b9f91c9dd9bc8fdc24b6
```
curl
```sh
curl -X POST 127.0.0.1:8000/greeting -H 'Content-Type: application/json' -d '{
"name": "nardiyansah"
}'
```
- -X, --request : define HTTP method
- -H, --header : define header (can more than one)
- -d : data

## Auth with header
checkout
```sh
git checkout 6e7543d567d59d586a21d196707671906def01ff
```
curl
```sh
curl 127.0.0.1:8000 -H 'X-key: secret'
```

## Download file
checkout
```sh
git checkout 8860f921ffc782ed118ca841c1d55e7ad856a098
```
curl
```sh
curl 127.0.0.1:8000/file -H 'X-key: secret' -o README.md
```
- -o : specified output file