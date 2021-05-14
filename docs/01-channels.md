# Channels

## Variables

| Name    | Type   | Required | Description                                            | Channel Type | Default Value | Validation                           |
| ------- | ------ | -------- | ------------------------------------------------------ | ------------ | ------------- | ------------------------------------ |
| type    | string | Yes      | Channel Type                                           | Both         |               | Must be either `text` or `voice`     |
| name    | string | Yes      | Name of the channel                                    | Both         |               | Max length is `100` chars            |
| bitrate | number | No       | Sound quality (kbps), only available                   | Voice        | `64`          | Min value of `4`, max value of `255` |
| topic   | string | No       | Discussion topic                                       | Text         |               | Max length of `1024` chars           |
| nsfw    | bool   | No       | Wether the channel is expected to contain NSFW content | Text         | `false`       |                                      |

## Example objects

## Voice

```json
{
  "id": "1sXx1jTWyKEZBdpuoTnhn1tvIdO",
  "type": "voice",
  "name": "Example Channel",
  "bitrate": 64
}
```

## Text

```json
{
  "id": "1sXx7zc9kNjxy7Ad0SeI4LFuFAc",
  "type": "text",
  "name": "Example Channel",
  "topic": "",
  "nsfw": false
}
```
