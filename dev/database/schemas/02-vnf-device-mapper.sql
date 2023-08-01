CREATE SEQUENCE agorangmanager.vnf_device_mapper_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;

CREATE TABLE agorangmanager.vnf_device_mapper (
    id                  int             NOT NULL,
    device_id           varchar(255)    NOT NULL,
    vnf_instance_id     int             NOT NULL,
    proxy_id            int             NOT NULL,
    CONSTRAINT vnf_device_mapper_pkey PRIMARY KEY (id)
);
