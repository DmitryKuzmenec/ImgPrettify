# ImgPrettify

## Installation
```
docker build . -t imgPrettify

docker run -rm -d -p 80:8080 --name=imgPrettify imgPrettify
```

## Using

```
curl --location --request POST 'http://localhost/v1/image/pretty' --form 'file=@"<file_path>"'

```

or using form at http://localhost