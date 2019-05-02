# Shift Object

A shift stores a frame of time between a start and finish timestamp.

- **id** *uint*
    
    The unique identifier for this shift.

- **start** *timestamp*

    The start time for this shift.

- **finish** *timestamp*

    The finish time for this shift.

- **total** *float* (readonly)

    The amount of hours between the start and finish. 
    >_NOTE: This is computed on every SELECT and is not stored in the database._

## Json Example

```json
{
    "id": 1,
    "start": "2018-04-01T08:00:00Z",
    "finish": "2018-04-01T16:00:00Z",
    "total": 8
}
```