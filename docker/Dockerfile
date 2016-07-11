FROM alpine
ARG version=0.1.0
ADD https://curl.haxx.se/ca/cacert.pem /etc/ssl/ca-bundle.pem
ADD https://github.com/iij/p2pubapi/releases/download/v${version}/p2pub-${version}.linux-386 /usr/bin/p2pub
RUN chmod +x /usr/bin/p2pub
