FROM debian:bullseye AS build

RUN apt update && \
    apt upgrade -y

RUN apt install -y golang git

RUN git clone http://git.zackmarvel.com/zack/cloudflare-ddns.git /src && \
    cd /src && \
    go build && \
    GOBIN=/usr/local/bin go install


FROM debian:bullseye

COPY --from=build /usr/local/bin/cloudflare-ddns /usr/local/bin/cloudflare-ddns
