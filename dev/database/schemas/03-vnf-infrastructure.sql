CREATE SEQUENCE agorangmanager.vnf_infrastructure_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;

CREATE TABLE agorangmanager.vnf_infrastructure (
    id              int             NOT NULL,   
    name            varchar(255)    NOT NULL,
    description     varchar(255)    NOT NULL,
    config_file     varchar(255)    NOT NULL,
    CONSTRAINT vnf_infrastructure_pkey PRIMARY KEY (id)
);
