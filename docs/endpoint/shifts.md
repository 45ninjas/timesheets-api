# Shifts Endpoint

# `GET` v1/shifts/
Gets a paginated list of all the shifts for the current user.

**Url Arguments**
- **page** The page number to view.
- **size** The size of the pages. Maximum: 200

## 200
**Headers**
- **X-total-items** The count of total items.

**Body**
Returns a list of [shifts](../../objects/shift.md)

## 204
No shifts for the current user exist yet.


# `POST` v1/shifts/
Inserts a shift into the database.

## Body (JSON)
- **Start** *timestamp*
- **Finish** *timestamp*

**Example:**
```json
{
    "start": "2018-04-01T10:00:00Z",
    "finish": "2018-04-01T16:00:00Z"
}
```

## 201
Returns the index of the freshly inserted shift.


# `GET` v1/shifts/{shiftID}/
Gets a shift with the provided ID.

## 200
Returns a single [shift](../../objects/shift.md).

## 404
A shift with that ID does not exist. Nothing is returned.

# `DELETE` v1/shifts/{shiftID}/
Removes a shift with the provided ID.


## 200
The shift was removed, nothing is returned.

## 404
A shift with that ID does not exist. Nothing is returned.


# `PUT` v1/shifts/{shiftID}/