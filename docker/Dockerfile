FROM alpine
ADD https://curl.haxx.se/ca/cacert.pem /etc/ssl/ca-bundle.pem
ADD https://github.com/iij/p2pubapi/releases/download/v0.1.0/p2pub-0.1.0.linux-386 /usr/bin/p2pub
RUN chmod +x /usr/bin/p2pub
