package main

import (
	"fmt"
	"time"
)

type Calon struct {
	Nama   string
	Partai string
	Suara  int
}

type Pemilih struct {
	Nama    string
	Pilihan string
}

type Pemilihan struct {
	DaftarCalon   []Calon
	DaftarPemilih []Pemilih
	WaktuMulai    time.Time
	WaktuSelesai  time.Time
}

func (p *Pemilihan) TambahCalon(nama, partai string) {
	p.DaftarCalon = append(p.DaftarCalon, Calon{Nama: nama, Partai: partai})
}

func (p *Pemilihan) EditCalon(index int, nama, partai string) {
	if index >= 0 && index < len(p.DaftarCalon) {
		p.DaftarCalon[index].Nama = nama
		p.DaftarCalon[index].Partai = partai
	}
}

func (p *Pemilihan) HapusCalon(index int) {
	if index >= 0 && index < len(p.DaftarCalon) {
		p.DaftarCalon = append(p.DaftarCalon[:index], p.DaftarCalon[index+1:]...)
	}
}

func (p *Pemilihan) TambahPemilih(nama, pilihan string) {
	p.DaftarPemilih = append(p.DaftarPemilih, Pemilih{Nama: nama, Pilihan: pilihan})
	for i := range p.DaftarCalon {
		if p.DaftarCalon[i].Nama == pilihan {
			p.DaftarCalon[i].Suara++
			break
		}
	}
}

func (p *Pemilihan) TampilkanDaftarCalon() {
	fmt.Println("Daftar Calon:")
	for i, c := range p.DaftarCalon {
		fmt.Printf("%d. %s (%s) - %d suara\n", i+1, c.Nama, c.Partai, c.Suara)
	}
}

func (p *Pemilihan) UrutkanBerdasarkanNama() {
	n := len(p.DaftarCalon)
	for i := 0; i < n-1; i++ {
		for j := 0; j < n-i-1; j++ {
			if p.DaftarCalon[j].Nama > p.DaftarCalon[j+1].Nama {
				p.DaftarCalon[j], p.DaftarCalon[j+1] = p.DaftarCalon[j+1], p.DaftarCalon[j]
			}
		}
	}
	p.TampilkanDaftarCalon()
}

func (p *Pemilihan) TampilkanPemenang() {
	if len(p.DaftarCalon) == 0 {
		fmt.Println("Tidak ada calon dalam daftar.")
		return
	}
	totalSuara := 0
	for _, c := range p.DaftarCalon {
		totalSuara += c.Suara
	}

	threshold := totalSuara/2 + 1
	for _, c := range p.DaftarCalon {
		if c.Suara >= threshold {
			fmt.Println("Pemenang Pemilihan:")
			fmt.Printf("%s (%s) - %d suara\n", c.Nama, c.Partai, c.Suara)
			return
		}
	}
	fmt.Println("Tidak ada pemenang yang memenuhi ambang batas.")
}

func (p *Pemilihan) CariBerdasarkanNamaCalon(nama string) {
	fmt.Printf("Mencari calon: %s\n", nama)
	for _, c := range p.DaftarCalon {
		if c.Nama == nama {
			fmt.Printf("%s (%s) - %d suara\n", c.Nama, c.Partai, c.Suara)
			fmt.Print("Apakah ingin memunculkan pemilih? (ya/tidak): ")
			var jawaban string
			fmt.Scan(&jawaban)
			if jawaban == "ya" {
				fmt.Println("Daftar Pemilih:")
				for _, pemilih := range p.DaftarPemilih {
					if pemilih.Pilihan == nama {
						fmt.Println(pemilih.Nama)
					}
				}
			}
			return
		}
	}
	fmt.Println("Calon tidak ditemukan.")
}

func (p *Pemilihan) AturWaktuPemilihan(waktuMulai, waktuSelesai string) {
	mulai, err1 := time.Parse("02/01/2006", waktuMulai)
	selesai, err2 := time.Parse("02/01/2006", waktuSelesai)
	if err1 != nil || err2 != nil {
		fmt.Println("Format tanggal tidak valid. Gunakan format DD/MM/YYYY.")
	} else {
		p.WaktuMulai = mulai
		p.WaktuSelesai = selesai
		fmt.Println("Waktu pemilihan berhasil diatur.")
	}
}

