# Aplikasi Manajemen Tenant

Aplikasi ini adalah program sederhana untuk manajemen penyewaan (tenant) yang dilakukan di suatu tempat. Program ini memungkinkan pengguna untuk menambah, mengubah, menghapus, serta melihat daftar tenant dan transaksi yang terjadi.

## Fitur

1. **Tambah Tenant :** Menambahkan data tenant baru ke dalam sistem.
2. **Ubah Tenant :** Mengubah data tenant yang sudah ada.
3. **Hapus Tenant :** Menghapus data tenant dari sistem.
4. **Tambah Transaksi :** Mencatat transaksi baru yang dilakukan oleh tenant.
5. **Tampilkan Transaksi :** Melihat daftar transaksi yang telah dilakukan beserta informasi tenant terkait.
6. **Urutkan Transaksi :** Mengurutkan transaksi secara menaik dan menurun

## Cara Penggunaan

1. **Menambah Tenant**: Pilih opsi "Tambah Tenant" dari menu dan masukkan data tenant yang baru.
2. **Mengubah Tenant**: Pilih opsi "Ubah Tenant" dari menu dan ikuti instruksi untuk mengubah data tenant yang sudah ada.
3. **Menghapus Tenant**: Pilih opsi "Hapus Tenant" dari menu dan ikuti instruksi untuk menghapus data tenant.
4. **Menambah Transaksi**: Pilih opsi "Tambah Transaksi" dari menu dan masukkan data transaksi baru untuk tenant yang dipilih.
5. **Melihat Transaksi**: Pilih opsi "Tampilkan Transaksi" dari menu untuk melihat daftar transaksi yang telah dicatat.
6. **Urutkan Transaksi :** Pilih opsi "Urutkan Transaksi" dari menu dan pilih nomor sesuai keinginan

## Catatan

- Program ini menggunakan fungsi `flushScanner()` untuk membersihkan input dari scanner setelah digunakan.
- Fungsi `cekTenant()` digunakan untuk memeriksa apakah tenant sudah ada dalam daftar atau belum.
- Transaksi dicatat dengan menghitung total uang yang dibayarkan oleh tenant dan menghitung keuntungan yang didapat (sebesar 25% dari total uang yang dibayarkan).
- Daftar tenant dan transaksi dapat diurutkan berdasarkan jumlah transaksi secara menaik atau menurun.
