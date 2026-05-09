# EFV API

EFV (Eligibility & Factor Verification) API — a small Go HTTP service used to verify applicant eligibility factors.

This repository contains a lightweight HTTP API with OpenAPI documentation and a small service layer for eligibility checks. It is intentionally minimal and meant for experimentation and local development.

**Version:** 0.1.7  •  **Go:** 1.20+

---

## Quick Links

- OpenAPI spec: `docs/openapi-efv.yaml`
- Swagger UI (local): http://localhost:8080/swagger/
 - Development server (OpenShift): https://efv-api-17db4f-dev.apps.silver.devops.gov.bc.ca/swagger/

---

## Table of contents

- Overview
- Getting started
  - Prerequisites
  - Run locally
  - Build
  - Environment
- API
  - Endpoints
  - Example requests
- Project layout
- Development notes
- Contributing
- License

---

## Overview

This service exposes endpoints for verifying eligibility factors (age, residence, income, citizenship, employment). The server serves a static Swagger UI in `docs/swagger-ui/` and the OpenAPI spec at `/openapi-efv.yaml`.

The current implementation provides mocked factor definitions and a simple `Verify` service that validates input and echoes results.

---

## Getting started

### Prerequisites

- Go 1.20 or later
- Git

### Run locally (development)

macOS / Linux:

```bash
git clone https://github.com/bcgov/efv-api.git
cd efv-api
export PORT=8080   # optional, defaults are used by config
go run .
```

Windows (PowerShell):

```powershell
git clone https://github.com/bcgov/efv-api.git
cd efv-api
$env:PORT = "8080"
go run .
```

The server listens on the configured port and exposes:

- `GET /health` — health check
- `GET /api/v1/eligibility/factors` — list factor definitions (mocked)
- `POST /api/v1/eligibility/verify` — verify multiple factors
- `GET /openapi-efv.yaml` — OpenAPI spec
- `GET /swagger/` — Swagger UI (static files)

### Build (production / CI)

Build without embedding VCS metadata (useful in CI or in environments where git info isn't desired):

```bash
# build
go build -buildvcs=false -o bin/efv-api ./...
# run
PORT=8080 ./bin/efv-api
```

On Windows (PowerShell):

```powershell
go build -buildvcs=false -o bin\efv-api.exe ./...
$env:PORT = "8080"
./bin\efv-api.exe
```

---

## Configuration / Environment

This project expects configuration to come from environment variables. The primary variable used in development is:

- `PORT` — TCP port to bind the server (default: 8080 if not set)

Other configuration values are read from `config/config.go` if present.

---

## API

See the OpenAPI spec at `docs/openapi-efv.yaml` for the authoritative API surface and generated examples.

Example: Verify eligibility (curl)

```bash
curl -s -X POST http://localhost:8080/api/v1/eligibility/verify \
  -H "Content-Type: application/json" \
  -d '{"applicantId":"12345","factors":[{"type":"age","value":"25"},{"type":"residence","value":"BC"}]}' | jq
```

Expected (mock) success sample:

```json
{
  "applicantId": "12345",
  "eligible": true,
  "factors": [
    {"eligible": true, "reason": "Mock: passes", "type": "age", "value": "25"},
    {"eligible": true, "reason": "Mock: passes", "type": "residence", "value": "BC"}
  ],
  "timestamp": "2025-11-19T10:30:00Z"
}
```

Error responses follow the `models.ErrorResponse` shape in the OpenAPI spec.

---

## Project layout (high level)

- `main.go` — application entrypoint, route registration
- `config/` — configuration helper(s)
- `docs/` — OpenAPI spec and Swagger UI static files
- `internal/handlers/` — HTTP handlers + route wiring
- `internal/models/` — DTOs and response/error models
- `internal/service/eligibility/` — business logic for eligibility verification
- `internal/mocks/` — mock data used by the service
- `tools/` — (previously contained small generators/validators; may be absent)

---

## Development notes

- The OpenAPI spec in `docs/openapi-efv.yaml` is the canonical API contract.
- If you change enum values in the spec, update the corresponding values in `internal/models` to keep runtime validation in sync.
- The project intentionally uses a small hand-written service layer rather than heavy OpenAPI codegen to keep the codebase simple and easy to inspect.

---

## Contributing

Contributions are welcome. Please open issues for discussion, and submit PRs targeting the `main` branch. Keep changes small and include tests where applicable.

Before submitting a PR:

- run `go build -buildvcs=false ./...`
- run the server and exercise endpoints locally

---

## License

This repository is provided as-is. See `LICENSE` for details (if present).

---
---