func (p *Pemilihan) ApakahPemilihanBuka() bool {
	sekarang := time.Now()
	return sekarang.After(p.WaktuMulai) && sekarang.Before(p.WaktuSelesai)
}

func menuAdmin(p *Pemilihan) {
	var pilihan int
	for {
		fmt.Println("Menu Admin:")
		fmt.Println("1. Tambah Calon")
		fmt.Println("2. Edit Calon")
		fmt.Println("3. Hapus Calon")
		fmt.Println("4. Atur Waktu Pemilihan")
		fmt.Println("5. Tampilkan Daftar Calon")
		fmt.Println("6. Keluar")
		fmt.Print("Pilih opsi: ")
		fmt.Scan(&pilihan)

		switch pilihan {
		case 1:
			var nama, partai string
			fmt.Print("Masukkan nama calon: ")
			fmt.Scan(&nama)
			fmt.Print("Masukkan partai calon: ")
			fmt.Scan(&partai)
			p.TambahCalon(nama, partai)
		case 2:
			var index int
			var nama, partai string
			p.TampilkanDaftarCalon()
			fmt.Print("Masukkan nomor calon yang akan diedit: ")
			fmt.Scan(&index)
			fmt.Print("Masukkan nama baru: ")
			fmt.Scan(&nama)
			fmt.Print("Masukkan partai baru: ")
			fmt.Scan(&partai)
			p.EditCalon(index-1, nama, partai)
		case 3:
			var index int
			p.TampilkanDaftarCalon()
			fmt.Print("Masukkan nomor calon yang akan dihapus: ")
			fmt.Scan(&index)
			p.HapusCalon(index - 1)
		case 4:
			var waktuMulai, waktuSelesai string
			fmt.Print("Masukkan waktu mulai (DD/MM/YYYY): ")
			fmt.Scan(&waktuMulai)
			fmt.Print("Masukkan waktu selesai (DD/MM/YYYY): ")
			fmt.Scan(&waktuSelesai)
			p.AturWaktuPemilihan(waktuMulai, waktuSelesai)
		case 5:
			p.TampilkanDaftarCalon()
		case 6:
			return
		default:
			fmt.Println("Pilihan tidak valid. Coba lagi.")
		}
	}
}

func menuPemilih(p *Pemilihan) {
	var pilihan int
	for {
		fmt.Println("Menu Pemilih:")
		if p.ApakahPemilihanBuka() {
			fmt.Println("1. Berikan Suara")
		}
		fmt.Println("2. Tampilkan Daftar Calon")
		fmt.Println("3. Cari Calon Berdasarkan Nama")
		fmt.Println("4. Tampilkan Pemenang")
		fmt.Println("5. Keluar")
		fmt.Print("Pilih opsi: ")
		fmt.Scan(&pilihan)

		switch pilihan {
		case 1:
			if p.ApakahPemilihanBuka() {
				var nama, pilihanCalon string
				fmt.Print("Masukkan nama Anda: ")
				fmt.Scan(&nama)
				fmt.Print("Masukkan nama calon yang dipilih: ")
				fmt.Scan(&pilihanCalon)
				p.TambahPemilih(nama, pilihanCalon)
			} else {
				fmt.Println("Pemilihan belum dibuka.")
			}
		case 2:
			p.UrutkanBerdasarkanNama()
		case 3:
			var nama string
			fmt.Print("Masukkan nama calon: ")
			fmt.Scan(&nama)
			p.CariBerdasarkanNamaCalon(nama)
		case 4:
			p.TampilkanPemenang()
		case 5:
			return
		default:
			fmt.Println("Pilihan tidak valid. Coba lagi.")
		}
	}
}

func main() {
	p := Pemilihan{}

	var tipePengguna int
	for {
		fmt.Println("Selamat datang di Sistem Pemilihan")
		fmt.Println("1. Admin")
		fmt.Println("2. Pemilih")
		fmt.Println("3. Keluar")
		fmt.Print("Pilih tipe pengguna: ")
		fmt.Scan(&tipePengguna)

		switch tipePengguna {
		case 1:
			menuAdmin(&p)
		case 2:
			menuPemilih(&p)
		case 3:
			fmt.Println("Keluar...")
			return
		default:
			fmt.Println("Pilihan tidak valid.")
		}
	}
}
