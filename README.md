# Flutter-App

my flutter 101 app project go go

## เริ่มต้นใช้งาน

```bash
# 1. ติดตั้ง dependencies
flutter pub get

# 2. ตั้งค่า environment
cp .env.example .env   # แล้วแก้ค่าใน .env ให้ถูกต้อง

# 3. รันแอป
flutter run
```

## Environment

ค่าลับต่าง ๆ เก็บในไฟล์ `.env` (ไม่ถูก commit)
ดูตัวอย่างตัวแปรที่ต้องตั้งค่าได้ที่ `.env.example`

> หมายเหตุ: หากต้องการอ่านค่า `.env` ในแอป Flutter แนะนำแพ็กเกจ
> [`flutter_dotenv`](https://pub.dev/packages/flutter_dotenv)
> และอย่าลืมเพิ่ม `.env` ใน `assets:` ของ `pubspec.yaml`
