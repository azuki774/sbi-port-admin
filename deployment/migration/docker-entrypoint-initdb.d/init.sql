CREATE TABLE daily_records (
    record_date DATE NOT NULL,
    fund_name VARCHAR(255) NOT NULL,
    amount INT,
    acquisition_price INT,
    now_price INT,
    theday_before INT,
    theday_before_ratio FLOAT,
    profit FLOAT,
    profit_ratio FLOAT,
    valuation FLOAT,
    PRIMARY KEY (record_date, fund_name),
    UNIQUE idx_df (fund_name, record_date),
    INDEX idx_d (record_date),
    INDEX idx_n (fund_name)
);

-- Migration v2

-- category_tag_master -> table_name の変換を行うマスタ
CREATE TABLE category_tag_master (
    category_tag_name VARCHAR(255) NOT NULL,
    table_name VARCHAR(255) NOT NULL,
    PRIMARY KEY (category_tag_name)
);

INSERT INTO `category_tag_master` (`category_tag_name`, `table_name`) VALUES
('nisa', 'daily_records');

