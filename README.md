# 🛒 OpenMarket — E-Commerce RESTful API Backend

> A full-featured **multi-vendor e-commerce** backend built from scratch with **Golang**, following **Clean Architecture** principles. The system supports the complete online shopping lifecycle — from user registration, product management, cart operations to transactional order checkout.

---

## 📌 Project Highlights

| Aspect             | Detail                                      |
| ------------------ | ------------------------------------------- |
| **Language**       | Go 1.25                                     |
| **Framework**      | Gin (HTTP) + GORM (ORM)                     |
| **Database**       | MySQL                                       |
| **Authentication** | JWT (HS256) with role-based access control  |
| **File Storage**   | AWS S3 / MinIO (compatible)                 |
| **Architecture**   | Clean Architecture (4-layer)                |
| **API Style**      | RESTful API v1 with JSON envelope responses |

---

## 🏗️ Architecture & Design Patterns

The project strictly follows **Clean Architecture** with 4 well-defined layers per module:

```
Transport (Gin Handlers)
    ↓
Business (Use Cases / Domain Logic)
    ↓
Repository (Interface Contracts)
    ↓
Storage (GORM / Database Implementation)
```

### Key Design Patterns Applied

| Pattern                           | Implementation                                                                                                                         |
| --------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------- |
| **Clean Architecture**            | Strict layer separation — business logic has zero dependency on frameworks                                                             |
| **Repository Pattern**            | All database access goes through interfaces; business layer depends only on abstractions                                               |
| **Unit of Work**                  | Cross-module database transactions via `TxStore` — composes multiple storage implementations on a single `gorm.DB` transaction         |
| **Dependency Injection**          | Business structs receive interface dependencies via constructors (`NewXxxBusiness(repo Interface)`)                                    |
| **DTO Separation**                | Dedicated structs for Create, Update, View, Filter per entity — domain models are never used directly for request/response             |
| **ID Obfuscation**                | Internal integer IDs are never exposed to clients; all external IDs use Base58-encoded composite UIDs (LocalID + ObjectType + ShardID) |
| **Soft Delete**                   | Status field pattern (1=active, 0=deleted) instead of hard deletes                                                                     |
| **Address Snapshotting**          | Order addresses are copied at checkout, decoupled from the user's mutable address book                                                 |
| **Decimal Money**                 | Uses `shopspring/decimal` for all monetary calculations — avoids floating-point precision issues                                       |
| **Centralized Error Handling**    | Rich `AppError` system with ~25 domain-specific error constructors and recursive error unwrapping                                      |
| **Panic-based Error Propagation** | Transport layer panics with `*AppError`, caught globally by Recover middleware — clean error flow in Gin handlers                      |

---

## 🧩 Module Overview

### 1. User Module (`module/user/`)

- **Registration & Login** — email/password authentication with secure password hashing (bcrypt)
- **Profile management** — view, update, soft-delete user account
- **Password change** — with old password verification
- **Address book** — full CRUD for shipping addresses (create, update, delete, list)

### 2. Seller Module (`module/seller/`)

- **Shop creation** — authenticated users can register as sellers (one shop per user, enforced)
- **Shop management** — update shop info, view own shop, soft-delete
- **Public listing** — browse all sellers with filtering

### 3. Product Module (`module/product/`) — _Richest domain, 25+ business logic files_

- **Product CRUD** — create, update, soft-delete products (seller-scoped ownership validation)
- **Variant management** — create product variants with SKU auto-generation, update, delete, toggle status
- **Duplicate detection** — check for duplicate attribute combinations before variant creation
- **Category system** — hierarchical categories with parent-child relationships
- **Attribute system** — flexible product attributes (e.g., Color, Size) linked to categories
- **Variant-Attribute mapping** — M:N relationship linking each variant to specific attribute values
- **Gallery management** — multiple product images with "main image" designation
- **Stock management** — adjust stock quantities per variant
- **Advanced filtering** — filter by seller, status, category, price range, keyword search
- **Dual views** — separate public product detail and seller management detail views

### 4. Cart Module (`module/cart/`)

