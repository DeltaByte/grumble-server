# Channels

## Variables

| Name    | Type   | Required | Description                                            | Channel Type | Default Value | Validation                           |
| ------- | ------ | -------- | ------------------------------------------------------ | ------------ | ------------- | ------------------------------------ |
| type    | string | Yes      | Channel Type                                           | Both         |               | Must be either `text` or `voice`     |
| name    | string | Yes      | Name of the channel                                    | Both         |               | Max length is `100` chars            |
| bitrate | number | No       | Sound quality (kbps), only available                   | Voice        | `64`          | Min value of `4`, max value of `255` |
| topic   | string | No       | Discussion topic                                       | Text         |               | Max length of `1024` chars           |
| nsfw    | bool   | No       | Wether the channel is expected to contain NSFW content | Text         | `false`       |                                      |

## Example Requests

### Creating a channel

**Method:** `POST`
**URI:** `/channels`

_Request Body (voice):_

```json
{
  "type": "voice",
  "name": "Example Channel",
  "bitrate": 64
}
```

_Request Body (text):_

```json
{
  "type": "text",
  "name": "Example Channel",
  "topic": "Q: What's tiny and yellow and very, very, dangerous?\nA: canary with the super-user password.",
  "nsfw": false
}
```

_Response Body:_

Expected HTTP status: `201`

```json
{
  "id": "1sasn9LLHI9IEmt9deyL7VbmlBb",
  "type": "text",
  "name": "Example Channel",
  "topic": "",
  "nsfw": false,
  "created_at": "0001-01-01T00:00:00Z",
  "updated_at": "0001-01-01T00:00:00Z"
}
```

### Listing all channels

**Method:** `GET`
**URI:** `/channels`

Because it is assumed that any clients will generally need to know all channels anyway, and there is unlikely to be a large quantity of channel, this endpoint does not currently support pagination.

_Response Body:_

```json
[
  {
    "id": "1sXjcGVow7fnuVbT6bcEvDVNasj",
    "type": "text",
    "name": "Test Channel 2",
    "topic": "very witty example channel topic",
    "nsfw": false,
    "created_at": "0001-01-01T00:00:00Z",
    "updated_at": "0001-01-01T00:00:00Z"
  },
  {
    "id": "1sXqfD3bUrueXFlNl2fypmzfUtG",
    "type": "voice",
    "name": "Test Channel 2",
    "bitrate": 64,
    "created_at": "0001-01-01T00:00:00Z",
    "updated_at": "0001-01-01T00:00:00Z"
  },
  {
    "id": "1sXrvYtijNiwc4Z1RqD7CwaqGlH",
    "type": "voice",
    "name": "Test Channel 2",
    "bitrate": 128,
    "created_at": "0001-01-01T00:00:00Z",
    "updated_at": "0001-01-01T00:00:00Z"
  }
]
```
