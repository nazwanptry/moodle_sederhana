package main

import (
	"fmt"
	"os"
)

type User struct {
	Username string
	Role     string
}

type Content struct {
	Title        string
	Type         string
	Questions    []Question
	Participants []Participant
}

type Question struct {
	Question      string
	Choices       [4]string
	CorrectAnswer string
}

type Participant struct {
	Username string
	Answers  []string
	Score    int
}

var contents []Content

func main() {
	fmt.Println("**************************")
	fmt.Println("* Created by Safa and Nazwan *")
	fmt.Println("**************************")
	for {
		user := login()
		if user.Role == "guru" {
			menuGuru(user)
		} else if user.Role == "siswa" {
			menuSiswa(user)
		} else {
			fmt.Println("Peran tidak valid, silakan coba lagi.")
		}
	}
}

func login() User {
	for {
		var username, role string
		fmt.Println("Selamat datang di aplikasi Moodle sederhana.")
		fmt.Println("1. Login")
		fmt.Println("2. Keluar")
		fmt.Print("Pilih opsi: ")
		var choice int
		fmt.Scan(&choice)
		switch choice {
		case 1:
			fmt.Print("Masukkan username: ")
			fmt.Scan(&username)
			fmt.Print("Masukkan peran (guru/siswa): ")
			fmt.Scan(&role)
			return User{Username: username, Role: role}
		case 2:
			fmt.Println("Terima kasih telah menggunakan aplikasi ini.")
			exit()
		default:
			fmt.Println("Opsi tidak valid, silakan coba lagi.")
		}
	}
}

func menuGuru(user User) {
	for {
		var choice int
		fmt.Println("\nMenu Guru")
		fmt.Println("1. Tambah Konten")
		fmt.Println("2. Edit Konten")
		fmt.Println("3. Hapus Konten")
		fmt.Println("4. Lihat Konten")
		fmt.Println("5. Lihat Nilai Mahasiswa")
		fmt.Println("6. Keluar")
		fmt.Print("Pilih opsi: ")
		fmt.Scan(&choice)
		switch choice {
		case 1:
			addContent()
		case 2:
			editContent()
		case 3:
			deleteContent()
		case 4:
			viewContent()
		case 5:
			viewScores()
		case 6:
			return
		default:
			fmt.Println("Opsi tidak valid.")
		}
	}
}

func menuSiswa(user User) {
	for {
		var choice int
		fmt.Println("\nMenu Siswa")
		fmt.Println("1. Ikuti Quiz")
		fmt.Println("2. Kerjakan Tugas")
		fmt.Println("3. Forum Diskusi")
		fmt.Println("4. Keluar")
		fmt.Print("Pilih opsi: ")
		fmt.Scan(&choice)
		switch choice {
		case 1:
			takeQuiz(user)
		case 2:
			submitAssignment(user)
		case 3:
			discussionForum(user)
		case 4:
			return
		default:
			fmt.Println("Opsi tidak valid.")
		}
	}
}

func addContent() {
	var title, contentType string
	fmt.Print("Masukkan judul konten: ")
	fmt.Scan(&title)
	fmt.Print("Masukkan tipe konten (tugas/quiz/forum): ")
	fmt.Scan(&contentType)
	content := Content{Title: title, Type: contentType}

	if contentType == "tugas" || contentType == "quiz" {
		var question, correctAnswer string
		var choices [4]string
		for {
			fmt.Print("Masukkan soal (ketik 'selesai' untuk berhenti): ")
			fmt.Scan(&question)
			if question == "selesai" {
				break
			}
			if contentType == "quiz" {
				fmt.Print("Masukkan jawaban a: ")
				fmt.Scan(&choices[0])
				fmt.Print("Masukkan jawaban b: ")
				fmt.Scan(&choices[1])
				fmt.Print("Masukkan jawaban c: ")
				fmt.Scan(&choices[2])
				fmt.Print("Masukkan jawaban d: ")
				fmt.Scan(&choices[3])
				fmt.Print("Masukkan jawaban yang benar (a/b/c/d): ")
				fmt.Scan(&correctAnswer)
			} else {
				correctAnswer = ""
			}
			content.Questions = append(content.Questions, Question{Question: question, Choices: choices, CorrectAnswer: correctAnswer})
		}
	}

	contents = append(contents, content)
	fmt.Println("Konten berhasil ditambahkan.")
}

