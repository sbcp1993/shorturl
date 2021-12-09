CREATE SCHEMA gotool;

CREATE TABLE gotool.Shorturl
(
    id integer SERIAL,
    long_url character varying(100) NOT NULL,
    short_url character varying(8) NOT NULL,
    CONSTRAINT shorturl_key UNIQUE (long_url, short_url)
);
