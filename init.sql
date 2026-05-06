-- Create schema erp
CREATE SCHEMA IF NOT EXISTS erp;

-- Set search path so we don't have to use erp. prefix for everything if we don't want to,
-- but the code uses erp.staffs explicitly so we must ensure erp schema exists.
SET search_path TO erp, public;

-- branches
CREATE TABLE IF NOT EXISTS erp.branches (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
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
