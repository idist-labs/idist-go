#Stage 1: Build source
FROM golang:1.20.1 as build
WORKDIR /app
COPY ./ /app/
RUN go mod tidy && go mod vendor
RUN /app/build.sh

#Stage 2: Dockerize production
FROM alpine:3.13.2
RUN apk add tzdata && cp /usr/share/zoneinfo/Asia/Ho_Chi_Minh /etc/localtime && echo "Asia/Ho_Chi_Minh" >  /etc/timezone
COPY --from=build /app/bin/api-linux-386 /app/api-linux-386
COPY configs/production /app/configs/production
EXPOSE 3000
WORKDIR /app
CMD ["/app/api-linux-386", "-e", "production"]