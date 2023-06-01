# Auth
Google SSO

# Postgres Database

## Table: Users
### Fields
userID (PK): string (google user ID)

## Table: Timers
### Fields
timerID (PK): uuid
name: string
initialSeconds: int64
remainingSeconds: int64
paused: bool
owner: FK(Users(userID))

## Table: EditAccess
### Fields
userID: FK(Users(userID))
timerID: FK(Timers(timerID))

## Table: ReadAccess
### Fields
userID: FK(Users(userID))
timerID: FK(Timers(timerID))