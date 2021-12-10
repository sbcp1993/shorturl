CREATE SCHEMA gotool;

CREATE TABLE gotool.Shorturl
(
    id SERIAL,
    long_url character varying(200) NOT NULL,
    short_url bytea NOT NULL,
    CONSTRAINT shorturl_key UNIQUE (long_url, short_url)
);
