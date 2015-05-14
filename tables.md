
# Tables

## checkins

* id (integer) PK
* account (integer)
* date (timestamp with timezone)
* type (varchar) - checkin/checkout

## auth_log

* host (inet)
* account (integer)
* date (timestamp with timezone)
* realm (text) - door

## account_aliases

* account (integer)
* alias (text)

## mac

* Only 12 rows. Probably not used.

## accounts

* id (integer) PK
* name (text)
* type (text) - product/user/p2k12
* last_billed (timestamp with timezone) - only 1 populated. Probably not used.

## member

* id (integer) PK
* date (timestamp with timezone)
* full_name (text)
* email (text)
* account (integer)
* organization (text)
* price (numeric)
* recurrence (interval) - "1 mon"
* flag (varchar)

## events

* id (integer) PK
* date (timestamp with timezone)
* type (text) - "inform-door-uri"
* account (integer)
* amount (numeric) - no entries. Probably not used.
