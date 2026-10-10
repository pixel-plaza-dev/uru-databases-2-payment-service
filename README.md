# uru-databases-2-payment-service

Payment microservice for Pixel Plaza, intended to record order and branch rent payments over gRPC with PostgreSQL. Only the scaffolding exists.

**Note:** This repository is archived and read-only. It has no `main.go` and no gRPC server yet.

## What exists

- **`model/postgresql/types.go`** — GORM models `User`, `Payment`, `PaymentAccount`, `OrderPayments`, `OrderPayment`, `BranchRentPayments` and `BranchRentPayment`.
- **`app/`** — constants for the listener (`PORT`) and the JWT public key (`PIXEL_PLAZA_JWT_PUBLIC_KEY`) environment variables.
- **`Dockerfile`** — builds a Go binary and starts it with `-m=prod` on port 8080. The build expects a `main` package that is not in the repo.
- **`go.mod`** — module `github.com/pixel-plaza-dev/uru-databases-2-payment-service`, Go 1.23.4, using `gorm.io/gorm`.

The intended API is the `Payment` service in `uru-databases-2-protobuf-common` (payment accounts, order and branch rent payments, verification).

## License

GPL-3.0. See [LICENSE](LICENSE).
