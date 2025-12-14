--
-- PostgreSQL database dump
--

-- Dumped from database version 17.4
-- Dumped by pg_dump version 17.4

SET statement_timeout = 0;
SET lock_timeout = 0;
SET idle_in_transaction_session_timeout = 0;
SET transaction_timeout = 0;
SET client_encoding = 'UTF8';
SET standard_conforming_strings = on;
SELECT pg_catalog.set_config('search_path', '', false);
SET check_function_bodies = false;
SET xmloption = content;
SET client_min_messages = warning;
SET row_security = off;

SET default_tablespace = '';

SET default_table_access_method = heap;

--
-- Name: categories; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.categories (
    id integer NOT NULL,
    name character varying(100) NOT NULL,
    description text,
    created_at timestamp without time zone DEFAULT now() NOT NULL,
    updated_at timestamp without time zone DEFAULT now() NOT NULL
);


ALTER TABLE public.categories OWNER TO postgres;

--
-- Name: categories_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.categories_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.categories_id_seq OWNER TO postgres;

--
-- Name: categories_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.categories_id_seq OWNED BY public.categories.id;


--
-- Name: items; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.items (
    id integer NOT NULL,
    category_id integer NOT NULL,
    name character varying(150) NOT NULL,
    price numeric(12,2) NOT NULL,
    purchase_date date NOT NULL,
    usage_days integer DEFAULT 0 NOT NULL,
    created_at timestamp without time zone DEFAULT now() NOT NULL,
    updated_at timestamp without time zone DEFAULT now() NOT NULL,
    CONSTRAINT items_price_check CHECK ((price > (0)::numeric)),
    CONSTRAINT items_usage_days_check CHECK ((usage_days >= 0))
);


ALTER TABLE public.items OWNER TO postgres;

--
-- Name: items_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.items_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.items_id_seq OWNER TO postgres;

--
-- Name: items_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.items_id_seq OWNED BY public.items.id;


--
-- Name: items_investment; Type: VIEW; Schema: public; Owner: postgres
--

CREATE VIEW public.items_investment AS
 SELECT id,
    name,
    price AS initial_price,
    EXTRACT(year FROM age(now(), (purchase_date)::timestamp with time zone)) AS years_used,
    (price * power(0.8, EXTRACT(year FROM age(now(), (purchase_date)::timestamp with time zone)))) AS current_value
   FROM public.items;


ALTER VIEW public.items_investment OWNER TO postgres;

--
-- Name: items_need_replace; Type: VIEW; Schema: public; Owner: postgres
--

CREATE VIEW public.items_need_replace AS
 SELECT id,
    category_id,
    name,
    price,
    purchase_date,
    usage_days,
    created_at,
    updated_at
   FROM public.items
  WHERE (usage_days > 100);


ALTER VIEW public.items_need_replace OWNER TO postgres;

--
-- Name: categories id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.categories ALTER COLUMN id SET DEFAULT nextval('public.categories_id_seq'::regclass);


--
-- Name: items id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.items ALTER COLUMN id SET DEFAULT nextval('public.items_id_seq'::regclass);


--
-- Data for Name: categories; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.categories (id, name, description, created_at, updated_at) FROM stdin;
1	Elektronik	Perangkat elektronik kantor	2025-12-12 01:35:27.178169	2025-12-12 01:35:27.178169
2	Furniture	Perabot kantor seperti kursi dan meja	2025-12-12 01:35:27.178169	2025-12-12 01:35:27.178169
3	ATK	Alat tulis kantor	2025-12-12 01:35:27.178169	2025-12-12 01:35:27.178169
4	Kebersihan	Peralatan kebersihan gedung	2025-12-12 01:35:27.178169	2025-12-12 01:35:27.178169
5	Testing	Testing description	2025-12-12 01:36:55.730516	2025-12-12 20:18:16.521866
7	test	desc	2025-12-13 14:13:24.171923	2025-12-13 14:13:24.171923
\.


--
-- Data for Name: items; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.items (id, category_id, name, price, purchase_date, usage_days, created_at, updated_at) FROM stdin;
1	1	Laptop Lenovo ThinkPad	12000000.00	2024-01-10	120	2025-12-12 01:35:27.178169	2025-12-12 01:35:27.178169
2	1	Monitor Dell 24 inch	3500000.00	2024-03-01	80	2025-12-12 01:35:27.178169	2025-12-12 01:35:27.178169
3	2	Kursi Ergonomis	1500000.00	2023-12-10	200	2025-12-12 01:35:27.178169	2025-12-12 01:35:27.178169
4	3	Stapler Besar	75000.00	2024-05-12	40	2025-12-12 01:35:27.178169	2025-12-12 01:35:27.178169
5	4	Vacuum Cleaner	950000.00	2023-11-15	180	2025-12-12 01:35:27.178169	2025-12-12 01:35:27.178169
\.


--
-- Name: categories_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.categories_id_seq', 7, true);


--
-- Name: items_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.items_id_seq', 5, true);


--
-- Name: categories categories_name_key; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.categories
    ADD CONSTRAINT categories_name_key UNIQUE (name);


--
-- Name: categories categories_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.categories
    ADD CONSTRAINT categories_pkey PRIMARY KEY (id);


--
-- Name: items items_name_key; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.items
    ADD CONSTRAINT items_name_key UNIQUE (name);


--
-- Name: items items_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.items
    ADD CONSTRAINT items_pkey PRIMARY KEY (id);


--
-- Name: idx_items_name; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_items_name ON public.items USING gin (to_tsvector('simple'::regconfig, (name)::text));


--
-- Name: items items_category_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.items
    ADD CONSTRAINT items_category_id_fkey FOREIGN KEY (category_id) REFERENCES public.categories(id) ON DELETE CASCADE;


--
-- PostgreSQL database dump complete
--

