# Inventory CLI

Aplikasi Command Line Interface (CLI) untuk mengelola sistem inventory secara lengkap. Dibangun menggunakan Go dengan framework Cobra untuk handling command dan PostgreSQL sebagai database.

## Deskripsi

Inventory CLI adalah aplikasi berbasis terminal yang memungkinkan pengguna untuk melakukan manajemen inventory secara komprehensif dengan tiga modul utama:

- **Category Management**: Mengelola kategori barang (CRUD operations)
- **Items Management**: Mengelola data barang/item inventory
- **Reports**: Menghasilkan laporan inventory

Aplikasi ini menggunakan arsitektur clean code dengan pemisahan layer repository, service, dan handler untuk memudahkan maintenance dan scalability.

## Cara Penggunaan

### Fitur Category Management

#### 1. Menambahkan Kategori Baru

Menambahkan kategori baru ke dalam database.

**Command:**

```bash
go run . add-category --name="Nama Kategori" --desc="Deskripsi Kategori"
```

**Contoh:**

```bash
go run . add-category --name="Elektronik" --desc="Kategori untuk barang elektronik"
go run . add-category -n "Furniture" -d "Kategori untuk perabotan rumah tangga"
```

**Parameter:**

- `--name` atau `-n` (required): Nama kategori
- `--desc` atau `-d` (required): Deskripsi kategori

**Output:**

```
Kategori berhasil dibuat.
```

#### 2. Menampilkan Semua Kategori

Mengambil dan menampilkan semua data kategori yang ada di database.

**Command:**

```bash
go run . list-category
```

**Output:**

```
ID: 1 | Name: Elektronik | Description: Kategori untuk barang elektronik
ID: 2 | Name: Furniture | Description: Kategori untuk perabotan rumah tangga
```

#### 3. Melihat Detail Kategori

Menampilkan detail kategori berdasarkan ID tertentu.

**Command:**

```bash
go run . detail-category --id=<ID>
```

**Contoh:**

```bash
go run . detail-category --id=1
go run . detail-category -i 3
```

**Parameter:**

- `--id` atau `-i` (required): ID kategori yang ingin dilihat

**Output:**

```
ID: 1
Name: Elektronik
Description: Kategori untuk barang elektronik
Created At: 2024-01-15 10:30:00
```

### Melihat Help

Untuk melihat bantuan dan daftar command yang tersedia:

**Command:**

```bash
go run . --help
go run . add-category --help
```

---

### Fitur Item Management

#### 1. Menambahkan Item Baru
