FROM debian:latest

RUN apt-get update && apt-get install -y curl
RUN curl -sSf https://install.surrealdb.com | sh && which surreal

WORKDIR /scripts/

COPY setup.surql .

EXPOSE 8000

CMD ["sh", "-c", "surreal start file:///var/lib/surrealdb/data"]
# also run the following command:
# surreal import --namespace fds --database main --endpoint http://localhost:8000 /scripts/setup.surql
