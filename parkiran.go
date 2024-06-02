package main

import "fmt"

//TUBES = Parkir
//Requirements Tubes
//Struct Array = Line 16
//Procedure = Line 95 ( 43 juga sih )
//Function = Line 274
//Sequential = Line 274
//Binary = Line 322
//SelectionSort = Line 232
//InsertionSort = Line 307

//struct array
type DataAdmin struct {
	Username string
	Password string
}

type PetugasParkir struct {
	ID       int
	Nama     string
	Username string
	Password string
}

type TransaksiParkir struct {
	JumlahKendaraan int
	TipeKendaraan   string
	WaktuMasuk      float32
	WaktuKeluar     float32
	BiayaParkir     float32
	Plat            string
}

const tPetugas = 100
const tParkir = 1000

type Petugas [tPetugas]PetugasParkir
type Transaksi [tParkir]TransaksiParkir

func main() {
	var pilihan int
	var PP Petugas
	var TP Transaksi
	var jumlahPetugas int
	var jumlahKendaraan int

	fmt.Println(" .----------------.  .----------------.  .----------------.  .----------------.  .----------------.  .----------------.  .----------------. ")
	fmt.Println("| .--------------. || .--------------. || .--------------. || .--------------. || .--------------. || .--------------. || .--------------. |")
	fmt.Println("| | _____  _____ | || |  _________   | || |   _____      | || |     ______   | || |     ____     | || | ____    ____ | || |  _________   | |")
	fmt.Println("| ||_   _||_   _|| || | |_   ___  |  | || |  |_   _|     | || |   .'___  |  | || |   .'    `.   | || ||_   \\  /   _|| || | |_   ___  |  | |")
	fmt.Println("| |  | | /\\ | |  | || |   | |_  \\_|  | || |    | |       | || |  / .'   \\_|  | || |  /  .--.  \\  | || |  |   \\/   |  | || |   | |_  \\_|  | |")
	fmt.Println("| |  | |/  \\| |  | || |   |  _|  _   | || |    | |   _   | || |  | |         | || |  | |    | |  | || |  | |\\  /| |  | || |   |  _|  _   | |")
	fmt.Println("| |  |   /\\   |  | || |  _| |___/ |  | || |   _| |__/ |  | || |  \\ `.___.'\\  | || |  \\  `--'  /  | || | _| |_\\/_| |_ | || |  _| |___/ |  | |")
	fmt.Println("| |  |__/  \\__|  | || | |_________|  | || |  |________|  | || |   `._____.'  | || |   `.____.'   | || ||_____||_____|| || | |_________|  | |")
	fmt.Println("| |              | || |              | || |              | || |              | || |              | || |              | || |              | |")
	fmt.Println("| '--------------' || '--------------' || '--------------' || '--------------' || '--------------' || '--------------' || '--------------' |")
	fmt.Println(" '----------------'  '----------------'  '----------------'  '----------------'  '----------------'  '----------------'  '----------------' ")

	for pilihan != 3 {
		fmt.Println("\nPilih Role :")
		fmt.Println("1. Admin")
		fmt.Println("2. Petugas Parkir")
		fmt.Println("3. Keluar")
		fmt.Println("Masukan Pilihan : ")
		fmt.Scanln(&pilihan)

		switch pilihan {
		case 1:
			admin(&PP, &jumlahPetugas)
		case 2:
			petugas(&PP, &TP, &jumlahKendaraan, jumlahPetugas)
		case 3:
			fmt.Println(" .----------------.  .----------------.  .----------------.  .----------------.  .----------------.  .----------------.  .----------------. ")
			fmt.Println("| .--------------. || .--------------. || .--------------. || .--------------. || .--------------. || .--------------. || .--------------. |")
			fmt.Println("| |  _________   | || |  ____  ____  | || |      __      | || | ____  _____  | || |  ___  ____   | || |    _______   | || |    _______   | |")
			fmt.Println("| | |  _   _  |  | || | |_   ||   _| | || |     /  \\     | || ||_   \\|_   _| | || | |_  ||_  _|  | || |   /  ___  |  | || |   /  ___  |  | |")
			fmt.Println("| | |_/ | | \\_|  | || |   | |__| |   | || |    / /\\ \\    | || |  |   \\ | |   | || |   | |_/ /    | || |  |  (__ \\_|  | || |  |  (__ \\_|  | |")
			fmt.Println("| |     | |      | || |   |  __  |   | || |   / ____ \\   | || |  | |\\ \\| |   | || |   |  __'.    | || |   '.___`-.   | || |   '.___`-.   | |")
			fmt.Println("| |    _| |_     | || |  _| |  | |_  | || | _/ /    \\ \\_ | || | _| |\\_\\ |_  | || |  _| |  \\ \\_  | || |  |`\\____) |  | || |  |`\\____) |  | |")
			fmt.Println("| |   |_____|    | || | |____||____| | || ||____|  |____|| || ||_____|\\____| | || | |____||____| | || |  |_______.'  | || |  |_______.'  | |")
			fmt.Println("| |              | || |              | || |              | || |              | || |              | || |              | || |              | |")
			fmt.Println("| '--------------' || '--------------' || '--------------' || '--------------' || '--------------' || '--------------' || '--------------' |")
			fmt.Println(" '----------------'  '----------------'  '----------------'  '----------------'  '----------------'  '----------------'  '----------------' ")

		default:
			fmt.Println("Pilihan tidak valid.")
		}
	}
}