- **Auto-creation** — cart is automatically created on first item addition
- **Add to cart** — with automatic quantity merge for existing items
- **Quantity adjustment** — update item quantities
- **Remove items** — remove specific items from cart
- **Detailed view** — enriched cart item list with product info, prices, stock status, and attributes
- **Transactional operations** — all cart mutations wrapped in database transactions

### 5. Order Module (`module/order/`)

- **Transactional checkout** — complete 8-step checkout flow in a single database transaction:
  1. Validate cart (must not be empty)
  2. Validate shipping address (must belong to user)
  3. Create order record (status = pending)
  4. Snapshot shipping address to order
  5. For each cart item: validate stock → load variant → calculate subtotal → create order item → **decrement stock atomically**
  6. Calculate and update total amount
  7. Clear cart after successful checkout
- **Order status FSM** — `pending → confirmed → shipping → completed` / `cancelled`
- **Order cancellation** — with status validation
- **Order listing** — with filters (status, price range) and pagination
- **Order detail** — enriched view with product names, images, attributes

### 6. Admin Module (`module/admin/`)

- **Separate authentication** — dedicated staff login system (independent from user auth)
- **Role-Based Access Control (RBAC)** — 3-tier roles: `admin > moderator > support`
- **Seller moderation** — list, view, update status, lock/unlock seller accounts
- **Staff profile** — view own staff info
- **Dev seed endpoint** — initial staff account creation for development

### 7. Upload Module (`module/upload/`)

- **Image upload** — supports AWS S3 and MinIO (S3-compatible)
- **Auto content-type detection**
- **Authenticated uploads** — requires JWT token

---

## 🔐 Authentication & Authorization

| Layer               | Mechanism                                                                                       |
| ------------------- | ----------------------------------------------------------------------------------------------- |
| **User Auth**       | JWT Bearer token in `Authorization` header → validate → load user from DB → check active status |
| **Seller Auth**     | Same JWT + verify user owns a seller shop                                                       |
| **Admin Auth**      | Separate JWT flow → loads staff from `staff_accounts` with roles                                |
| **RBAC Middleware** | `RequireRole(roles...)` — checks if staff has required role before proceeding                   |

---

## 📡 API Endpoints (40+ endpoints)

### Public

| Method | Endpoint                        | Description                  |
| ------ | ------------------------------- | ---------------------------- |
| POST   | `/v1/auth/login`                | User login                   |
| POST   | `/v1/users`                     | User registration            |
| GET    | `/v1/sellers`                   | List all sellers             |
| GET    | `/v1/sellers/:id`               | Get seller detail            |
| GET    | `/v1/categories`                | List categories              |
| GET    | `/v1/categories/:id/attributes` | Get category attributes      |
| GET    | `/v1/products`                  | List products (with filters) |
| GET    | `/v1/products/:id`              | Product detail               |
| GET    | `/v1/products/:id/variants`     | List variants                |
| GET    | `/v1/products/:id/galleries`    | Get product images           |
| GET    | `/v1/products/:id/attributes`   | Get product attributes       |

### Authenticated (User)

| Method | Endpoint             | Description        |
| ------ | -------------------- | ------------------ |
| GET    | `/v1/users/profile`  | Get profile        |
| PATCH  | `/v1/users`          | Update profile     |
| PATCH  | `/v1/users/password` | Change password    |
| DELETE | `/v1/users`          | Delete account     |
| CRUD   | `/v1/address/*`      | Address management |
| CRUD   | `/v1/sellers/*`      | Shop management    |
| CRUD   | `/v1/carts/*`        | Cart operations    |
| CRUD   | `/v1/orders/*`       | Order operations   |
| POST   | `/v1/upload`         | Image upload       |

### Seller Dashboard

| Method | Endpoint                              | Description        |
| ------ | ------------------------------------- | ------------------ |
| CRUD   | `/v1/seller/products/*`               | Product management |
| CRUD   | `/v1/seller/products/:id/variants/*`  | Variant management |
| CRUD   | `/v1/seller/products/:id/galleries/*` | Gallery management |

### Admin Panel

