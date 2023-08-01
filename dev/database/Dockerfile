from postgres:12

ENV POSTGRES_HOST_AUTH_METHOD trust
ENV POSTGRES_PASSWORD postgres
ENV POSTGRES_DB agora-vnf-manager

COPY schemas/01-agora.sql               /docker-entrypoint-initdb.d/01-agora.sql
COPY schemas/02-vnf-device-mapper.sql   /docker-entrypoint-initdb.d/02-vnf-device-mapper.sql
COPY schemas/03-vnf-infrastructure.sql  /docker-entrypoint-initdb.d/03-vnf-infrastructure.sql
COPY schemas/04-vnf-instance.sql        /docker-entrypoint-initdb.d/04-vnf-instance.sql

RUN chmod 777 /docker-entrypoint-initdb.d/01-agora.sql
RUN chmod 777 /docker-entrypoint-initdb.d/02-vnf-device-mapper.sql
RUN chmod 777 /docker-entrypoint-initdb.d/03-vnf-infrastructure.sql
RUN chmod 777 /docker-entrypoint-initdb.d/04-vnf-instance.sql