//Prosedur
func admin(PP *Petugas, jumlahPetugas *int) {
	var userAdmin DataAdmin
	var exit bool

	exit = false

	for exit == false {
		fmt.Println("\nMasukan Username Dan Password Admin")
		fmt.Println("Ketik 'keluar' jika anda bukan admin atau kembali ke menu awal")
		fmt.Print("Username : ")
		fmt.Scanln(&userAdmin.Username)
		if userAdmin.Username == "keluar" || userAdmin.Username == "Keluar" {
			exit = true
		} else {
			fmt.Print("Password : ")
			fmt.Scanln(&userAdmin.Password)
			if userAdmin.Username == "Admin" && userAdmin.Password == "Admin" {
				fmt.Println("Selamat Datang Admin!")
				roleAdmin(PP, jumlahPetugas)
			} else {
				fmt.Println("Username atau Password salah, coba lagi.")
			}
		}
	}
}

func roleAdmin(PP *Petugas, jumlahPetugas *int) {
	var pilihan int
	var exit bool

	exit = false

	for !exit {
		fmt.Println("\nPilih Menu : ")
		fmt.Println("1. Tambah Petugas Parkir")
		fmt.Println("2. Cetak Daftar Petugas Parkir")
		fmt.Println("3. Edit Petugas Parkir")
		fmt.Println("4. Hapus Data Petugas Parkir")
		fmt.Println("5. Keluar (kembali ke menu pemilihan role)")
		fmt.Scanln(&pilihan)

		switch pilihan {
		case 1:
			tambahPetugasParkir(PP, jumlahPetugas)
		case 2:
			cetakDaftarPetugasParkir(PP, jumlahPetugas)
		case 3:
			editPetugas(PP, jumlahPetugas)
		case 4:
			hapusPetugas(PP, jumlahPetugas)
		case 5:
			exit = true
		default:
			fmt.Println("Pilihan Tidak Valid!")
		}
	}
}

func tambahPetugasParkir(petugasParkir *Petugas, jumlahPetugas *int) {
	var exit bool

	exit = false

	for exit == false {
		if *jumlahPetugas >= tPetugas {
			fmt.Println("Kapasitas petugas parkir sudah penuh.")
			exit = true
		} else {
			var id int
			validID := false
			for validID == false {
				fmt.Println("Masukkan Data Petugas Parkir:")
				fmt.Print("Masukan ID Petugas: ")
				fmt.Scanln(&id)
				if pengecekanIDBySequential(petugasParkir, jumlahPetugas, id) {
					fmt.Println("ID yang dimasukan sudah digunakan, silakan masukkan ID yang berbeda.")
				} else {
					validID = true
				}
			}
			petugasParkir[*jumlahPetugas].ID = id
			fmt.Print("Nama: ")
			fmt.Scanln(&petugasParkir[*jumlahPetugas].Nama)
			fmt.Print("Username: ")
			fmt.Scanln(&petugasParkir[*jumlahPetugas].Username)
			fmt.Print("Password: ")
			fmt.Scanln(&petugasParkir[*jumlahPetugas].Password)

			*jumlahPetugas++

			fmt.Println("Apakah ingin menambah petugas lagi? (yes/no)")
			var response string
			fmt.Scanln(&response)
			if response != "yes" {
				exit = true
			}
		}
	}
}

