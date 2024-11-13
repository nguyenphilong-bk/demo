SET statement_timeout = 0;
SET lock_timeout = 0;
SET client_encoding = 'UTF8';
SET standard_conforming_strings = on;
SET check_function_bodies = false;
SET client_min_messages = warning;

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

CREATE TABLE campaigns (
    id uuid DEFAULT gen_random_uuid(),
    name VARCHAR(255),
    discount_rate DECIMAL(5, 2) CHECK (discount_rate <= 100),
    voucher_limit INT,
    start_date TIMESTAMP,
    end_date TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT NOW(),
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    created_by uuid,
    deleted_by uuid
);

ALTER TABLE "campaigns" OWNER TO postgres;

ALTER TABLE ONLY "campaigns"
    ADD CONSTRAINT campaign_id PRIMARY KEY (id);

CREATE TABLE vouchers (
    id uuid DEFAULT gen_random_uuid(),
    campaign_id uuid,
    user_id uuid,
    discount_rate DECIMAL(5, 2) CHECK (discount_rate <= 100),
    code VARCHAR(255),
    status VARCHAR(255),
    expiration_date TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT NOW(),
    created_at TIMESTAMP NOT NULL DEFAULT NOW()
);

ALTER TABLE "vouchers" OWNER TO postgres;

ALTER TABLE ONLY "vouchers"
    ADD CONSTRAINT campaign_id PRIMARY KEY (id);

REVOKE ALL ON SCHEMA public FROM PUBLIC;
REVOKE ALL ON SCHEMA public FROM postgres;
GRANT ALL ON SCHEMA public TO postgres;
GRANT ALL ON SCHEMA public TO PUBLIC;

