FROM alpine:latest
RUN mkdir /app
WORKDIR /app
COPY consignment-service ./
COPY ./cmd/cmd ./
CMD [ "./consignment-service" ]
ENTRYPOINT [ "./consignment-service" ]