CREATE TABLE stocks (
    id INTEGER PRIMARY KEY,
    status INTEGER NOT NULL,            -- 1: ok, 2: waitIN, 3: outed, 4: rejected
    name TEXT NOT NULL,                 -- 货品名称
    product_type INTEGER NOT NULL,      -- 货品类型 1:主材，2:辅材，3:半成品，4:成品，5:鸡货，6:报废品
    product_attr INTEGER NOT NULL,      -- 货品属性 1:医疗器械，2:非医疗器械，3:健康产品
    supplier TEXT NOT NULL,             -- 供应商
    model TEXT NOT NULL,                -- 规格型号
    unit TEXT NOT NULL,                 -- 单位
    price INTEGER NOT NULL,             -- 单价，分
    batch_no_in TEXT NOT NULL,          -- 入库批号
    way_in INTEGER NOT NULL,            -- 入库方式 1:客供，2:内供，3:外发，外购
    location INTEGER NOT NULL,          -- 存放仓库 5301,5302,5402,5403,5404,5405,5406,5407,5408,5409,5410,5411
    batch_no_produce TEXT NOT NULL,     -- 生产批号
    produce_date INTEGER NOT NULL,      -- 生产日期
    disinfection_no TEXT NOT NULL,      -- 灭菌批号
    disinfection_date INTEGER NOT NULL, -- 灭菌日期
    stock_date INTEGER NOT NULL,        -- 入库时间
    stock_num INTEGER NOT NULL,         -- 入库数量
    current_num INTEGER NOT NULL,       -- 当前数量
    value INTEGER NOT NULL              -- 当前价值
);

CREATE TABLE stock_applications (
    id INTEGER PRIMARY KEY,
    stock_id INTEGER NOT NULL,
    application_date TEXT NOT NULL,
    batch_no_in TEXT NOT NULL, -- 入库批号
    status INTEGER NOT NULL, -- 1: initiate, 2: wait approve, 3: prooved, 4: rejected
    application_user TEXT NOT NULL,
    approve_user TEXT NOT NULL,
    approve_date TEXT NOT NULL,
    create_date TEXT NOT NULL
);

CREATE TABLE stock_out_applications (
    id INTEGER PRIMARY KEY,
    stockids TEXT NOT NULL,
    number INTEGER NOT NULL,
    status INTEGER NOT NULL, -- 1: initiate, 2: wait approve, 3: prooved, 4: rejected
    application_user TEXT NOT NULL,
    approve_user TEXT NOT NULL,
    create_date TEXT NOT NULL
);