func pengecekanIDBySequential(petugasParkir *Petugas, jumlahPetugas *int, id int) bool {
	for i := 0; i < *jumlahPetugas; i++ {
		if petugasParkir[i].ID == id {
			return true
		}
	}
	return false
}

func cetakDaftarPetugasParkir(petugasParkir *Petugas, jumlahPetugas *int) {
	var pilihan int

	fmt.Println("\nPilih Urutan Cetak :")
	fmt.Println("1. Ascending")
	fmt.Println("2. Descending")
	fmt.Print("Masukan Pilihan : ")
	fmt.Scanln(&pilihan)

	if pilihan == 1 {
		selectionSort(petugasParkir, *jumlahPetugas, true)
		printPetugas(petugasParkir, *jumlahPetugas)
	} else if pilihan == 2 {
		selectionSort(petugasParkir, *jumlahPetugas, false)
		printPetugas(petugasParkir, *jumlahPetugas)
	} else {
		fmt.Println("Pilihan tidak valid.")
	}
}

func printPetugas(petugasParkir *Petugas, jumlahPetugas int) {
	fmt.Println("Daftar Petugas Parkir:")
	for i := 0; i < jumlahPetugas; i++ {
		fmt.Printf("ID: %d, Nama: %s, Username: %s\n", petugasParkir[i].ID, petugasParkir[i].Nama, petugasParkir[i].Username)
	}
}

//selectionSort
func selectionSort(petugasParkir *Petugas, jumlahPetugas int, ascending bool) {
	for i := 0; i < jumlahPetugas-1; i++ {
		idx := i
		for j := i + 1; j < jumlahPetugas; j++ {
			if ascending {
				if petugasParkir[j].ID < petugasParkir[idx].ID {
					idx = j
				}
			} else {
				if petugasParkir[j].ID > petugasParkir[idx].ID {
					idx = j
				}
			}
		}
		petugasParkir[i], petugasParkir[idx] = petugasParkir[idx], petugasParkir[i]
	}
}

func editPetugas(petugasParkir *Petugas, jumlahPetugas *int) {
	var id int

	fmt.Print("Masukkan ID Petugas yang akan diedit: ")
	fmt.Scanln(&id)

	foundIndex := cariPetugasByIDSequential(petugasParkir, *jumlahPetugas, id)

	if foundIndex != -1 {
		fmt.Println("Masukkan data baru untuk petugas:")
		fmt.Print("Nama: ")
		fmt.Scanln(&petugasParkir[foundIndex].Nama)
		fmt.Print("Username: ")
		fmt.Scanln(&petugasParkir[foundIndex].Username)
		fmt.Print("Password: ")
		fmt.Scanln(&petugasParkir[foundIndex].Password)

		fmt.Println("Data petugas berhasil diubah.")
	} else {
		fmt.Println("Petugas dengan ID tersebut tidak ditemukan.")
	}
}

//Function dan Sequential Search
func cariPetugasByIDSequential(petugasParkir *Petugas, jumlahPetugas int, id int) int {
	for i := 0; i < jumlahPetugas; i++ {
		if petugasParkir[i].ID == id {
			return i
		}
	}
	return -1
}

