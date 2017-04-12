FROM alpine

ENV POSTGRES_HOST=localhost \
    POSTGRES_PORT=5432 \
    POSTGRES_USER=proxy \
    POSTGRES_PASSWORD=proxy \
    POSTGRES_DBNAME=proxy

COPY . /usr/local/bin

RUN chmod +x /usr/local/bin/*

EXPOSE 3000

ENTRYPOINT ["docker-entrypoint.sh"]
