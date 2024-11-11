SET statement_timeout = 0;
SET lock_timeout = 0;
SET client_encoding = 'UTF8';
SET standard_conforming_strings = on;
SET check_function_bodies = false;
SET client_min_messages = warning;

--
-- Name: plpgsql; Type: EXTENSION; Schema: -; Owner:
--

CREATE EXTENSION IF NOT EXISTS plpgsql WITH SCHEMA pg_catalog;


--
-- Name: EXTENSION plpgsql; Type: COMMENT; Schema: -; Owner:
--

COMMENT ON EXTENSION plpgsql IS 'PL/pgSQL procedural language';


CREATE FUNCTION created_at_column() RETURNS trigger
    LANGUAGE plpgsql
    AS $$

BEGIN
	NEW.updated_at = EXTRACT(EPOCH FROM NOW());
	NEW.created_at = EXTRACT(EPOCH FROM NOW());
    RETURN NEW;
END;

$$;


ALTER FUNCTION public.created_at_column() OWNER TO postgres;

--
-- TOC entry 190 (class 1255 OID 36646)
-- Name: update_at_column(); Type: FUNCTION; Schema: public; Owner: postgres
--

CREATE FUNCTION update_at_column() RETURNS trigger
    LANGUAGE plpgsql
    AS $$

BEGIN
    NEW.updated_at = EXTRACT(EPOCH FROM NOW());
    RETURN NEW;
END;

$$;


ALTER FUNCTION public.update_at_column() OWNER TO postgres;


SET search_path = public, pg_catalog;

SET default_tablespace = '';

SET default_with_oids = false;

--
-- Name: campaign; Type: TABLE; Schema: public; Owner: postgres; Tablespace:
--


CREATE TABLE "user" (
    id uuid DEFAULT gen_random_uuid(),
    email character varying,
    password character varying,
    name character varying,
    user_type character varying,
    updated_at TIMESTAMP NOT NULL DEFAULT NOW(),
    created_at TIMESTAMP NOT NULL DEFAULT NOW()
);


ALTER TABLE "user" OWNER TO postgres;

ALTER TABLE ONLY "user"
    ADD CONSTRAINT user_id PRIMARY KEY (id);


CREATE TRIGGER create_user_created_at BEFORE INSERT ON "user" FOR EACH ROW EXECUTE PROCEDURE created_at_column();
CREATE TRIGGER update_user_updated_at BEFORE UPDATE ON "user" FOR EACH ROW EXECUTE PROCEDURE update_at_column();

CREATE TABLE campaigns (
    id uuid DEFAULT gen_random_uuid(),
    name VARCHAR(255),
    discount_rate DECIMAL(5, 2) CHECK (discount_rate <= 100),
    voucher_limit INT,
    start_date TIMESTAMP,
    end_date TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT NOW(),
    created_at TIMESTAMP NOT NULL DEFAULT NOW()
);

ALTER TABLE "campaigns" OWNER TO postgres;

ALTER TABLE ONLY "campaigns"
    ADD CONSTRAINT campaign_id PRIMARY KEY (id);


REVOKE ALL ON SCHEMA public FROM PUBLIC;
REVOKE ALL ON SCHEMA public FROM postgres;
GRANT ALL ON SCHEMA public TO postgres;
GRANT ALL ON SCHEMA public TO PUBLIC;