func hapusPetugas(petugasParkir *Petugas, jumlahPetugas *int) {
	var id int
	fmt.Print("Masukkan ID Petugas yang akan dihapus: ")
	fmt.Scanln(&id)

	insertionSort(petugasParkir, *jumlahPetugas)

	foundIndex := cariPetugasByIDBinary(petugasParkir, *jumlahPetugas, id)

	if foundIndex != -1 {
		for i := foundIndex; i < *jumlahPetugas-1; i++ {
			petugasParkir[i] = petugasParkir[i+1]
		}

		*jumlahPetugas--
		petugasParkir[*jumlahPetugas] = PetugasParkir{}

		fmt.Println("Petugas berhasil dihapus.")
	} else {
		fmt.Println("Petugas dengan ID tersebut tidak ditemukan.")
	}
}

//Insertion Sort
func insertionSort(petugasParkir *Petugas, jumlahPetugas int) {
	i := 1
	for i <= jumlahPetugas-1 {
		j := i
		temp := petugasParkir[j]
		for j > 0 && temp.ID < petugasParkir[j-1].ID {
			petugasParkir[j] = petugasParkir[j-1]
			j = j - 1
		}
		petugasParkir[j] = temp
		i = i + 1
	}
}

//Binary Search
func cariPetugasByIDBinary(petugasParkir *Petugas, jumlahPetugas int, id int) int {
	kr := 0
	kn := jumlahPetugas - 1
	found := false

	for kr <= kn && !found {
		med := (kr + kn) / 2
		if petugasParkir[med].ID < id {
			kr = med + 1
		} else if petugasParkir[med].ID > id {
			kn = med - 1
		} else {
			found = true
			return med
		}
	}
	return -1
}

func petugas(petugasParkir *Petugas, transaksiParkir *Transaksi, jumlahKendaraan *int, jumlahPetugas int) {
	var username string
	var password string
	var exit bool

	exit = false

	for exit == false {
		fmt.Println("\nMasukan Username Dan Password Petugas Parkir")
		fmt.Println("Ketik 'keluar' jika anda bukan petugas parkir atau kembali ke menu awal")
		fmt.Println("Masukkan Username: ")
		fmt.Scanln(&username)
		if username == "keluar" || username == "Keluar" {
			exit = true
		} else {
			fmt.Println("Masukan Password: ")
			fmt.Scanln(&password)

			foundIndex := loginPetugas(*petugasParkir, jumlahPetugas, username, password)

			if foundIndex == -1 {
				fmt.Println("Username atau Password salah, coba lagi.")
			} else {
				fmt.Print("Selamat Datang! ", petugasParkir[foundIndex].Nama)
				fmt.Println()
				entitasPetugas(transaksiParkir, jumlahKendaraan)
			}
		}
	}
}

func loginPetugas(PP Petugas, jumlahPetugas int, username string, password string) int {
	for i := 0; i < jumlahPetugas; i++ {
		if PP[i].Username == username && PP[i].Password == password {
			return i
		}
	}
	return -1
}

func entitasPetugas(transaksiParkir *Transaksi, jumlahKendaraan *int) {
	var pilihan int
	var exit bool

	exit = false

	for pilihan != 3 && exit != true {
		fmt.Println("\nPilih Menu : ")
		fmt.Println("1. Tambah Transaksi Parkir")
		fmt.Println("2. Cetak Daftar Kendaraan")
		fmt.Println("3. Cetak Total Biaya Kendaraan")
		fmt.Println("4. Keluar ( kembali ke menu pemilihan role)")
		fmt.Scanln(&pilihan)
		switch pilihan {
		case 1:
			tambahDataParkir(transaksiParkir, jumlahKendaraan)
		case 2:
			cetakDaftarKendaraanMasuk(transaksiParkir, jumlahKendaraan)
		case 3:
			cetakTotalBiayaParkir(transaksiParkir, jumlahKendaraan)
		case 4:
			exit = true
		default:
			fmt.Println("Pilihan tidak valid!")
		}
	}
}

