# Inventory CLI

Aplikasi Command Line Interface (CLI) untuk mengelola sistem inventory secara lengkap. Dibangun menggunakan Go dengan framework Cobra untuk handling command dan PostgreSQL sebagai database.

## Deskripsi

Inventory CLI adalah aplikasi berbasis terminal yang memungkinkan pengguna untuk melakukan manajemen inventory secara komprehensif dengan tiga modul utama:

- **Category Management**: Mengelola kategori barang (CRUD operations)
- **Items Management**: Mengelola data barang/item inventory (CRUD operations)
- **Reports**: Menghasilkan laporan inventory termasuk analisis investasi dan depresiasi

## Cara Penggunaan

### Melihat Help

Untuk melihat bantuan dan daftar command yang tersedia:

**Command:**

```bash
go run . --help
go run . add-category --help
go run . list-items --help
```

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

---

### Fitur Item Management

#### 1. Menampilkan Semua Barang/item

Mengambil dan menampilkan semua data barang dari database.

**Command:**

```bash
go run . list-items
```

**Output:**

```
ID: 1 | Name: Laptop Dell | Price: Rp 15.000.000 | Category: Elektronik
ID: 2 | Name: Meja Kerja | Price: Rp 2.500.000 | Category: Furniture
```

#### 2. Melihat Detail Barang/item

Menampilkan detail barang berdasarkan ID tertentu.

**Command:**

```bash
go run . detail-item --id=<ID>
```

**Contoh:**

```bash
go run . detail-item --id=1
go run . detail-item -i 2
```

**Parameter:**

- `--id` atau `-i` (required): ID item yang ingin dilihat

**Output:**

```
ID: 1
Name: Laptop Dell
Price: Rp 15.000.000
Category: Elektronik
Purchase Date: 2023-06-15
Usage Days: 180
```

#### 3. Mencari Barang/item Berdasarkan Nama

Mencari barang dengan keyword tertentu.

**Command:**

```bash
go run . search-item --name="keyword"
```

**Contoh:**

```bash
go run . search-item --name="laptop"
go run . search-item -n "meja"
```

**Parameter:**

- `--name` atau `-n` (required): Keyword nama item yang dicari

**Output:**

```
Hasil pencarian untuk "laptop":
ID: 1 | Name: Laptop Dell | Price: Rp 15.000.000 | Category: Elektronik
ID: 3 | Name: Laptop HP | Price: Rp 12.000.000 | Category: Elektronik
```

#### 4. Update Data Barang/item

Mengubah/update data barang yang sudah ada.

**Command:**

```bash
go run . update-item --id=<ID> [options]
```

**Contoh:**

```bash
go run . update-item --id=1 --name="Laptop Dell XPS" --price=18000000
go run . update-item -i 2 -n "Meja Gaming" -p 3500000 -c 2
go run . update-item --id=1 --date="2024-01-15" --usage=90
```

**Parameter:**

- `--id` atau `-i` (required): ID item yang akan diupdate
- `--name` atau `-n` (optional): Nama baru item
- `--price` atau `-p` (optional): Harga baru item
- `--category` atau `-c` (optional): Category ID baru
- `--date` atau `-d` (optional): Tanggal pembelian baru (format: YYYY-MM-DD)
- `--usage` atau `-u` (optional): Jumlah hari pemakaian baru

**Catatan:** Parameter yang tidak diisi akan tetap menggunakan nilai lama.

#### 5. Barang/item yang Perlu Diganti

Menampilkan daftar barang yang sudah digunakan lebih dari 100 hari dan perlu diganti.

**Command:**

```bash
go run . need-replacement
```

**Output:**

```
Barang yang perlu diganti (usage > 100 hari):

ID: 1 | Name: Laptop Dell | Usage Days: 180 | Category: Elektronik
ID: 5 | Name: Printer Canon | Usage Days: 150 | Category: Elektronik

Total: 2 barang perlu diganti
```

---

### Fitur Laporan Investasi

Fitur ini menghitung nilai investasi barang dengan mempertimbangkan depresiasi menggunakan **metode saldo menurun (declining balance method)** dengan tingkat depresiasi **20% per tahun**.

#### Formula Depresiasi:

```
Nilai Setelah Depresiasi = Harga Beli × (1 - Tingkat Depresiasi)^Tahun Pemakaian
Nilai Setelah Depresiasi = Harga Beli × (0.8)^(Usage Days / 365)
Depresiasi = Harga Beli - Nilai Setelah Depresiasi
```

#### 1. Ringkasan Total Investasi

Menampilkan total nilai investasi seluruh barang setelah perhitungan depresiasi sebesar 20% per tahun.

**Command:**

```bash
go run . investment-summary
```

**Output:**

```
=== LAPORAN INVESTASI DAN DEPRESIASI ===

Total Nilai Investasi Awal    : Rp 50.000.000
Total Depresiasi               : Rp 8.500.000
Total Nilai Setelah Depresiasi : Rp 41.500.000

Jumlah Barang: 10 items
```

**Keterangan:**

- **Total Nilai Investasi Awal**: Jumlah seluruh harga beli barang
- **Total Depresiasi**: Total penyusutan nilai dari seluruh barang
- **Total Nilai Setelah Depresiasi**: Nilai investasi saat ini setelah dikurangi depresiasi

#### 2. Detail Investasi per Barang

Menampilkan nilai investasi dan depresiasi untuk barang tertentu berdasarkan ID.

**Command:**

```bash
go run . investment-detail --id=<ID>
```

**Contoh:**

```bash
go run . investment-detail --id=1
go run . investment-detail -i 3
```

**Parameter:**

- `--id` atau `-i` (required): ID item yang ingin dilihat detail investasinya

**Output:**

```
=== DETAIL INVESTASI BARANG ===

Nama Barang          : Laptop Dell
Kategori             : Elektronik
Tanggal Pembelian    : 2023-06-15
Hari Pemakaian       : 180 hari (0.49 tahun)

Harga Beli           : Rp 15.000.000
Tingkat Depresiasi   : 20% per tahun
Depresiasi           : Rp 2.850.000
Nilai Setelah        : Rp 12.150.000
Depresiasi

Persentase Penyusutan: 19.0%
```

**Keterangan:**

- **Harga Beli**: Nilai awal investasi barang
- **Tingkat Depresiasi**: Rate depresiasi yang digunakan (20% per tahun)
- **Depresiasi**: Jumlah nilai yang sudah berkurang
- **Nilai Setelah Depresiasi**: Nilai barang saat ini
- **Persentase Penyusutan**: Persentase nilai yang sudah hilang dari harga awal

---

## Link Demo dan Penjelasan Program

---

- [Demo dan Penjelasan Program](https://drive.google.com/file/d/1Li3kgZjVfE4doOPGgr9oOSRQvLGAz1Y1/view?usp=sharing)