func editContent() {
	var title string
	fmt.Print("Masukkan judul konten yang ingin diubah: ")
	fmt.Scan(&title)

	for i, content := range contents {
		if content.Title == title {
			fmt.Print("Masukkan judul baru: ")
			fmt.Scan(&contents[i].Title)
			fmt.Println("Konten berhasil diubah.")
			return
		}
	}
	fmt.Println("Konten tidak ditemukan.")
}

func deleteContent() {
	var title string
	fmt.Print("Masukkan judul konten yang ingin dihapus: ")
	fmt.Scan(&title)

	for i, content := range contents {
		if content.Title == title {
			contents = append(contents[:i], contents[i+1:]...)
			fmt.Println("Konten berhasil dihapus.")
			return
		}
	}
	fmt.Println("Konten tidak ditemukan.")
}

func viewContent() {
	for _, content := range contents {
		fmt.Printf("Judul: %s, Tipe: %s\n", content.Title, content.Type)
		if content.Type == "tugas" || content.Type == "quiz" {
			for _, question := range content.Questions {
				fmt.Printf("- Soal: %s\n", question.Question)
				if content.Type == "quiz" {
					fmt.Printf("  a. %s\n", question.Choices[0])
					fmt.Printf("  b. %s\n", question.Choices[1])
					fmt.Printf("  c. %s\n", question.Choices[2])
					fmt.Printf("  d. %s\n", question.Choices[3])
					fmt.Printf("  Jawaban yang benar: %s\n", question.CorrectAnswer)
				}
			}
		}
		fmt.Printf("Jumlah Peserta: %d\n", len(content.Participants))
	}
}

func takeQuiz(user User) {
	var title, answer string
	fmt.Print("Masukkan judul quiz: ")
	fmt.Scan(&title)

	for i, content := range contents {
		if content.Title == title && content.Type == "quiz" {
			participant := Participant{Username: user.Username}
			score := 0
			for _, question := range content.Questions {
				fmt.Println(question.Question)
				fmt.Printf("a. %s\n", question.Choices[0])
				fmt.Printf("b. %s\n", question.Choices[1])
				fmt.Printf("c. %s\n", question.Choices[2])
				fmt.Printf("d. %s\n", question.Choices[3])
				fmt.Print("Jawaban: ")
				fmt.Scan(&answer)
				participant.Answers = append(participant.Answers, answer)
				if answer == question.CorrectAnswer {
					score++
				}
			}
			participant.Score = score
			contents[i].Participants = append(contents[i].Participants, participant)
			fmt.Printf("Quiz selesai. Skor Anda: %d\n", score)
			return
		}
	}
	fmt.Println("Quiz tidak ditemukan.")
}

func submitAssignment(user User) {
	var title, answer string
	fmt.Print("Masukkan judul tugas: ")
	fmt.Scan(&title)

	for i, content := range contents {
		if content.Title == title && content.Type == "tugas" {
			participant := Participant{Username: user.Username}
			for _, question := range content.Questions {
				fmt.Println(question.Question)
				fmt.Print("Jawaban: ")
				fmt.Scan(&answer)
				participant.Answers = append(participant.Answers, answer)
			}
			contents[i].Participants = append(contents[i].Participants, participant)
			fmt.Println("Tugas selesai.")
			return
		}
	}
	fmt.Println("Tugas tidak ditemukan.")
}

func discussionForum(user User) {
	var title string
	fmt.Print("Masukkan judul forum: ")
	fmt.Scan(&title)

	for i, content := range contents {
		if content.Title == title && content.Type == "forum" {
			fmt.Println("Masukkan pertanyaan atau jawaban (ketik 'selesai' untuk berhenti):")
			for {
				var post string
				fmt.Print("Pesan: ")
				fmt.Scan(&post)
				if post == "selesai" {
					break
				}
				fmt.Printf("%s: %s\n", user.Username, post)
				content.Participants = append(content.Participants, Participant{Username: user.Username, Answers: []string{post}})
			}
			contents[i] = content // Update the content with new posts
			fmt.Println("Diskusi forum selesai.")
			return
		}
	}
	fmt.Println("Forum tidak ditemukan.")
}

func viewScores() {
	for _, content := range contents {
		fmt.Printf("Judul: %s, Tipe: %s\n", content.Title, content.Type)
		if content.Type == "tugas" || content.Type == "quiz" {
			fmt.Println("Nilai Mahasiswa:")
			for _, participant := range content.Participants {
				fmt.Printf("- %s: %d\n", participant.Username, participant.Score)
			}
		}
	}
}

func exit() {
	fmt.Println("Program dihentikan.")
	os.Exit(0)
}
