--
-- PostgreSQL database dump
--

-- Dumped from database version 14.1 (Ubuntu 14.1-2.pgdg20.04+1)
-- Dumped by pg_dump version 14.1 (Ubuntu 14.1-2.pgdg20.04+1)

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

ALTER TABLE ONLY public.articlestable DROP CONSTRAINT articlestable_pkey;
ALTER TABLE public.articlestable ALTER COLUMN id DROP DEFAULT;
DROP SEQUENCE public.id;
DROP SEQUENCE public.articlestable_id_seq;
DROP TABLE public.articlestable;
SET default_tablespace = '';

SET default_table_access_method = heap;

--
-- Name: articlestable; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.articlestable (
    id integer NOT NULL,
    title character varying(50) NOT NULL,
    description character varying(256) NOT NULL,
    content character varying(2048) NOT NULL
);


ALTER TABLE public.articlestable OWNER TO postgres;

--
-- Name: articlestable_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.articlestable_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.articlestable_id_seq OWNER TO postgres;

--
-- Name: articlestable_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.articlestable_id_seq OWNED BY public.articlestable.id;


--
-- Name: id; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.id
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.id OWNER TO postgres;

--
-- Name: articlestable id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.articlestable ALTER COLUMN id SET DEFAULT nextval('public.articlestable_id_seq'::regclass);


--
-- Name: articlestable articlestable_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.articlestable
    ADD CONSTRAINT articlestable_pkey PRIMARY KEY (id);


--
-- PostgreSQL database dump complete
--

