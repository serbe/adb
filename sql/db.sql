CREATE DATABASE pr
  WITH OWNER = pr
       ENCODING = 'UTF8'
       TABLESPACE = pg_default
       CONNECTION LIMIT = -1;

\connect pr

CREATE TABLE proxies
(
  id bigserial,
  hostname text NOT NULL,
  scheme text NOT NULL DEFAULT 'http'::text,
  host text NOT NULL DEFAULT ''::text,
  port text NOT NULL DEFAULT ''::text,
  work boolean NOT NULL DEFAULT false,
  anon boolean NOT NULL DEFAULT false,
  response bigint NOT NULL DEFAULT 0,
  checks integer NOT NULL DEFAULT 0,
  create_at timestamp with time zone NOT NULL DEFAULT now(),
  update_at timestamp with time zone NOT NULL DEFAULT now(),
  CONSTRAINT proxies_pkey PRIMARY KEY (id)
)
WITH (
  OIDS=FALSE
);
ALTER TABLE public.proxies
  OWNER TO pr;