# API Endpoint Reference

`atimer` provides a lightweight HTTP API endpoint to manage and schedule tasks.

## POST /api

Schedules a new timer task.

### Request Format
* **Content-Type**: `application/x-www-form-urlencoded` or `multipart/form-data`

### Form Parameters
| Parameter | Type | Required | Description |
|---|---|---|---|
| `id` | `string` | **Yes** | Unique identifier for the timer task. |
| `timer_time` | `integer` | **Yes** | Duration in **seconds** before the timer expires and fires. |
| `callback_url` | `string` | **Yes** | HTTP endpoint where the POST request notification is sent upon expiration. |

### Example Request
```bash
curl -X POST http://localhost:8080/api \
  -d "id=task_1" \
  -d "timer_time=10" \
  -d "callback_url=http://example.com/callback"
```

### Success Response
* **Status**: `200 OK`
* **Content-Type**: `text/plain`
* **Body**:
```text
success: routed to heap 1
```

### Error Responses
* **Status**: `400 Bad Request`
  * Missing parameters: `Missing required parameters: id, timer_time, or callback_url`
  * Non-integer `timer_time`: `Invalid timer_time, must be an integer`
* **Status**: `405 Method Not Allowed`
  * Triggered for non-`POST` requests: `Method not allowed`
