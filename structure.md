# Cây thư mục dự án OpenMarket

## Sơ đồ thư mục

```text
OpenMarket/
├─ .env
├─ .git/
├─ .gitignore
├─ go.mod
├─ go.sum
├─ main.go
├─ structure.md
├─ common/
│  ├─ app_error.go
│  ├─ app_response.go
│  ├─ const.go
│  ├─ image.go
│  ├─ paging.go
│  └─ sql_model.go
├─ component/
│  ├─ appctx/
│  │  └─ app_context.go
│  ├─ hasher/
│  │  └─ hasher.go
│  ├─ tokenprovider/
│  │  ├─ provider.go
│  │  └─ jwt/
│  │     └─ jwt.go
│  └─ uploadProvider/
│     ├─ provider.go
│     └─ aws_s3.go
├─ middleware/
│  ├─ authenticate.go
│  └─ recover.go
├─ module/
│  ├─ cart/ (trống)
│  ├─ inventory/ (trống)
│  ├─ order/ (trống)
│  ├─ variant/ (trống)
│  ├─ product/
│  │  ├─ business/
│  │  ├─ model/
│  │  ├─ storage/
│  │  └─ transport/gin/
│  ├─ seller/
│  │  ├─ business/
│  │  ├─ model/
│  │  ├─ storage/
│  │  └─ transport/gin/
│  ├─ upload/
│  │  ├─ business/
│  │  │  └─ upload.go
│  │  └─ transport/gin/
│  │     └─ upload_image.go
│  └─ user/
│     ├─ business/
│     ├─ model/
│     ├─ storage/
│     └─ transport/gin/
```

## Mô tả nhanh

- **common/**: Các kiểu phản hồi chung, error handling, image model, phân trang, và model SQL cơ bản.
- **component/**: `appctx` quản lý context/kết nối DB; `hasher` mã hoá mật khẩu; `tokenprovider` JWT auth; `uploadProvider` S3/MinIO.
- **middleware/**: `authenticate` xác thực JWT; `recover` xử lý panic.
- **module/**: Các domain chính (user, seller, product, upload đã có code; auth/cart/inventory/order/variant còn trống).
- **main.go**: Khởi tạo ứng dụng, kết nối DB, đăng ký router Gin.

---

## Changelog & Optimizations

### v1.1.0 - Upload Image & S3/MinIO Integration (2025-12-26)

#### ✅ Tính năng mới

| Tính năng                 | Mô tả                                     | File liên quan                     |
| ------------------------- | ----------------------------------------- | ---------------------------------- |
| Upload ảnh lên MinIO/S3   | Hỗ trợ upload ảnh qua multipart/form-data | `module/upload/`                   |
| Image dimension detection | Tự động đọc width/height của ảnh          | `module/upload/business/upload.go` |
| S3 Provider abstraction   | Interface cho upload provider             | `component/uploadProvider/`        |

#### 🔧 Cấu hình S3/MinIO (.env)

```env
S3BucketName=duy
S3Region=us-east-1
S3APIKey=minioadmin
S3SecretKey=minioadmin
S3_ENDPOINT=http://localhost:9000
S3_PUBLIC_DOMAIN=http://localhost:9000
```

#### 🐛 Bugs đã fix

| Lỗi                     | Nguyên nhân                                      | Cách fix                                                  |
| ----------------------- | ------------------------------------------------ | --------------------------------------------------------- |
| `NoSuchBucket`          | Bucket name trong `.env` không khớp với MinIO    | Đồng bộ `S3BucketName` với bucket thực tế trên MinIO      |
| `MissingRegion`         | Biến môi trường không đọc đúng                   | Kiểm tra tên biến trong `.env` khớp với `os.Getenv()`     |
| `image: unknown format` | File upload không phải ảnh hợp lệ (JPEG/PNG/GIF) | Import `_ "image/jpeg"`, `_ "image/png"`, `_ "image/gif"` |

#### 📝 API Endpoints

| Method | Endpoint        | Mô tả                  | Auth               |
| ------ | --------------- | ---------------------- | ------------------ |
| POST   | `/upload`       | Upload ảnh             | No                 |
| POST   | `/users`        | Đăng ký user           | No                 |
| POST   | `/authenticate` | Đăng nhập              | No                 |
| GET    | `/profile`      | Lấy thông tin user     | Yes (Bearer Token) |
| PUT    | `/users/:id`    | Cập nhật user          | No                 |
| DELETE | `/users/:id`    | Xóa user (soft delete) | No                 |
| POST   | `/products`     | Tạo product            | No                 |
| GET    | `/products`     | Danh sách products     | No                 |
| POST   | `/sellers`      | Tạo seller             | No                 |

---

### v1.0.0 - Initial Release

#### ✅ Core Features

- Clean Architecture (Transport → Business → Storage)
- JWT Authentication
- Password hashing với bcrypt
- Soft delete pattern
- Pagination support
- Error handling chuẩn hóa

#### 🏗️ Architecture Pattern

```text
┌─────────────────────────────────────────────────────────┐
│                      Transport (Gin)                     │
│   - Nhận HTTP request                                   │
│   - Validate input                                       │
│   - Gọi Business layer                                   │
│   - Trả HTTP response                                    │
└─────────────────────────────────────────────────────────┘
                            │
                            ▼
┌─────────────────────────────────────────────────────────┐
│                      Business Logic                      │
│   - Xử lý nghiệp vụ                                     │
│   - Validate business rules                              │
│   - Gọi Storage qua Interface                           │
└─────────────────────────────────────────────────────────┘
                            │
                            ▼
┌─────────────────────────────────────────────────────────┐
│                    Storage (GORM/MySQL)                  │
│   - CRUD operations                                      │
│   - Query database                                       │
└─────────────────────────────────────────────────────────┘
```

#### 🔐 Authentication Flow

```text
1. POST /authenticate → Nhận email + password
2. Validate credentials → So sánh bcrypt hash
3. Generate JWT token → Trả về cho client
4. Client gửi: Authorization: Bearer <token>
5. Middleware validate token → Inject user vào context
```

---

## TODO / Roadmap

- [ ] Thêm role-based access control (RBAC)
- [ ] Implement cart module
- [ ] Implement order module
- [ ] Implement inventory module
- [ ] Implement variant module (product variants)
- [ ] Add Redis caching
- [ ] Add rate limiting
- [ ] Swagger/OpenAPI documentation
- [ ] Unit tests & integration tests
- [ ] Docker & docker-compose setup

---

## Lưu ý

- File `app_respose.go` đặt tên sai chính tả (nên đổi thành `app_response.go`)
- Package name trong `module/upload/business/upload.go` là `uploadgin` (nên đổi thành `uploadbusiness` cho nhất quán)
