# Prometheus metrics
Recon have basic prometheus metrics on root path
```http request
GET /

### 

recon_request_duration_seconds_bucket{code="200",le="0.005"} 2
recon_request_duration_seconds_bucket{code="200",le="0.01"} 2
recon_request_duration_seconds_bucket{code="200",le="0.025"} 2
```