func tambahDataParkir(transaksiParkir *Transaksi, jumlahKendaraan *int) {
	var exit, Test, TestKeluar bool
	var B, C float32
	Test = false

	exit = false

	for !exit {
		if *jumlahKendaraan >= tParkir {
			fmt.Println("Kapasitas parkir sudah penuh.")
			exit = true
		} else {
			var transaksi TransaksiParkir
			validVehicle := false

			fmt.Println("Masukkan Data Parkir:")
			fmt.Print("Tipe Kendaraan: ")
			fmt.Println("Masukan angka '1' jika Mobil")
			fmt.Println("Masukan angka '2' jika Motor")
			fmt.Print("Tipe Kendaran: ")
			for validVehicle == false {
				fmt.Scanln(&transaksi.TipeKendaraan)
				if transaksi.TipeKendaraan == "1" {
					transaksi.TipeKendaraan = "Mobil"
					validVehicle = true
				} else if transaksi.TipeKendaraan == "2" {
					transaksi.TipeKendaraan = "Motor"
					validVehicle = true
				} else {
					validVehicle = false
					fmt.Print("Pilihan Tidak Valid")
				}
			}
			fmt.Print("Plat Kendaraan: ")
			fmt.Scanln(&transaksi.Plat)

			fmt.Print("Waktu Masuk (dalam format jam, misal 13.30): ")
			for Test == false {
				fmt.Scanln(&transaksi.WaktuMasuk)
				C = (transaksi.WaktuMasuk)
				if C < 24.00 {
					C := int(C)
					B = transaksi.WaktuMasuk - float32(C)
					if B <= 0.60 {
						Test = true
					} else {
						Test = false
						fmt.Println("Masukan Waktu Dengan Format Yang Benar")
					}
				} else {
					Test = false
					fmt.Println("Masukan Waktu Dengan Format Yang Benar")
				}
			}
			fmt.Print("Waktu Keluar (dalam format jam, misal 15.45): ")
			for TestKeluar == false {
				fmt.Scanln(&transaksi.WaktuKeluar)
				C = (transaksi.WaktuKeluar)
				if C < 24.00 {
					C := int(C)
					B = transaksi.WaktuKeluar - float32(C)
					if B < 0.60 {
						TestKeluar = true
					} else {
						TestKeluar = false
						fmt.Println("Masukan Waktu Dengan Format Yang Benar")
					}
				} else {
					TestKeluar = false
					fmt.Println("Masukan Waktu Dengan Format Yang Benar")
				}
			}

			if transaksi.TipeKendaraan == "Mobil" {
				transaksi.BiayaParkir = (transaksi.WaktuKeluar - transaksi.WaktuMasuk) * 5000
			} else if transaksi.TipeKendaraan == "Motor" {
				transaksi.BiayaParkir = (transaksi.WaktuKeluar - transaksi.WaktuMasuk) * 3000
			}
			transaksi.JumlahKendaraan = *jumlahKendaraan + 1

			transaksiParkir[*jumlahKendaraan] = transaksi
			*jumlahKendaraan++

			fmt.Println("Apakah ingin menambah data parkir lagi? (yes/no)")
			var response string
			fmt.Scanln(&response)
			if response != "yes" {
				exit = true
			}
		}
	}
}

func cetakDaftarKendaraanMasuk(transaksiParkir *Transaksi, jumlahKendaraan *int) {
	fmt.Println("| Jenis kendaraan |   | Plat Nomor |   | Waktu Masuk |   | Waktu Keluar |   | Biaya Parkir |")
	for i := 0; i < *jumlahKendaraan; i++ {
		t := transaksiParkir[i]
		fmt.Printf("| %-15s |   | %-10s |   | %-10.2f |   | %-12.2f |   | %-10.2f |\n", t.TipeKendaraan, t.Plat, t.WaktuMasuk, t.WaktuKeluar, t.BiayaParkir)
	}
}

func cetakTotalBiayaParkir(transaksiParkir *Transaksi, jumlahKendaraan *int) {
	var totalBiaya float32

	for i := 0; i < *jumlahKendaraan; i++ {
		totalBiaya += transaksiParkir[i].BiayaParkir
	}

	fmt.Printf("Total Biaya Parkir Hari Ini: %.2f\n", totalBiaya)
	entitasPetugas(transaksiParkir, jumlahKendaraan)
}
