-- Create schema erp
CREATE SCHEMA IF NOT EXISTS erp;

-- Set search path so we don't have to use erp. prefix for everything if we don't want to,
-- but the code uses erp.staffs explicitly so we must ensure erp schema exists.
SET search_path TO erp, public;

-- branches
CREATE TABLE IF NOT EXISTS erp.branches (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    code VARCHAR(100) NOT NULL UNIQUE,
    name VARCHAR(255) NOT NULL UNIQUE,

    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW(),
    deleted_at TIMESTAMP
);

-- departments
CREATE TABLE IF NOT EXISTS erp.departments (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    name VARCHAR(255) NOT NULL,
    parent_id UUID REFERENCES erp.departments(id) ON DELETE SET NULL,

    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW(),
    deleted_at TIMESTAMP
);

-- positions
CREATE TABLE IF NOT EXISTS erp.positions (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    name VARCHAR(100) NOT NULL UNIQUE
);

-- permissions
CREATE TABLE IF NOT EXISTS erp.permissions (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    code VARCHAR(100) UNIQUE NOT NULL,
    name VARCHAR(255) NOT NULL,
    module VARCHAR(50)
);

CREATE UNIQUE INDEX IF NOT EXISTS idx_permissions_code ON erp.permissions(code);

-- roles
CREATE TABLE IF NOT EXISTS erp.roles (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    name VARCHAR(100) UNIQUE NOT NULL,
    description TEXT,

    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW()
);

-- role_permissions
CREATE TABLE IF NOT EXISTS erp.role_permissions (
    role_id UUID REFERENCES erp.roles(id) ON DELETE CASCADE,
    permission_id UUID REFERENCES erp.permissions(id) ON DELETE CASCADE,
    PRIMARY KEY (role_id, permission_id)
);

-- staffs
CREATE TABLE IF NOT EXISTS erp.staffs (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    employee_code VARCHAR(50) UNIQUE,

    first_name VARCHAR(100),
    last_name VARCHAR(100),

    username VARCHAR(255) UNIQUE NOT NULL, -- Changed from email to username to match code
    password_hash TEXT NOT NULL,

    branch_id UUID REFERENCES erp.branches(id) ON DELETE SET NULL,
    department_id UUID REFERENCES erp.departments(id) ON DELETE SET NULL,
    position_id UUID REFERENCES erp.positions(id) ON DELETE SET NULL,

    status VARCHAR(20) DEFAULT 'active',
    CONSTRAINT check_staff_status CHECK (
        status IN ('active', 'resigned', 'suspended')
    ),

    created_by UUID,
    updated_by UUID,

    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW(),
    deleted_at TIMESTAMP
);

-- staff_roles
CREATE TABLE IF NOT EXISTS erp.staff_roles (
    staff_id UUID REFERENCES erp.staffs(id) ON DELETE CASCADE,
    role_id UUID REFERENCES erp.roles(id) ON DELETE CASCADE,
    PRIMARY KEY (staff_id, role_id)
);

-- sessions
CREATE TABLE IF NOT EXISTS erp.sessions (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    staff_id UUID REFERENCES erp.staffs(id) ON DELETE CASCADE,

    refresh_token TEXT NOT NULL,
    user_agent TEXT,
    ip_address TEXT,

    expired_at TIMESTAMP NOT NULL,
    created_at TIMESTAMP DEFAULT NOW()
);

-- activity_logs
CREATE TABLE IF NOT EXISTS erp.activity_logs (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),

    staff_id UUID,
    action VARCHAR(100),
    entity VARCHAR(100),
    entity_id UUID,

    old_data JSONB,
    new_data JSONB,

    created_at TIMESTAMP DEFAULT NOW()
);

-- INDEX
CREATE INDEX IF NOT EXISTS idx_staffs_username ON erp.staffs(username);
CREATE INDEX IF NOT EXISTS idx_staffs_employee_code ON erp.staffs(employee_code);
CREATE INDEX IF NOT EXISTS idx_staffs_status ON erp.staffs(status);
CREATE INDEX IF NOT EXISTS idx_staffs_branch ON erp.staffs(branch_id);
CREATE INDEX IF NOT EXISTS idx_staffs_dept ON erp.staffs(department_id);
CREATE INDEX IF NOT EXISTS idx_staffs_name ON erp.staffs(first_name, last_name);

-- INSERT branches
INSERT INTO erp.branches (id, name, created_at, updated_at, deleted_at, code) VALUES ('c74528d8-4d3f-48ca-8db1-ad471e683649', 'กรุงเทพฯ', '2026-05-05 10:12:28', '2026-05-05 10:12:32.998935', NULL, 'BBK');
INSERT INTO erp.branches (id, name, created_at, updated_at, deleted_at, code) VALUES ('c90cd835-fbed-4ac7-9c33-9d1b3b955439', 'สมุทรปราการ', '2026-05-05 10:21:21.35922', '2026-05-05 10:21:21.35922', NULL, 'SKS');
INSERT INTO erp.branches (id, name, created_at, updated_at, deleted_at, code) VALUES ('2af952c1-5548-422f-9879-17f06fe95fdc', 'นนทบุรี', '2026-05-05 10:22:02.728097', '2026-05-05 10:22:02.728097', NULL, 'NTB');

