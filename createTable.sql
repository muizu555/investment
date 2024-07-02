CREATE TABLE TradeHistory (
  UserID VARCHAR(255) NOT NULL,
  FundID VARCHAR(255) NOT NULL,
  Quantity BIGINT NOT NULL,
  TradeDate VARCHAR(255) NOT NULL
);

CREATE TABLE ReferencePrices (
  FundID VARCHAR(255) NOT NULL,
  ReferencePriceDate VARCHAR(255) NOT NULL,
  ReferencePrice BIGINT NOT NULL
);








