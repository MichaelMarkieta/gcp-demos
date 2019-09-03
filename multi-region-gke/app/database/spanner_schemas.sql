CREATE TABLE Categories (
  CategoryId STRING(36) NOT NULL,
  CategoryName STRING(1024),
  CategoryDescription STRING(1024)
) PRIMARY KEY (CategoryId);

CREATE TABLE Products (
  CategoryId STRING(36) NOT NULL,
  ProductId STRING(36) NOT NULL,
  ProductName STRING(1024),
  ProductDescription STRING(1024),
) PRIMARY KEY (CategoryId, ProductId),
  INTERLEAVE IN PARENT Categories ON DELETE CASCADE;

CREATE TABLE Options (
  CategoryId STRING(36) NOT NULL,
  ProductId STRING(36) NOT NULL,
  OptionId STRING(36) NOT NULL, 
  OptionName STRING(1024),
  OptionDescription STRING(1024),
  OptionValue STRING(1024),
) PRIMARY KEY (CategoryId, ProductId, OptionId),
  INTERLEAVE IN PARENT Products ON DELETE CASCADE;
