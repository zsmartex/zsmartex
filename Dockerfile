FROM golang:1.21-alpine AS builder

WORKDIR /build

# RUN apk update
# RUN apk add --no-cache git

# ARG GITHUB_TOKEN

ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOARCH="amd64" \
    GOOS=linux

COPY go.mod go.sum ./

# RUN git config --global url."https://${GITHUB_TOKEN}:x-oauth-basic@github.com/".insteadOf "https://github.com/"

RUN go mod download

COPY . .
RUN go build -o user-server ./cmd/user/server
RUN go build -o user-client ./cmd/user/client

# Use alpine for production mode
FROM alpine

# ARG KAIGARA_VERSION=v1.0.34

# RUN apk add ca-certificates curl
WORKDIR /app

ENV APP_HOME=/app

# ARG MAXMINDDB_LINK
# Open Source license key provided by Openware has some download rate and amount limits
# We strongly suggest you to create your oun key and pass via --build-arg MAXMINDDB_LICENSE_KEY
# All the guidance on how to create license key you can find here - https://blog.maxmind.com/2019/12/18/significant-changes-to-accessing-and-using-geolite2-databases/
# ARG MAXMINDDB_LICENSE_KEY=T6ElPBlyOOuCyjzw
# ENV MAXMINDDB_LINK=${MAXMINDDB_LINK:-https://download.maxmind.com/app/geoip_download?edition_id=GeoLite2-Country&suffix=tar.gz&license_key=${MAXMINDDB_LICENSE_KEY}}

# Download MaxMind Country DB
# RUN wget -O ${APP_HOME}/geolite.tar.gz ${MAXMINDDB_LINK} \
#     && mkdir -p ${APP_HOME}/geolite \
#     && tar xzf ${APP_HOME}/geolite.tar.gz -C ${APP_HOME}/geolite --strip-components 1 \
#     && rm ${APP_HOME}/geolite.tar.gz
# ENV MAXMINDDB_PATH=${APP_HOME}/geolite/GeoLite2-Country.mmdb

# RUN mkdir -p ${APP_HOME}/config

# Download list of Cloudflare IP Ranges (v4 and v6)
# RUN curl https://www.cloudflare.com/ips-v4 >> ${APP_HOME}/config/cloudflare_ips.yml \
#     && echo >> ${APP_HOME}/config/cloudflare_ips.yml \
#     && curl https://www.cloudflare.com/ips-v6 >> ${APP_HOME}/config/cloudflare_ips.yml \
#     && echo >> ${APP_HOME}/config/cloudflare_ips.yml

# RUN curl -Lo /usr/bin/kaigara https://github.com/openware/kaigara/releases/download/${KAIGARA_VERSION}/kaigara \
#   && chmod +x /usr/bin/kaigara

COPY --from=builder /build/user-server ./
COPY --from=builder /build/user-client ./
