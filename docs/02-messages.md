# Messages

All messages belong to a single text channel.

## Variables

| Name | Type   | Required | Description                                                        | Default Value | Validation                          |
| ---- | ------ | -------- | ------------------------------------------------------------------ | ------------- | ----------------------------------- |
| body | string | Yes      | Plaintext message content                                          |               | Length between `1` and `2048` chars |
| ttl  | number | No       | How many seconds until the message should automatically be deleted |               | maximum value of `2592000` (30 days) |