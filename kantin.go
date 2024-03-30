package main

import (
	"bufio"
	"fmt"
	"os"
)

const NMAX = 1000

type tenant struct {
	nama       string
	totalUang  int
	nTransaksi int
	untung     int
}

type sewa struct {
	tab [NMAX]tenant
	n   int
}

func flushScanner() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
}

func menu(sewaTab *sewa) {
	var pilihan int

	fmt.Println()
	fmt.Println("----====[ MENU ]====----")
	fmt.Println("1. Lihat Tenant")
	fmt.Println("2. Tambah Tenant")
	fmt.Println("3. Ubah Tenant")
	fmt.Println("4. Hapus Tenant")
	fmt.Println("5. Tambah Transaksi")
	fmt.Println("6. Tampilkan Transaksi")
	fmt.Println("7. Keluar")
	fmt.Println("------------------------")
	fmt.Print("Pilihan: ")
	fmt.Scan(&pilihan)
	flushScanner()

	if pilihan == 1 {
		tampilkanTenant(*sewaTab)
		menu(sewaTab)
	} else if pilihan == 2 {
		tambahTenant(sewaTab)
	} else if pilihan == 3 {
		ubahTenant(sewaTab)
	} else if pilihan == 4 {
		hapusTenant(sewaTab)
	} else if pilihan == 5 {
		tambahTransaksi(sewaTab)
	} else if pilihan == 6 {
		tampilkanTransaksi(sewaTab)
	} else if pilihan == 7 {
		fmt.Println("Terima Kasih.")
	} else {
		fmt.Println("Pilihan Tidak Valid")
		menu(sewaTab)
	}
}

func ubahTenant(sewaTab *sewa) {
	var temp, temp2 tenant
	var pilihan int

	for pilihan != 3 {
		fmt.Println()
		fmt.Println("----====[ UBAH TENANT ]====----")
		fmt.Println("1. Ubah Tenant")
		fmt.Println("2. Lihat Tenant")
		fmt.Println("3. Kembali Ke Menu")
		fmt.Println("--------------------------------")
		fmt.Print("Pilihan: ")
		fmt.Scan(&pilihan)
		flushScanner()

		if pilihan == 1 {
			fmt.Print("Masukkan nama tenant: ")
			fmt.Scan(&temp.nama)
			id := cekTenant(sewaTab, temp)

			if id != -1 {
				fmt.Printf("Masukkan nama tenant baru %s: ", temp.nama)
				fmt.Scan(&temp2.nama)
				if cekTenant(sewaTab, temp2) != -1 {
					fmt.Println("Nama tenant sudah ada.")
				} else {
					sewaTab.tab[id].nama = temp2.nama
					fmt.Printf("Tenant %s berhasil diubah menjadi %s\n", temp.nama, temp2.nama)
				}
			} else {
				fmt.Println("Nama tenant tidak ditemukan.")
			}
		} else if pilihan == 2 {
			tampilkanTenant(*sewaTab)
		} else if pilihan == 3 {
			menu(sewaTab)
		} else {
			fmt.Println("Pilihan tidak valid.")
		}
	}
}

func tambahTenant(sewaTab *sewa) {
	var temp tenant
	var pilihan int

	for pilihan != 3 {
		fmt.Println()
		fmt.Println("----====[ TAMBAH TENANT ]====----")
		fmt.Println("1. Tambah Tenant")
		fmt.Println("2. Lihat Tenant")
		fmt.Println("3. Kembali Ke Menu")
		fmt.Println("---------------------------------")
		fmt.Print("Pilihan: ")
		fmt.Scan(&pilihan)
		flushScanner()

		if pilihan == 1 {
			fmt.Print("Masukkan Nama Tenant: ")
			fmt.Scan(&temp.nama)
			if cekTenant(sewaTab, temp) != -1 {
				fmt.Println("Nama Tenant Sudah Ada.")
			} else {
				sewaTab.tab[sewaTab.n] = temp
				sewaTab.n++
				fmt.Printf("Tenant %s Berhasil Ditambahkan\n", temp.nama)
			}

		} else if pilihan == 2 {
			tampilkanTenant(*sewaTab)
		} else if pilihan == 3 {
			menu(sewaTab)
		} else {
			fmt.Println("Pilihan Tidak Valid")
		}
	}
}

func cekTenant(sewaTab *sewa, t tenant) int {
	for i := 0; i < sewaTab.n; i++ {
		if sewaTab.tab[i].nama == t.nama {
			return i
		}
	}
	return -1
}

func hapusTenant(sewaTab *sewa) {
	var temp, clear tenant
	var pilihan int

	for pilihan != 3 {
		fmt.Println()
		fmt.Println("----====[ HAPUS TENANT ]====----")
		fmt.Println("1. Hapus tenant")
		fmt.Println("2. Lihat Tenant")
		fmt.Println("3. Kembali Ke Menu")
		fmt.Println("--------------------------------")
		fmt.Print("Pilihan: ")
		fmt.Scan(&pilihan)
		flushScanner()

		if pilihan == 1 {
			fmt.Print("Masukkan nama tenant: ")
			fmt.Scan(&temp.nama)
			id := cekTenant(sewaTab, temp)

			if id != -1 {
				for i := id; i < sewaTab.n-1; i++ {
					sewaTab.tab[i] = sewaTab.tab[i+1]
				}
				sewaTab.tab[sewaTab.n-1] = clear
				sewaTab.n--
				fmt.Printf("Tenant %s berhasil dihapus\n", temp.nama)
			} else {
				fmt.Println("Nama tenant tidak ditemukan.")
			}
		} else if pilihan == 2 {
			tampilkanTenant(*sewaTab)
		} else if pilihan == 3 {
			menu(sewaTab)
		} else {
			fmt.Println("Pilihan tidak valid.")
		}
	}
}

