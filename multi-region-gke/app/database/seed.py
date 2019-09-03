from google.cloud import spanner

def create_database(instance_id, database_id):
    spanner_client = spanner.Client()
    instance = spanner_client.instance(instance_id)

    database = instance.database(database_id, ddl_statements=[
        """CREATE TABLE Categories (
             CategoryId STRING(36) NOT NULL,
             CategoryName STRING(1024),
             CategoryDescription STRING(1024)
           ) PRIMARY KEY (CategoryId);""",
        """CREATE TABLE Products (
             CategoryId STRING(36) NOT NULL,
             ProductId STRING(36) NOT NULL,
             ProductName STRING(1024),
             ProductDescription STRING(1024),
           ) PRIMARY KEY (CategoryId, ProductId),
             INTERLEAVE IN PARENT Categories ON DELETE CASCADE;""",
        """CREATE TABLE Options (
             CategoryId STRING(36) NOT NULL,
             ProductId STRING(36) NOT NULL,
             OptionId STRING(36) NOT NULL,
             OptionName STRING(1024),
             OptionDescription STRING(1024),
             OptionValue STRING(1024),
           ) PRIMARY KEY (CategoryId, ProductId, OptionId),
             INTERLEAVE IN PARENT Products ON DELETE CASCADE;"""
    ])

    operation = database.create()

    print('Waiting for operation to complete...')
    operation.result()

    print('Created database {} on instance {}'.format(database_id, instance_id))

def seed_tables(instance_id, database_id):
    spanner_client = spanner.Client()
    instance = spanner_client.instance(instance_id)

    for i in range(0,10):
       print i

if __name__ == "__main__":
    spanner_client = spanner.Client()
    instance_id = 'threecontinent'
    instance = spanner_client.instance(instance_id)
    database_id = 'ecommerce'
    try:
        database = instance.database(database_id)
    except:
        create_database(instance_id, database_id)
