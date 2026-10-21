# Article Service - API Specification
**Project:** Sharing Vision Article API Submission 2026

**Base URL VPS:** `https://submission.metrolabs.site`

**Base URL lokal:** `http://localhost:8081`  
**Format:** JSON  
**Service:** `article-service` — Go + Gin + MySQL

## Membuat artikel baru
Endpoint: `POST /api/article`

Request Body : 

```json
{
  "title": "",
  "content": "",
  "category":"",
  "status": "",
}
```

Response Body (Success) : 
```json
{
  
}
```

## Menampilkan article 

Endpoint: `GET /article/:limit/:offset`

Path Parameters:
- `limit` (integer): Jumlah maksimum artikel yang ditampilkan dalam satu permintaan.
- `offset` (integer): Jumlah artikel yang dilewati sebelum mengambil data.

Contoh Request:
```http GET /article/10/0 ```

Request tersebut mengambil maksimal 10 artikel, dimulai dari data pertama (`offset = 0`).

**Response Body (Success):**
```json
[
  {
    "title": "",
    "content": "",
    "category": "",
    "status": ""
  }
]
```
## Menampilkan article dengan id

Endpoint: `GET /article/:id`


Response Body (Success) :
```json
[{
  "title": "",
  "content": "",
  "category":"",
  "status":""
}]
```

## Merubah data article dengan id
Endpoint: `PUT /article/:id`

Response Body(success) : 

```json
{
  "title": "",
  "content": "",
  "category":"",
  "status":"" 
}
```

## Menghapus data article dengan id
Endpoint: `DELETE /article/:id`

Response Body(success) :

```json
{
  "message": "data berhasil dihapus"
}
```

## Model Data

| Field | Tipe JSON | Wajib pada create | Aturan |
|---|---|---:|---|
| `id` | integer | Tidak | ID unik database |
| `title` | string | Ya | Minimal 20 karakter |
| `content` | string | Ya | Minimal 200 karakter |
| `category` | string | Ya | Minimal 3 karakter |
| `status` | string | Ya | `publish`, `draft`, atau `thrash` |
| `created_date` | string date-time | Tidak | Waktu dibuat server |
| `updated_date` | string date-time | Tidak | Waktu diperbarui server |

Status `thrash` mengikuti ejaan pada PDF. Gunakan konsisten pada API dan database.

## Validasi

Untuk create dan update:
- `title` wajib, minimal 20 karakter.
- `content` wajib, minimal 200 karakter.
- `category` wajib, minimal 3 karakter.
- `status` wajib, hanya `publish`, `draft`, atau `thrash`.
- String kosong atau hanya spasi dianggap kosong (rekomendasi).
- `limit` bilangan bulat positif dan `offset` bilangan bulat nol atau lebih (rekomendasi).
- `id` bilangan bulat positif (rekomendasi).
- Untuk `PATCH`, rekomendasi: gabungkan field parsial dengan data tersimpan, lalu validasi kondisi akhir artikel.




