--
-- PostgreSQL database dump
--

-- Dumped from database version 10.6 (Ubuntu 10.6-0ubuntu0.18.04.1)
-- Dumped by pg_dump version 10.6 (Ubuntu 10.6-0ubuntu0.18.04.1)

SET statement_timeout = 0;
SET lock_timeout = 0;
SET idle_in_transaction_session_timeout = 0;
SET client_encoding = 'UTF8';
SET standard_conforming_strings = on;
SELECT pg_catalog.set_config('search_path', '', false);
SET check_function_bodies = false;
SET client_min_messages = warning;
SET row_security = off;

--
-- Name: plpgsql; Type: EXTENSION; Schema: -; Owner: 
--

CREATE EXTENSION IF NOT EXISTS plpgsql WITH SCHEMA pg_catalog;


--
-- Name: EXTENSION plpgsql; Type: COMMENT; Schema: -; Owner: 
--

COMMENT ON EXTENSION plpgsql IS 'PL/pgSQL procedural language';


SET default_tablespace = '';

SET default_with_oids = false;

--
-- Name: ips; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.ips (
    id integer NOT NULL,
    ip text DEFAULT ''::text,
    "timestamp" timestamp with time zone DEFAULT now()
);


ALTER TABLE public.ips OWNER TO postgres;

--
-- Name: ips_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.ips_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.ips_id_seq OWNER TO postgres;

--
-- Name: ips_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.ips_id_seq OWNED BY public.ips.id;


--
-- Name: monthly_visitors; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.monthly_visitors (
    id integer NOT NULL,
    month text DEFAULT ''::text,
    count integer DEFAULT 0,
    year integer DEFAULT 0
);


ALTER TABLE public.monthly_visitors OWNER TO postgres;

--
-- Name: monthly_visitors_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.monthly_visitors_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.monthly_visitors_id_seq OWNER TO postgres;

--
-- Name: monthly_visitors_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.monthly_visitors_id_seq OWNED BY public.monthly_visitors.id;


--
-- Name: star_wars_characters; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.star_wars_characters (
    id integer NOT NULL,
    name text DEFAULT ''::text,
    home_world text DEFAULT ''::text,
    born text DEFAULT ''::text,
    died text DEFAULT ''::text,
    species text DEFAULT ''::text,
    gender text DEFAULT ''::text,
    associated text DEFAULT ''::text,
    affiliation text DEFAULT ''::text,
    masters text DEFAULT ''::text,
    apprentices text DEFAULT ''::text
);


ALTER TABLE public.star_wars_characters OWNER TO postgres;

--
-- Name: star_wars_characters_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.star_wars_characters_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.star_wars_characters_id_seq OWNER TO postgres;

--
-- Name: star_wars_characters_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.star_wars_characters_id_seq OWNED BY public.star_wars_characters.id;


--
-- Name: ips id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.ips ALTER COLUMN id SET DEFAULT nextval('public.ips_id_seq'::regclass);


--
-- Name: monthly_visitors id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.monthly_visitors ALTER COLUMN id SET DEFAULT nextval('public.monthly_visitors_id_seq'::regclass);


--
-- Name: star_wars_characters id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.star_wars_characters ALTER COLUMN id SET DEFAULT nextval('public.star_wars_characters_id_seq'::regclass);


--
-- Data for Name: ips; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.ips (id, ip, "timestamp") FROM stdin;
6	::1	2019-02-21 19:41:33.350487-07
7	172.16.3.43	2019-02-21 20:32:02.887307-07
8		2019-02-26 17:46:46.205683-07
\.


--
-- Data for Name: monthly_visitors; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.monthly_visitors (id, month, count, year) FROM stdin;
1	February	8	2019
\.


--
-- Data for Name: star_wars_characters; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.star_wars_characters (id, name, home_world, born, died, species, gender, associated, affiliation, masters, apprentices) FROM stdin;
5	TEST	TEST	TEST	TEST	TEST	TEST	""	"TEST"	""	""
6	TEST	TEST	TEST	TEST	TEST	TEST	TEST	TEST	TEST	TEST
7	1	2	3	4	5	6	7, 8, 9, 10	7, 8, 9, 10	7, 8, 9, 10	7, 8, 9, 10
\.


--
-- Name: ips_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.ips_id_seq', 8, true);


--
-- Name: monthly_visitors_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.monthly_visitors_id_seq', 1, true);


--
-- Name: star_wars_characters_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.star_wars_characters_id_seq', 7, true);


--
-- PostgreSQL database dump complete
--

