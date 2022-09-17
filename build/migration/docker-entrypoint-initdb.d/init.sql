CREATE TABLE daily_record (
    id INT AUTO_INCREMENT,
    fund_name VARCHAR(255) NOT NULL,
    record_date DATE NOT NULL,
    amount INT,
    acquisition_price int,
    now_price int,
    theday_before int,
    theday_before_ratio FLOAT,
    profit FLOAT,
    profit_ratio FLOAT,
    valuation FLOAT,
    PRIMARY KEY (id),
    UNIQUE idx_df (fund_name, record_date),
    INDEX idx_d (record_date),
    INDEX idx_n (fund_name)
);