-- INSERT departments
INSERT INTO erp.departments (id, name, parent_id, created_at, updated_at, deleted_at) VALUES ('1335968c-a95a-4736-ad6c-7b8910eb2db8', 'ฝ่ายบริหาร', NULL, '2026-05-05 10:15:11', '2026-05-05 10:15:13.747853', NULL);
INSERT INTO erp.departments (id, name, parent_id, created_at, updated_at, deleted_at) VALUES ('f40e64b6-b06c-4af3-9d11-18becd982ee6', 'ขาย', NULL, '2026-05-05 10:22:26.768928', '2026-05-05 10:22:26.768928', NULL);
INSERT INTO erp.departments (id, name, parent_id, created_at, updated_at, deleted_at) VALUES ('6b913531-e4a9-4ade-9fd5-bb88063ff1d9', 'คลังสินค้า', NULL, '2026-05-05 10:22:36.351768', '2026-05-05 10:22:36.351768', NULL);
INSERT INTO erp.departments (id, name, parent_id, created_at, updated_at, deleted_at) VALUES ('7350790b-c361-4bcd-9d3b-6ac159b94806', 'จัดซื้อ', NULL, '2026-05-05 10:22:44.423748', '2026-05-05 10:22:44.423748', NULL);

-- INSERT permissions
INSERT INTO erp.permissions (id, code, name, module) VALUES ('8ffe0618-9934-4f60-8e2f-dbf444f6b756', 'AD', 'admin', 'system');
INSERT INTO erp.permissions (id, code, name, module) VALUES ('2b63bac7-78a5-4397-9974-b310444d0fa2', 'ST-STAFF', 'ระบบพนักงาน', 'system');

-- INSERT positions
INSERT INTO erp.positions (id, name) VALUES ('5e6a5ea5-b617-4f69-a294-5709c66b1fc3', 'admin');
INSERT INTO erp.positions (id, name) VALUES ('7a62f48d-1130-4531-81ac-92984bb13c6f', 'sell');
INSERT INTO erp.positions (id, name) VALUES ('a08a1a9a-db37-492a-b954-794f175c847a', 'staff');

-- INSERT roles
INSERT INTO erp.roles (id, name, description, created_at, updated_at) VALUES ('5b388dc6-9fa9-4954-a3ea-b21db19a4e33', 'admin', 'แอมินของระบบ', '2026-05-05 10:16:20.580114', '2026-05-05 10:16:20.580114');
INSERT INTO erp.roles (id, name, description, created_at, updated_at) VALUES ('7cbc8bc1-10d6-43e9-82b4-7a0d574a949a', 'sell', 'พนักงานขาย', '2026-05-05 10:25:26.868449', '2026-05-05 10:25:26.868449');

-- INSERT role_permissions
INSERT INTO erp.role_permissions (role_id, permission_id) VALUES ('5b388dc6-9fa9-4954-a3ea-b21db19a4e33', '8ffe0618-9934-4f60-8e2f-dbf444f6b756');



--------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------

-- product_categories
CREATE TABLE erp.product_categories (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),

    category_code VARCHAR(50) UNIQUE,
    category_name VARCHAR(255) NOT NULL,

    status VARCHAR(20) DEFAULT 'active',

    CONSTRAINT check_product_categories_status CHECK (
        status IN ('active', 'in_active')
    ),

    created_at TIMESTAMP DEFAULT NOW()
);

-- product_brands
CREATE TABLE erp.product_brands (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),

    brand_code VARCHAR(50) UNIQUE,
    brand_name VARCHAR(255) NOT NULL,

    status VARCHAR(20) DEFAULT 'active',

    CONSTRAINT check_product_brands_status CHECK (
        status IN ('active', 'in_active')
    ),

    created_at TIMESTAMP DEFAULT NOW()
);

-- product_models
CREATE TABLE erp.product_models (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),

    models_code VARCHAR(50) UNIQUE,

    models_name VARCHAR(255) NOT NULL,

    brand_id UUID REFERENCES erp.product_brands(id),

    status VARCHAR(20) DEFAULT 'active',

    CONSTRAINT check_product_models_status CHECK (
        status IN ('active', 'in_active')
    ),

    created_at TIMESTAMP DEFAULT NOW()
);

-- products
CREATE TABLE erp.products (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),

    product_code VARCHAR(50) UNIQUE,

    product_name VARCHAR(255) NOT NULL,

    category_id UUID REFERENCES erp.product_categories(id),

    brand_id UUID REFERENCES erp.product_brands(id),

    model_id UUID REFERENCES erp.product_models(id),

    descriptions TEXT,

    unit VARCHAR(50),

    standard_price NUMERIC(18,2)
        DEFAULT 0
        CHECK (standard_price >= 0),

    status VARCHAR(20) DEFAULT 'active',

    CONSTRAINT check_products_status CHECK (
        status IN ('active', 'in_active')
    ),

    created_by UUID,
    created_at TIMESTAMP DEFAULT NOW(),

    updated_by UUID,
    updated_at TIMESTAMP
);

-- INDEX
CREATE INDEX IF NOT EXISTS idx_products_category_id
ON erp.products(category_id);

CREATE INDEX IF NOT EXISTS idx_products_brand_id
ON erp.products(brand_id);

CREATE INDEX IF NOT EXISTS idx_products_model_id
ON erp.products(model_id);

CREATE INDEX IF NOT EXISTS idx_products_product_name
ON erp.products(product_name);

CREATE INDEX IF NOT EXISTS idx_products_status
ON erp.products(status);