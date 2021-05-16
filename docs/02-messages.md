# Messages

All messages belong to a single text channel.

## Variables

| Name | Type   | Required | Description                                                        | Default Value | Validation                          |
| ---- | ------ | -------- | ------------------------------------------------------------------ | ------------- | ----------------------------------- |
| body | string | Yes      | Plaintext message content                                          |               | Length between `1` and `2048` chars |
| ttl  | number | No       | How many seconds until the message should automatically be deleted |               | maximum value of `2592000` (30 days)|

## Example Requests

### Create a Message

**Method:** `POST`
**URI:** `/channels/{{ CHANNEL_ID }}/messages`

_Request Body:_

```json
{
  "body": "Did you ever hear the tragedy of Darth Plagueis The Wise?"
}
```

_Response body:_

```json
{
    "id": "1sauqZrq3X6l5NeEkQYh6SkzgG7",
    "channel_id": "1saXxyQcF4gDITDsCWOolIDwIH1",
    "body": "Did you ever hear the tragedy of Darth Plagueis The Wise?",
    "ttl": 1337,
    "created_at": "2021-05-16T00:00:12.164025714Z",
    "updated_at": "2021-05-16T00:00:12.164039005Z"
}
```

### Listing all messages

**Method:** `GET`
**URI:** `/channels/{{ CHANNEL_ID }}/messages`

This endpoint supports cursor-based pagination via the `cursor` and `count` headers/query-parameters, the pagination is limited to returning `1000` messages in a single response.

_Response body:_

```json
[
    {
        "id": "1saYRhaS48p1LkJmiWu1bTo0ggT",
        "channel_id": "1saXxyQcF4gDITDsCWOolIDwIH1",
        "body": "foobar",
        "created_at": "2021-05-15T20:56:00.21899005Z",
        "updated_at": "2021-05-15T20:56:00.219004772Z"
    },
    {
        "id": "1saYRxoZyx9rPrqrs3CQPQSctFK",
        "channel_id": "1saXxyQcF4gDITDsCWOolIDwIH1",
        "body": "foobar",
        "created_at": "2021-05-15T20:56:02.119908041Z",
        "updated_at": "2021-05-15T20:56:02.119921869Z"
    },
    {
        "id": "1saYUxcYslGuaRaJRu3ZH6RWiFe",
        "channel_id": "1saXxyQcF4gDITDsCWOolIDwIH1",
        "body": "foobar",
        "created_at": "2021-05-15T20:56:26.692884694Z",
        "updated_at": "2021-05-15T20:56:26.692914437Z"
    }
]
```
