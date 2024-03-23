CREATE TABLE IF NOT EXISTS users (
    id UUID PRIMARY KEY,
    email VARCHAR NOT NULL, -- 電子郵件
    user_name VARCHAR NOT NULL, -- 帳號
    password VARCHAR NOT NULL, -- 密碼
    phone_number VARCHAR, -- 電話號碼
    active BOOLEAN NOT NULL, -- 是否啟用
    point INTEGER NOT NULL, -- 點數
    name VARCHAR , -- 名字
    role_id UUID, -- 角色ID
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    created_by UUID,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_by UUID,
    deleted_at TIMESTAMP
);

create index idx_users_id
    on users using hash (id);


    

