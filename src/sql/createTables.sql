--
-- PostgreSQL database dump
--

-- Dumped from database version 10.10 (Ubuntu 10.10-0ubuntu0.18.04.1)
-- Dumped by pg_dump version 10.10 (Ubuntu 10.10-0ubuntu0.18.04.1)

SET statement_timeout = 0;
SET lock_timeout = 0;
SET idle_in_transaction_session_timeout = 0;
SET client_encoding = 'UTF8';
SET standard_conforming_strings = on;
SELECT pg_catalog.set_config('search_path', '', false);
SET check_function_bodies = false;
SET xmloption = content;
SET client_min_messages = warning;
SET row_security = off;

SET default_tablespace = '';

SET default_with_oids = false;

--
-- Name: ip_locations; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.ip_locations (
    id integer NOT NULL,
    as_asn integer DEFAULT 0,
    as_domain text DEFAULT ''::text,
    as_name text DEFAULT ''::text,
    as_route text DEFAULT ''::text,
    ip text DEFAULT ''::text,
    location_city text DEFAULT ''::text,
    location_country text DEFAULT ''::text,
    location_lat double precision DEFAULT 0.0,
    location_lng double precision DEFAULT 0.0,
    location_postalcode text DEFAULT ''::text,
    location_region text DEFAULT ''::text,
    location_timezone text DEFAULT ''::text,
    isp text DEFAULT ''::text
);


ALTER TABLE public.ip_locations OWNER TO postgres;

--
-- Name: ip_locations_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.ip_locations_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.ip_locations_id_seq OWNER TO postgres;

--
-- Name: ip_locations_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.ip_locations_id_seq OWNED BY public.ip_locations.id;


--
-- Name: ip_locations id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.ip_locations ALTER COLUMN id SET DEFAULT nextval('public.ip_locations_id_seq'::regclass);


--
-- Data for Name: ip_locations; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.ip_locations (id, as_asn, as_domain, as_name, as_route, ip, location_city, location_country, location_lat, location_lng, location_postalcode, location_region, location_timezone, isp) FROM stdin;
12	13415	http://www.firstdigital.com	FirstDigital	66.60.96.0/19	66.60.123.222	Sandy	US	40.5688000000000031	-111.861699999999999	84094	Utah	-06:00	FirstDigital Communications, LLC
\.


--
-- Name: ip_locations_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.ip_locations_id_seq', 12, true);


--
-- PostgreSQL database dump complete
--

