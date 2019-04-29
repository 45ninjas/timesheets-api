# Timesheets project
The timesheets project is designed to help freelancers track, automate and review where their time has been spent.
Timesheets is split up into multiple parts.

- **API**
    The interaction for the website and the database. This is 100% independant of the website.

- **Website**
    The website should be extremly simple and can behave like a mobile app if viewed on a mobile device. Because the API is independant of the website it can be a single-page site. It does not have to, but it can.
    the website can also be a portal for employers to view hours.

- **App**
    The app is a little different to the website because it can work offline. When the app goes online, it syncs up with the database. This might be implemented many months in the future. OR COULD BE INTERGATED INTO THE WEBSITE AS AN OFFLINE PAGE.

The timesheets api tracks timesheets for users.

The only table in the database is shifts.

# Current endpoints.

### Shifts

`/v1/shifts`

A shift is a block of time between the 'start' and 'finish' times.

Example of a few shifts.


| start               | finish              | id | total  |
|--------------------:|--------------------:|---:|-------:|
| 2018-04-01 08:00:00 | 2018-04-01 16:00:00 |  1 |  8.0000 |
| 2018-04-02 08:00:00 | 2018-04-02 14:45:00 |  2 |  6.7500 |
| 2018-04-03 10:00:00 | 2018-04-03 14:00:00 |  3 |  4.0000 |
| 2018-04-04 06:30:00 | 2018-04-04 18:00:00 |  4 | 11.5000 |
| 2018-04-05 08:00:00 | 2018-04-05 11:23:00 |  5 |  3.3833 |

Shifts can `GET`, `POST`, `DELETE` and `UPDATE`.


### Days

`/v1/days/{date}`

A day shows all the shifts that started on the provied date.

Days can only `GET`


### Frame

`/v1/frames/{frame}/{index}`

A frame is a span of `days`.

**Frame**
- week
    A week is seven days.

- fortnight
    A fortnight is 14 days (two weeks).

- month
    A month is around 4 weeks, depends on the month.

- year
    A year is mostly 365 days.

**Index**
Index is an integer . The range of the integer is defined by the Frame.
For example, if the frame is *month* the index will range from 1-12.

# Future features.

There are a few features that would be nice to add later on. However they are not important and don't have to be enabled/created for the project to work.

### Comments/tasks/category.
Sometimes you need to know what you had done during that shift. A seperate table could contain comments so not EVERY shift has a comment. You could also use existing comments for multiple shifts.

### Employers/clients
Same as comments, we can add a table that stores employers. We can assign each shift to an employer.

### Sceduled submissions
Each *employer* can be sent an email or pdf summary of the last *frame*.