CREATE TYPE tender_type AS ENUM (
    'Construction',
    'Delivery',
    'Manufacture'
);

CREATE TYPE tender_status AS ENUM (
    'Created',
    'Published',
    'Closed'
);

CREATE TABLE tender (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    status tender_status NOT NULL DEFAULT 'Created',
    organization_id UUID REFERENCES organization(id) ON DELETE CASCADE,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE tender_content (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    tender_id UUID REFERENCES tender(id) ON DELETE CASCADE,
    name VARCHAR(100) NOT NULL,
    description TEXT NOT NULL,
    type tender_type NOT NULL,
    version INTEGER NOT NULL DEFAULT 1
);

CREATE TABLE tender_content_ref (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    tender_id UUID REFERENCES tender(id) ON DELETE CASCADE,
    content_id UUID REFERENCES tender_content(id) ON DELETE CASCADE
);