func tampilkanTenant(sewaTab sewa) {
	fmt.Println()
	fmt.Println("----====[ TENANT TERDAFTAR ]====----")
	j := 1
	for i := 0; i < sewaTab.n; i++ {
		if sewaTab.tab[i].nama != "" {
			fmt.Printf("%d. %s\n", j, sewaTab.tab[i].nama)
			j++
		}
	}
	fmt.Println("------------------------------------")
}

func tambahTransaksi(sewaTab *sewa) {
	var pilihan int

	for pilihan != 3 {
		fmt.Println()
		fmt.Println("----====[ TAMBAH TRANSAKSI ]====----")
		fmt.Println("1. Tambah transaksi")
		fmt.Println("2. Lihat Tenant")
		fmt.Println("3. Kembali ke menu")
		fmt.Println("------------------------------------")
		fmt.Print("Pilihan: ")
		fmt.Scan(&pilihan)
		flushScanner()

		if pilihan == 1 {
			var temp tenant

			fmt.Print("Masukkan nama tenant: ")
			fmt.Scan(&temp.nama)
			id := cekTenant(sewaTab, temp)

			if id != -1 {
				temp.totalUang = -1
				for temp.totalUang != 0 {
					fmt.Printf("Masukkan besar transaksi tenant %s (0 untuk keluar dari transaksi): ", temp.nama)
					fmt.Scan(&temp.totalUang)
					flushScanner()
					sewaTab.tab[id].totalUang += temp.totalUang
					sewaTab.tab[id].untung += (temp.totalUang * 25) / 100
					sewaTab.tab[id].nTransaksi++
				}
				sewaTab.tab[id].nTransaksi--
			} else {
				fmt.Println("Nama tenant tidak ditemukan.")
			}
		} else if pilihan == 2 {
			tampilkanTenant(*sewaTab)
		} else if pilihan == 3 {
			menu(sewaTab)
		} else {
			fmt.Println("Pilihan tidak valid.")
		}
	}
}

func urutkanTransaksi(sewaTab *sewa, mode string) {
	if mode == "asc" {
		for i := 0; i < sewaTab.n; i++ {
			for j := i + 1; j < sewaTab.n; j++ {
				if sewaTab.tab[i].nTransaksi > sewaTab.tab[j].nTransaksi {
					temp := sewaTab.tab[i]
					sewaTab.tab[i] = sewaTab.tab[j]
					sewaTab.tab[j] = temp
				}
			}
		}
	}
	if mode == "desc" {
		for i := 0; i < sewaTab.n; i++ {
			for j := i + 1; j < sewaTab.n; j++ {
				if sewaTab.tab[i].nTransaksi < sewaTab.tab[j].nTransaksi {
					temp := sewaTab.tab[i]
					sewaTab.tab[i] = sewaTab.tab[j]
					sewaTab.tab[j] = temp
				}
			}
		}
	}
}

func tampilkanTransaksi(sewaTab *sewa) {
	var pilihan int

	for pilihan != 2 {
		fmt.Println()
		fmt.Println("----====[ MENU LIHAT TRANSAKSI ]====----")
		fmt.Println("1. Lihat transaksi")
		fmt.Println("2. Kembali ke menu")
		fmt.Println("----------------------------------------")
		fmt.Print("Pilihan: ")
		fmt.Scan(&pilihan)
		flushScanner()

		if pilihan == 1 {
			pilihan = 0

			for pilihan != 1 {
				fmt.Println()
				fmt.Println("----====[ DAFTAR TRANSAKSI ]====----")
				fmt.Println()
				fmt.Println("NO (ID) | NAMA (N) | TOTAL UANG (U) | JUMLAH TRANSAKSI (T) | KEUNTUNGAN (K)")
				fmt.Println("---------------------------------------------------------------------------")
				j := 1
				for i := 0; i < sewaTab.n; i++ {
					if sewaTab.tab[i].nama != "" {
						fmt.Printf("%d (ID) | %s (N) | %d (U) | %d (T) | %d (K)\n", j, sewaTab.tab[i].nama, sewaTab.tab[i].totalUang, sewaTab.tab[i].nTransaksi, sewaTab.tab[i].untung)
						j++
					}
				}
				fmt.Println("---------------------------------------------------------------------------")
				fmt.Println("NO (ID) | NAMA (N) | TOTAL UANG (U) | JUMLAH TRANSAKSI (T) | KEUNTUNGAN (K)")
				fmt.Println()
				fmt.Println("----====[ MENU DAFTAR TRANSAKSI ]====----")
				fmt.Println("1. Kembali ke menu lihat transaksi")
				fmt.Println("2. Urutkan jumlah transaksi secara menaik")
				fmt.Println("3. Urutkan jumlah transaksi secara menurun")
				fmt.Println("------------------------")
				fmt.Print("Pilihan: ")
				fmt.Scan(&pilihan)
				flushScanner()

				if pilihan == 1 {
					break
				} else if pilihan == 2 {
					urutkanTransaksi(sewaTab, "asc")
				} else if pilihan == 3 {
					urutkanTransaksi(sewaTab, "desc")
				} else {
					fmt.Println("Pilihan tidak valid.")
				}
			}
		} else if pilihan == 2 {
			menu(sewaTab)
		} else {
			fmt.Println("Pilihan tidak valid.")
		}
	}
}

func main() {
	var sewaTab sewa
	menu(&sewaTab)
}