# Brettspiel Tracker API
---

## Setup

Requirements:

- Postgres Database

## API Endpoints

Get a random Boardgame from the specified User BoardgameGeek Collection


```/api/random/username=[USERNAME]```

Filter are supported as Parameters. Just append them with an "&"

### Example:
```/api/random/username=[USERNAME]&minPlayer=4```

### Supported Parameters
|  Parameter  | Description |
| ----------- | ----------- |
| username    | !required   |
| minPlayer   |             |
| maxPlayer   |             |
| minPlaytime |             |
| maxPlaytime |             |
| rating      |             |