| Method | Endpoint                       | Description               |
| ------ | ------------------------------ | ------------------------- |
| POST   | `/v1/admin/auth/login`         | Staff login               |
| GET    | `/v1/admin/profile`            | Staff profile             |
| GET    | `/v1/admin/sellers`            | List sellers (moderation) |
| PATCH  | `/v1/admin/sellers/:id/status` | Update seller status      |
| POST   | `/v1/admin/sellers/:id/lock`   | Lock seller               |
| POST   | `/v1/admin/sellers/:id/unlock` | Unlock seller             |

---

## 📂 Project Structure

```
OpenMarket/
├── main.go                    # Application entry point
├── main_route.go              # Centralized route registration
├── common/                    # Shared types, errors, utilities
│   ├── app_error.go           # Rich error system (~25 error types)
│   ├── uid.go                 # Base58 UID encoding/decoding
│   ├── sql_model.go           # Base model with soft-delete
│   ├── paging.go              # Offset + cursor pagination
│   └── image.go               # JSONB image type
├── component/                 # Infrastructure components
│   ├── appctx/                # Dependency injection container
│   ├── hasher/                # Password hashing (bcrypt)
│   ├── tokenprovider/         # JWT token provider
│   ├── uploadProvider/        # S3/MinIO file upload
│   └── validate/              # Email validation
├── middleware/                 # HTTP middleware
│   ├── authenticate.go        # User JWT authentication
│   ├── admin_auth.go          # Admin JWT + RBAC
│   └── recover.go             # Global error recovery
└── module/                    # Domain modules (Clean Arch)
    ├── user/                  # User management
    ├── seller/                # Seller/shop management
    ├── product/               # Products, variants, categories
    ├── cart/                  # Shopping cart
    ├── order/                 # Order & checkout
    ├── admin/                 # Admin panel & moderation
    └── upload/                # File upload
```

Each module follows the same internal structure:

```
module/<name>/
├── model/         # Domain models + DTOs
├── business/      # Use-case logic
├── repository/    # Interface definitions
├── storage/       # Database implementation (GORM)
└── transport/gin/ # HTTP handlers
```

---

## ⚙️ Technical Highlights

- **Transactional Order Checkout** — The order creation process orchestrates across 4 modules (order, cart, product, user) within a single database transaction using the Unit of Work pattern. Stock decrement, order creation, and cart clearing are atomic.

- **Custom UID System** — Internal auto-increment IDs are never exposed. A composite UID (LocalID + ObjectType + ShardID) encoded in Base58 provides URL-friendly, type-safe, obfuscated identifiers.

- **Flexible Product Catalog** — Supports hierarchical categories, configurable attributes per category, multi-variant products with attribute combinations, and a gallery system with main image selection.

- **Dual Authentication System** — Separate authentication flows for end-users and admin staff, with RBAC middleware supporting role hierarchy.

- **Standardized API Responses** — All endpoints return consistent JSON envelopes with `data`, `paging`, and `filter` fields. Error responses include structured error keys for client-side handling.

- **Production-Ready Error Handling** — Centralized error types with HTTP status mapping, recursive error unwrapping for debugging, and global panic recovery middleware.

---

## 🚀 Getting Started

### Prerequisites

- Go 1.25+
- MySQL 8.0+
- AWS S3 or MinIO (for file uploads)

### Environment Variables

```env
DB_USER=root
DB_PASSWORD=your_password
DB_HOST=localhost
DB_PORT=3306
DB_NAME=open_market

SYSTEM_SECRET=your_jwt_secret

S3BucketName=your_bucket
S3Region=ap-southeast-1
S3APIKey=your_key
S3SecretKey=your_secret
S3_ENDPOINT=your_endpoint
S3_PUBLIC_DOMAIN=your_domain

PORT=:8080
```

### Run

```bash
# Clone the repository
git clone https://github.com/dykh061/OpenMarket.git
cd OpenMarket

# Install dependencies
go mod tidy

# Set up environment variables
cp .env.example .env  # Edit with your config

# Run the server
go run .
```

---

## 👤 Author

**Duy Khanh**  
GitHub: [@dykh061](https://github.com/dykh061)

---

_This project was built as a portfolio project to demonstrate backend engineering skills including system design, clean architecture implementation, transactional business logic, and production-ready API development with Go._
