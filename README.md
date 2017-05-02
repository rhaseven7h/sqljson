# Golang (Go) SQL+JSON+Validator Support for Null Database/SQL Types

For fields:

- sql.NullString
- sql.NullBool
- sql.NullInt64
- sql.NullFloat64

Adds support to Null* field types for database/sql.

Validator library is:

- [Go Playground Validator: gopkg.in/go-playground/validator.v9](gopkg.in/go-playground/validator.v9)

Allows simultaneous use of validator validations, null sql values in struct fields for database/sql, and still support JSON Marshal and Unmarshal for those Null* fields.

TODO: Add instruction on registering validator custom types.

TODO: Add basic usage.
