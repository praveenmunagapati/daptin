package datastore

import "github.com/artpar/api2go"

var StandardColumns = []api2go.ColumnInfo{
  {
    Name: "id",
    ColumnName: "id",
    DataType: "int(11)",
    IsPrimaryKey: true,
    IsAutoIncrement: true,
    IncludeInApi: false,
    ColumnType: "id",
  },
  {
    Name: "created_at",
    ColumnName: "created_at",
    DataType: "timestamp",
    DefaultValue: "current_timestamp",
    ColumnType: "datetime",
  },
  {
    Name: "updated_at",
    ColumnName: "updated_at",
    DataType: "timestamp",
    DefaultValue: "null",
    IsNullable: true,
    ColumnType: "datetime",
  },
  {
    Name: "deleted_at",
    ColumnName: "deleted_at",
    DataType: "timestamp",
    IncludeInApi: false,
    IsIndexed: true,
    IsNullable: true,
    ColumnType: "datetime",
  },
  {
    Name: "reference_id",
    ColumnName: "reference_id",
    DataType: "varchar(40)",
    IsIndexed: true,
    ColumnType: "alias",
  },
  {
    Name: "permission",
    ColumnName: "permission",
    IncludeInApi: false,
    DataType: "int(11)",
    IsIndexed: false,
    ColumnType: "value",

  },
  {
    Name: "status",
    ColumnName: "status",
    DataType: "varchar(20)",
    DefaultValue: "'pending'",
    IsIndexed: true,
    ColumnType: "label",
  },
}

var StandardRelations = []api2go.TableRelation{
  api2go.NewTableRelation("world_column", "belongs_to", "world"),
  api2go.NewTableRelation("action", "belongs_to", "world"),
}

var StandardTables = []TableInfo{
  {
    TableName: "world",
    IsHidden: true,
    Columns: []api2go.ColumnInfo{
      {
        Name: "table_name",
        ColumnName: "table_name",
        IsNullable:false,
        IsUnique: true,
        DataType: "varchar(30)",
        ColumnType: "name",
      },
      {
        Name: "schema_json",
        ColumnName: "schema_json",
        DataType: "text",
        IsNullable: false,
        ColumnType: "json",
      },
      {
        Name: "default_permission",
        ColumnName: "default_permission",
        DataType: "int(4)",
        IsNullable: false,
        DefaultValue: "644",
        ColumnType: "value",
      },

      {
        Name: "is_top_level",
        ColumnName: "is_top_level",
        DataType: "bool",
        IsNullable: false,
        DefaultValue: "true",
        ColumnType: "truefalse",
      },
      {
        Name: "is_hidden",
        ColumnName: "is_hidden",
        DataType: "bool",
        IsNullable: false,
        DefaultValue: "false",
        ColumnType: "truefalse",
      },

    },
  },
  {
    TableName: "world_column",
    IsHidden: true,
    Columns: []api2go.ColumnInfo{
      {
        Name: "name",
        ColumnName: "name",
        DataType: "varchar(100)",
        IsNullable: false,
        ColumnType: "name",
      },
      {
        Name: "column_name",
        ColumnName: "column_name",
        DataType: "varchar(100)",
        IsNullable: false,
        ColumnType: "name",
      },
      {
        Name: "column_type",
        ColumnName: "column_type",
        DataType: "varchar(100)",
        IsNullable: false,
        ColumnType: "label",
      },
      {
        Name: "is_primary_key",
        ColumnName: "is_primary_key",
        DataType: "bool",
        IsNullable: false,
        DefaultValue: "false",
        ColumnType: "truefalse",
      },
      {
        Name: "is_auto_increment",
        ColumnName: "is_auto_increment",
        DataType: "bool",
        IsNullable: false,
        DefaultValue: "false",
        ColumnType: "truefalse",
      },
      {
        Name: "is_indexed",
        ColumnName: "is_indexed",
        DataType: "bool",
        IsNullable: false,
        DefaultValue: "false",
        ColumnType: "truefalse",
      },
      {
        Name: "is_unique",
        ColumnName: "is_unique",
        DataType: "bool",
        IsNullable: false,
        DefaultValue: "false",
        ColumnType: "truefalse",
      },
      {
        Name: "is_nullable",
        ColumnName: "is_nullable",
        DataType: "bool",
        IsNullable: false,
        DefaultValue: "false",
        ColumnType: "truefalse",
      },
      {
        Name: "is_foreign_key",
        ColumnName: "is_foreign_key",
        DataType: "bool",
        IsNullable: false,
        DefaultValue: "false",
        ColumnType: "truefalse",
      },
      {
        Name: "include_in_api",
        ColumnName: "include_in_api",
        DataType: "bool",
        IsNullable: false,
        DefaultValue: "true",
        ColumnType: "truefalse",
      },
      {
        Name: "foreign_key_data",
        ColumnName: "foreign_key_data",
        DataType: "varchar(100)",
        IsNullable: true,
        ColumnType: "content",
      },
      {
        Name: "default_value",
        ColumnName: "default_value",
        DataType: "varchar(100)",
        IsNullable: true,
        ColumnType: "content",
      },
      {
        Name: "data_type",
        ColumnName: "data_type",
        DataType: "varchar(50)",
        IsNullable: true,
        ColumnType: "label",
      },
    },
  },
  {
    TableName: "user",
    Columns: []api2go.ColumnInfo{
      {
        Name: "name",
        ColumnName: "name",
        DataType: "varchar(80)",
        ColumnType: "name",
      },
      {
        Name: "email",
        ColumnName: "email",
        DataType: "varchar(80)",
        IsIndexed: true,
        IsUnique: true,
        ColumnType: "email",
      },

      {
        Name: "password",
        ColumnName: "password",
        DataType: "varchar(100)",
        ColumnType: "password",
        IsNullable: true,
      },
    },
  },
  {
    TableName: "usergroup",
    Columns: []api2go.ColumnInfo{
      {
        Name: "name",
        ColumnName: "name",
        DataType: "varchar(80)",
        ColumnType: "name",
      },
    },
  },
  {
    TableName: "action",
    Columns: []api2go.ColumnInfo{
      {
        Name: "action_name",
        ColumnName: "action_name",
        DataType: "varchar(100)",
        ColumnType: "name",
      },
      {
        Name: "label",
        ColumnName: "label",
        DataType: "varchar(100)",
        ColumnType: "label",
      },
      {
        Name: "in_fields",
        ColumnName: "in_fields",
        DataType: "text",
        ColumnType: "json",
      },
      {
        Name: "out_fields",
        ColumnName: "out_fields",
        DataType: "text",
        ColumnType: "json",
      },
    },
  },
}

type TableInfo struct {
  TableName         string
  TableId           int
  DefaultPermission int
  Columns           []api2go.ColumnInfo
  Relations         []api2go.TableRelation
  IsTopLevel        bool
  IsHidden          bool
}
