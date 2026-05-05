## branches

```
CREATE TABLE branches (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    name VARCHAR(255) NOT NULL UNIQUE,

    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW(),
    deleted_at TIMESTAMP
);
```

## departments

```
CREATE TABLE departments (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    name VARCHAR(255) NOT NULL,
    parent_id UUID REFERENCES departments(id) ON DELETE SET NULL,

    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW(),
    deleted_at TIMESTAMP
);
```

## positions

```
CREATE TABLE positions (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    name VARCHAR(100) NOT NULL UNIQUE
);
```

## permissions

```
CREATE TABLE permissions (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    code VARCHAR(100) UNIQUE NOT NULL,
    name VARCHAR(255) NOT NULL,
    module VARCHAR(50)
);

CREATE UNIQUE INDEX idx_permissions_code ON permissions(code);
```

## roles

```
CREATE TABLE roles (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    name VARCHAR(100) UNIQUE NOT NULL,
    description TEXT,

    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW()
);
```

## role_permissions

```
CREATE TABLE role_permissions (
    role_id UUID REFERENCES roles(id) ON DELETE CASCADE,
    permission_id UUID REFERENCES permissions(id) ON DELETE CASCADE,
    PRIMARY KEY (role_id, permission_id)
);
```

## staffs

```
CREATE TABLE staffs (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    employee_code VARCHAR(50) UNIQUE,

    first_name VARCHAR(100),
    last_name VARCHAR(100),

    email VARCHAR(255) UNIQUE NOT NULL,
    password_hash TEXT NOT NULL,

    branch_id UUID REFERENCES branches(id) ON DELETE SET NULL,
    department_id UUID REFERENCES departments(id) ON DELETE SET NULL,
    position_id UUID REFERENCES positions(id) ON DELETE SET NULL,

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
```

## staff_roles

```
CREATE TABLE staff_roles (
    staff_id UUID REFERENCES staffs(id) ON DELETE CASCADE,
    role_id UUID REFERENCES roles(id) ON DELETE CASCADE,
    PRIMARY KEY (staff_id, role_id)
);
```

## sessions

```
CREATE TABLE sessions (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    staff_id UUID REFERENCES staffs(id) ON DELETE CASCADE,

    refresh_token TEXT NOT NULL,
    user_agent TEXT,
    ip_address TEXT,

    expired_at TIMESTAMP NOT NULL,
    created_at TIMESTAMP DEFAULT NOW()
);
```

## activity_logs

```
CREATE TABLE activity_logs (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),

    staff_id UUID,
    action VARCHAR(100),
    entity VARCHAR(100),
    entity_id UUID,

    old_data JSONB,
    new_data JSONB,

    created_at TIMESTAMP DEFAULT NOW()
);
```

## INDEX

```
CREATE INDEX idx_staffs_email ON staffs(email);
CREATE INDEX idx_staffs_employee_code ON staffs(employee_code);
CREATE INDEX idx_staffs_status ON staffs(status);
CREATE INDEX idx_staffs_branch ON staffs(branch_id);
CREATE INDEX idx_staffs_dept ON staffs(department_id);
CREATE INDEX idx_staffs_name ON staffs(first_name, last_name);
```