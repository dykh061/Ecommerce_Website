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
│  ├─ app_respose.go
│  ├─ paging.go
│  └─ sql_model.go
├─ component/
│  ├─ appctx/
│  │  └─ app_context.go
│  └─ hasher/
│     └─ hasher.go
├─ module/
│  ├─ auth/ (trống)
│  ├─ cart/ (trống)
│  ├─ inventory/ (trống)
│  ├─ order/ (trống)
│  ├─ variant/ (trống)
│  ├─ product/
│  │  ├─ business/
│  │  │  ├─ create_product.go
│  │  │  ├─ delete_product.go
│  │  │  └─ list_product.go
│  │  ├─ model/
│  │  │  ├─ filter.go
│  │  │  └─ product.go
│  │  ├─ storage/
│  │  │  ├─ create.go
│  │  │  ├─ delete.go
│  │  │  ├─ find.go
│  │  │  ├─ list.go
│  │  │  └─ storage.go
│  │  └─ transport/
│  │     └─ gin/
│  │        ├─ create_product.go
│  │        ├─ delete_product.go
│  │        └─ list_product.go
│  ├─ seller/
│  │  ├─ business/
│  │  │  └─ create_seller.go
│  │  ├─ model/
│  │  │  └─ seller.go
│  │  ├─ storage/
│  │  │  ├─ create.go
│  │  │  └─ storage.go
│  │  └─ transport/
│  │     └─ gin/
│  │        └─ create_seller.go
│  └─ user/
│     ├─ business/
│     │  ├─ create_user.go
│     │  ├─ delete_user.go
│     │  ├─ find_user.go
│     │  ├─ password_hasher.go
│     │  └─ update_user.go
│     ├─ model/
│     │  └─ user.go
│     ├─ storage/
│     │  ├─ create.go
│     │  ├─ delete.go
│     │  ├─ find.go
│     │  ├─ storage.go
│     │  └─ update.go
│     └─ transport/
│        └─ gin/
│           ├─ create_user.go
│           ├─ delete_user.go
│           ├─ find_user.go
│           └─ update_user.go
└─ common note: file app_respose.go hiện đặt tên sai chính tả "respose" (nếu cần có thể đổi thành "response").
```

## Mô tả nhanh

- **common/**: Các kiểu phản hồi chung (`app_respose.go`), phân trang (`paging.go`), và model SQL cơ bản (`sql_model.go`).
- **component/**: `appctx` quản lý context/kết nối DB; `hasher` là bộ mã hoá mật khẩu (bcrypt).
- **module/**: Các domain chính (user, seller, product đã có code; auth/cart/inventory/order/variant còn trống).
- **main.go**: Khởi tạo ứng dụng, kết nối DB, đăng ký router Gin.
- `.env`, `.gitignore`, `go.mod`, `go.sum`: Cấu hình môi trường, bỏ qua file nhạy cảm, khai báo module và checksum Go.
