CREATE TABLE stocks (
	id INTEGER PRIMARY KEY,
	name TEXT NOT NULL,
    status INTEGER NOT NULL, -- 1: ok, 2: waitIN, 3: outed
	product_type INTEGER NOT NULL, -- 成品类型 1:原材料，2:成品
	type INTEGER NOT NULL, -- 货品类型
	supplier TEXT NOT NULL, -- 供应商
    model TEXT NOT NULL, -- 规格型号
    unit TEXT NOT NULL, -- 单位
    price INTEGER NOT NULL, -- 单价，分
    batch_no_in TEXT NOT NULL, -- 入库批号
    way_in TEXT NOT NULL, -- 入库方式
    location TEXT NOT NULL, -- 存放仓库
    batch_no_produce TEXT NOT NULL, -- 生产批号
    produce_date TEXT NOT NULL, -- 生产日期
    stock_date TEXT NOT NULL, -- 入库时间
    stock_num INTEGER NOT NULL, -- 入库数量
    current_num INTEGER NOT NULL, -- 当前数量
    value INTEGER NOT NULL -- 当前价值
);

CREATE TABLE stock_applications (
	id INTEGER PRIMARY KEY,
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
  status INTEGER NOT NULL, -- 1: initiate, 2: wait approve, 3: prooved, 4: rejected
  application_user TEXT NOT NULL,
  approve_user TEXT NOT NULL,
  create_date TEXT NOT NULL
);
