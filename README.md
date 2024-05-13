## Uy vazifa

1. ### `GET` buyrug'ini yuborganingizda joriy vaqtni `RFC 3339` formatida qaytaradigan kichik web-server yozing (masalan `/time/RFC3339` endpoint)
2. ### Vaqtni `JSON` sifatida qaytarish imkoniyatini qo'shing. `JSON` yoki matn qaytarilishini nazorat qilish uchun Qabul qilish sarlavhasidan foydalaning (standart matn uchun). `JSON` quyidagi tarzda tuzilishi kerak:
```json
{
  "day_of_week": "Monday",
  "day_of_month": 10,
  "month": "April",
  "year": 2023,
  "hour": 20,
  "minute": 15,
  "second": 20
}
```