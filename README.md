# ImgPrettify

## Installation
```
docker build . -t prettify

docker run --rm -d -p 80:8080 --name=imgPrettify prettify
```

## Using

```
curl --location --request POST 'http://localhost/v1/image/pretty' --form 'file=@"<in_file>"' --output <out_file>

```

or using form at http://localhost