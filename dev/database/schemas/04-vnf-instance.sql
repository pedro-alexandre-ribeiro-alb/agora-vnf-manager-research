CREATE SEQUENCE agorangmanager.vnf_instance_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;

CREATE TABLE agorangmanager.vnf_instance (
    id                      int             NOT NULL,
    name                    varchar(255)    NOT NULL,
    description             varchar(255)    NOT NULL,
    type                    varchar(255)    NOT NULL,
    vnf_infra_id            int             NOT NULL,
    discovered              boolean         NOT NULL,
    management_interface    varchar(255)    NOT NULL,
    control_interface       varchar(255)    NOT NULL,
    vendor                  varchar(255)    NOT NULL,
    version                 varchar(255)    NOT NULL,
    CONSTRAINT vnf_instance_pkey PRIMARY KEY (id),
    CONSTRAINT vnf_instance_fkey1 FOREIGN KEY (vnf_infra_id) REFERENCES agorangmanager.vnf_infrastructure(id) 
);
