CREATE TABLE public.requests (
    id bigint NOT NULL,
    method varchar(255),
    params jsonb,
    created_at timestamp without time zone DEFAULT now() NOT NULL
);