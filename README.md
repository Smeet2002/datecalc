# datecalc

datecalc
=======

#### Simplest date and time calculator

datecalc is a small program written in GO calculating date and time of the future and past events.\
It is accepting the following argument : "+N" or "-N", where N is the number.\
"+" = calculating future dates from now.\
"-" = calculating past dates from now.\
"N" = the number used to calculate "seconds", "minutes", "hours", "days", "months" and "years" from now.


How to use
-----------

    datecalc +10

    Current date and time: 2024-08-01 14:48 seconds: 31

    +10 seconds: 2024-08-01 14:48 seconds: 41
    +10 minutes: 2024-08-01 14:58 seconds: 31
    +10 hours: 2024-08-02 00:48
    +10 days: 2024-08-11 14:48
    +10 months: 2025-06-01
    +10 years: 2034-08-01


    datecalc -15

    Current date and time: 2024-08-01 14:46 seconds: 57

    -15 seconds: 2024-08-01 14:46 seconds: 42
    -15 minutes: 2024-08-01 14:31 seconds: 57
    -15 hours: 2024-07-31 23:46
    -15 days: 2024-07-17 14:46
    -15 months: 2023-05-01
    -15 years: 2009-08-01

