
## Member resource

### Paths

* /member
* /member/{id}

### Json structure

* id (members.id) *
* fullName (members.full_name) *
* email (members.email) *
* created (members.date) *
* organization (members.organization) *
* price (members.price)
* recurrence (members.recurrence) - enum with "1 mon"
* flag (members.flag)
* account (href to account url /account/{id})

## User resource

### Paths

* /user
* /user/{name}
* /user/{name}/invoices
* /user/{name}/members

### Filter

* type = 'user'

### Json structure

* id (accounts.id)
* name (accounts.name)
* alias (account_aliases.alias)
* members (list of all members for account)

## Product resource

### Paths

* /product
* /product/{id}

### Filter

* type = 'product'

### Json structure

* id (accounts.id)
* name (accounts.name)
* alias (account_aliases.alias)


### Ignored columns

* members.flag (seems that all flags are NULL)
* accounts.last_billed (only used in 2 rows - let's create some billing history later)
