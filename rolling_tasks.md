

## Store cookie client side to track daily intake

- Server needs to generate Cookie
- Create Table with list of cookie to current state
- Ensure data and cookie will be deleted at 00:00 every day
    - means a cookie must only be valid until the end of a given day
- Allow to add intakes to the current day
- Allow to remove intake from the current day????


## Implement consensus algorithm for verification of an added food item

- if a item is being added, the item will be not be marks as valid unless the item was approved by others (how many??? I dont know)
- buffer newly added items and present them to connected users???? SSE Events: while on page client has open SSE to receive verification questions


## Implement search recipe