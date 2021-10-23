--- change / add / modify this schema as necessary
CREATE TABLE public.requests (
    created_at timestamp without time zone DEFAULT now() NOT NULL,
    data jsonb
);

CREATE TABLE public.connections (
    id BIGSERIAL,
    created_at timestamp without time zone DEFAULT now() NOT NULL
);