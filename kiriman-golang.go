package main

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

func updateKiriman() {
	db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/kiriman_console")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	var nomor int
	var tanggal string
	var bulan string
	var tahun int
	var hari string
	var jumlah int

	fmt.Print("No Kirim \t: ")
	fmt.Scanln(&nomor)
	fmt.Print("Hari \t: ")
	fmt.Scanln(&hari)
	fmt.Print("Tanggal \t: ")
	fmt.Scanln(&tanggal)
	fmt.Print("Bulan \t: ")
	fmt.Scanln(&bulan)
	fmt.Print("Tahun \t: ")
	fmt.Scanln(&tahun)
	fmt.Print("Jumlah Kirim \t: ")
	fmt.Scanln(&jumlah)

	stmt, err := db.Prepare("UPDATE kiriman SET no_kirim=?, tgl=?, bulan=?, tahun=?, hari=?, jml_kiriman=? WHERE no_kirim=?")
	checkErr(err)

	res, err := stmt.Exec(nomor, tanggal, bulan, tahun, hari, jumlah)
	checkErr(err)
	id, err := res.LastInsertId()
	checkErr(err)
	fmt.Println(id)
	fmt.Println("Data berhasil diubah!")
	db.Close()
}

func insertKiriman() {
	db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/kiriman_console")
	checkErr(err)
	defer db.Close()

	var noKir int
	var tanggal string
	var bulan string
	var tahun int
	var hari string
	var jumlahKir int

	fmt.Print("No Kirim \t: ")
	fmt.Scanln(&noKir)
	fmt.Print("Hari \t: ")
	fmt.Scanln(&hari)
	fmt.Print("Tanggal \t: ")
	fmt.Scanln(&tanggal)
	fmt.Print("Bulan \t: ")
	fmt.Scanln(&bulan)
	fmt.Print("Tahun \t: ")
	fmt.Scanln(&tahun)
	fmt.Print("Jumlah Kirim : ")
	fmt.Scanln(&jumlahKir)

	stmt, err := db.Prepare("INSERT kiriman SET no_kirim=?, tgl=?, bulan=?, tahun=?, hari=?, jml_kiriman=?")
	checkErr(err)
	res, err := stmt.Exec(noKir, tanggal, bulan, tahun, hari, jumlahKir)
	checkErr(err)
	id, err := res.LastInsertId()
	checkErr(err)
	fmt.Println(id)
	fmt.Println("Data berhasil dimasukan!")
	db.Close()
}

func deleteKiriman() {
	db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/kiriman_console")
	checkErr(err)
	defer db.Close()
	var noKir int
	fmt.Print("No Kirim : ")
	fmt.Scanln(&noKir)

	stmt, err := db.Prepare("DELETE FROM kiriman WHERE no_kirim=?")
	checkErr(err)
	res, err := stmt.Exec(noKir)
	checkErr(err)
	affect, err := res.RowsAffected()
	checkErr(err)
	fmt.Println(affect)
	db.Close()
}

func selectKiriman() {
	db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/kiriman_console")
	checkErr(err)
	defer db.Close()
	fmt.Println("No Kirim\t", "Hari\t\t", "Tanggal\t\t\t", "Jumlah Kirim\t")
	rows, err := db.Query("SELECT * FROM kiriman")
	checkErr(err)
	for rows.Next() {
		var nomor int
		var tanggal int
		var bulan string
		var tahun int
		var hari string
		var jumlah int
		err = rows.Scan(&nomor, &tanggal, &bulan, &tahun, &hari, &jumlah)
		if err != nil {
			panic(err.Error())
		}
		fmt.Println(nomor, "\t\t", hari, "\t\t", tanggal, "-", bulan, "-", tahun, "\t\t", jumlah)
	}
}

func sumKiriman() {

}

func menuKiriman() {
	fmt.Println("+---------------------------------------+")
	fmt.Println("|		Menu Kiriman		|")
	fmt.Println("+---------------------------------------+")
	fmt.Println("|	1. Tambah Data Kiriman		|")
	fmt.Println("|	2. Lihat Data Kiriman		|")
	fmt.Println("|	3. Ubah Data Kiriman		|")
	fmt.Println("|	4. Hapus Data Kiriman		|")
	fmt.Println("|	5. Jumlah Data Kiriman		|")
	fmt.Println("|	0. Kembali 			|")
	fmt.Println("+---------------------------------------+")
	var x int
	fmt.Print("=> ")
	fmt.Scanln(&x)
	switch x {
	case 1:
		fmt.Println("Tambah Data Kiriman")
		insertKiriman()
		menuKiriman()
		break
	case 2:
		fmt.Println("Lihat Data Kiriman")
		selectKiriman()
		menuKiriman()
		break
	case 3:
		fmt.Println("Ubah Data Kiriman")
		updateKiriman()
		menuKiriman()
		break
	case 4:
		fmt.Println("Hapus Data Kiriman")
		selectKiriman()
		deleteKiriman()
		menuKiriman()
		break
	case 5:
		fmt.Println("Jumlah Data Kiriman")
		break
	case 0:
		mainMenu()
		break
	}
}

func mainMenu() {
	fmt.Println("+-------------------------------+")
	fmt.Println("|	Aplikasi Kiriman	|")
	fmt.Println("+-------------------------------+")
	fmt.Println("|	1. Kiriman		|")
	fmt.Println("|	2. Cari			|")
	fmt.Println("|	0. Keluar		|")
	fmt.Println("+-------------------------------+")
	var x int
	fmt.Print("=> ")
	fmt.Scanln(&x)
	switch x {
	case 1:
		fmt.Println("Kiriman")
		menuKiriman()
		break
	case 2:
		fmt.Println("Cari")
		break
	case 0:
		fmt.Println("Terimakasih sudah menggunakan layanan yang anda buat sendiri...!")
		os.Exit(0)
		break
	default:
		fmt.Println("\nSalah inputan!\n")
		mainMenu()
		break
	}
}

func connectDB() {
	fmt.Println("Connect to database...(kiriman)")
	db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/kiriman_console")
	checkErr(err)
	con := "Connected"
	fmt.Println(con)
	defer db.Close()
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	//connectDB()
	mainMenu()
}
