CREATE DATABASE pr
  WITH OWNER = pr
       ENCODING = 'UTF8'
       TABLESPACE = pg_default
       CONNECTION LIMIT = -1;

\connect pr

CREATE TABLE links
(
  hostname text NOT NULL,
  update_at timestamp with time zone NOT NULL DEFAULT now(),
  iterate boolean NOT NULL DEFAULT false,
  num bigint NOT NULL DEFAULT 0,
  CONSTRAINT links_pkey PRIMARY KEY (hostname)
)
WITH (
  OIDS=FALSE
);
ALTER TABLE links
  OWNER TO pr;

CREATE TABLE proxies
(
  hostname text NOT NULL,
  host text NOT NULL DEFAULT ''::text,
  port text NOT NULL DEFAULT ''::text,
  work boolean NOT NULL DEFAULT false,
  anon boolean NOT NULL DEFAULT false,
  checks integer NOT NULL DEFAULT 0,
  create_at timestamp with time zone NOT NULL DEFAULT now(),
  update_at timestamp with time zone NOT NULL DEFAULT now(),
  response bigint NOT NULL DEFAULT 0,
  id bigint NOT NULL DEFAULT nextval('proxies_id_seq'::regclass),
  scheme text NOT NULL DEFAULT 'http'::text,
  CONSTRAINT proxies_pkey PRIMARY KEY (id)
)
WITH (
  OIDS=FALSE
);
ALTER TABLE public.proxies
  OWNER TO pr;