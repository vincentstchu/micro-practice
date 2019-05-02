FROM alpine:latest
RUN mkdir /app
WORKDIR /app
COPY user-service ./
COPY user-cli/user-cli ./
CMD ["./user-service"]