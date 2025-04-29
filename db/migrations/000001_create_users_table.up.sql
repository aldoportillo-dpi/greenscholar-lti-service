CREATE TABLE public."USERS"
(
    "ID" character varying(30),
    PRIMARY KEY ("ID")
);

ALTER TABLE IF EXISTS public."USERS"
    OWNER to dpi;