-- TODO: find formatter to clean this things up

CREATE SCHEMA IF NOT EXISTS loft;
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE IF NOT EXISTS loft.loft (
    id uuid PRIMARY KEY DEFAULT uuid_generate_v4(),
    name text NOT NULL,
    join_code text NOT NULL,
    created_at timestamptz NOT NULL DEFAULT NOW()
);

CREATE TABLE IF NOT EXISTS loft.join_request(
    id uuid PRIMARY KEY DEFAULT uuid_generate_v4(),
    loft_id uuid not null REFERENCES loft(id),
    name text not null,
    message text,
    created_at timestamptz not null DEFAULT NOW()
);

CREATE TABLE IF NOT EXISTS loft.member(
    id uuid PRIMARY KEY DEFAULT uuid_generate_v4(),
    loft_id uuid not null REFERENCES loft.loft(id),
    approved_at timestamptz,
    approved_by_member_id uuid,
    name text not null,
    phone text,
    email text,
    join_request_id uuid,
    FOREIGN KEY (approved_by_member_id) REFERENCES loft.member(id)
);

CREATE TABLE IF NOT EXISTS loft.note_format(
    id uuid PRIMARY KEY DEFAULT uuid_generate_v4(),
    name TEXT NOT NULL
);

INSERT INTO loft.note_format (name) VALUES (`COMMON_MARK_V_0_28`) RETURNING id;

-- TODO: find out how to return id as the new default value

CREATE TABLE IF NOT EXISTS loft.note(
    id uuid PRIMARY KEY,
    loft_id uuid REFERENCES loft.loft(id),
    creator_id uuid REFERENCES loft.member(id),
    created_at timestamptz DEFAULT NOW(),
    format uuid not null REFERENCES loft.note_format(id),
    content text not null
);

CREATE TABLE IF NOT EXISTS loft.task(
    id uuid PRIMARY KEY default uuid_generate_v4(),
    loft_id uuid not null REFERENCES loft.loft(id),
    creator_id uuid not null REFERENCES loft.member(id),
    created_at timestamptz not null DEFAULT NOW(),
    assignee_id uuid REFERENCES loft.member(id),
    title text not null,
    due_date date
);

CREATE TABLE IF NOT EXISTS loft.event(
    id	        uuid                 primary key                 default uuid_generate_v4(),
    loft_id 	uuid not null        REFERENCES loft.loft(id),
    creator_id	uuid not null        REFERENCES loft.member(id),
    created_at	timestamptz not null                             DEFAULT NOW(),
    start_time	timestamptz,
    end_time	timestamptz,
    title	    text        not null
);

CREATE TABLE IF NOT EXISTS loft.message_type(
    id uuid PRIMARY KEY DEFAULT uuid_generate_v4(),
    type text not null
);

INSERT INTO loft.message_type (type) VALUES ('TEXT') RETURNING id;
-- TODO: image reference stored on a cloud service (cloudinary/s3/sth)

CREATE TABLE IF NOT EXISTS loft.message(
    id	uuid PRIMARY KEY default uuid_generate_v4(),
    loft_id	uuid not null REFERENCES loft.loft(id),
    created_at	timestamptz not null DEFAULT NOW(),
    sender_id	uuid not null REFERENCES loft.member(id),
    content	text not null,
    type uuid not null REFERENCES loft.message_type
);

CREATE TABLE IF NOT EXISTS loft.session(
    id uuid PRIMARY KEY default uuid_generate_v4(),
    member_id uuid REFERENCES member(id),
    created_at timestamptz DEFAULT NOW(),
    last_used_at timestamptz,
    last_used_ip cidr
);