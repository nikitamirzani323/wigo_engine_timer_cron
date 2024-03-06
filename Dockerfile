FROM golang:alpine AS engine_timer

WORKDIR /appbuilds

COPY . .

RUN go mod tidy
RUN go build -o binary


FROM alpine:latest as agenconsumergeneratorrelease
WORKDIR /app
RUN apk add tzdata
COPY --from=engine_timer /appbuilds/binary .
COPY --from=engine_timer /appbuilds/.env /app/.env
ENV DB_USER="admindb"
ENV DB_PASS="asd123QWE"
ENV DB_HOST="128.199.124.131"
ENV DB_PORT="5432"
ENV DB_NAME="admindb"
ENV DB_SCHEMA="db_tot_wigo"
ENV DB_DRIVER="postgres"
ENV DB_REDIS_HOST="128.199.124.131"
ENV DB_REDIS_PORT="6379"
ENV DB_REDIS_PASSWORD="asdQWE123!@#"
ENV DB_REDIS_NAME="6"
ENV DB_CONF_COMPANY="NUKE"
ENV DB_CONF_CURR="IDR"
ENV TZ=Asia/Jakarta

RUN ln -snf /usr/share/zoneinfo/$TZ /etc/localtime && echo $TZ > /etc/timezone

ENTRYPOINT [ "./binary" ]