# Product API (Go backend)

REST API สำหรับ CRUD สินค้า เขียนด้วย Go มาตรฐาน (`net/http`) เก็บข้อมูลในหน่วยความจำ

## โครงสร้างไฟล์

| ไฟล์ | หน้าที่ |
|------|--------|
| `main.go` | จุดเริ่มโปรแกรม + ตั้งเส้นทาง (routing) + CORS |
| `handlers.go` | ฟังก์ชันรับ HTTP request แต่ละแบบ แปลง JSON |
| `product.go` | โมเดล `Product` + ที่เก็บข้อมูล `ProductStore` |
| `product_test.go` | เทสต์ของ store |

## วิธีรัน

```bash
cd backend
go run .
```
เซิร์ฟเวอร์จะทำงานที่ http://localhost:8080

## วิธีเทสต์

```bash
cd backend
go test ./...
```

## Endpoints

| Method | Path | ทำอะไร |
|--------|------|--------|
| GET | `/products` | ดูสินค้าทั้งหมด |
| GET | `/products/{id}` | ดูสินค้า 1 ชิ้น |
| POST | `/products` | เพิ่มสินค้าใหม่ |
| PUT | `/products/{id}` | แก้ไขสินค้า |
| DELETE | `/products/{id}` | ลบสินค้า |

## ตัวอย่างการเรียก (curl)

```bash
# ดูทั้งหมด
curl http://localhost:8080/products

# เพิ่มใหม่
curl -X POST http://localhost:8080/products \
  -H "Content-Type: application/json" \
  -d "{\"name\":\"นม\",\"price\":25,\"stock\":50}"

# แก้ไข id 1
curl -X PUT http://localhost:8080/products/1 \
  -H "Content-Type: application/json" \
  -d "{\"name\":\"กาแฟเย็น\",\"price\":50,\"stock\":90}"

# ลบ id 1
curl -X DELETE http://localhost:8080/products/1
```

## ข้อมูลสินค้า (JSON)

```json
{
  "id": 1,
  "name": "กาแฟ",
  "price": 45,
  "stock": 100
}
```

> หมายเหตุ: ข้อมูลเก็บใน RAM จะหายเมื่อปิดเซิร์ฟเวอร์ — ขั้นต่อไปค่อยต่อ database จริง
