package main

import (
	"os"
	"strconv"
)

const (
	IM             = 139968
	IA             = 3877
	IC             = 29573
	SEED           = 42
	BUFLINES       = 100
	LINELEN        = 60
	SMALL_DATASET  = 6250000
	MEDIUM_DATASET = 12500000
	LARGE_DATASET  = 25000000
	alu            = "GGCCGGGCGCGGTGGCTCACGCCTGTAATCCCAGCACTTTGG" +
		"GAGGCCGAGGCGGGCGGATCACCTGAGGTCAGGAGTTCGAGA" +
		"CCAGCCTGGCCAACATGGTGAAACCCCGTCTCTACTAAAAAT" +
		"ACAAAAATTAGCCGGGCGTGGTGGCGCGCGCCTGTAATCCCA" +
		"GCTACTCGGGAGGCTGAGGCAGGAGAATCGCTTGAACCCGGG" +
		"AGGCGGAGGTTGCAGTGAGCCGAGATCGCGCCACTGCACTCC" +
		"AGCCTGGGCGACAGAGCGAGACTCCGTCTCAAAAA"
	iub         = "acgtBDHKMNRSVWY"
	homosapiens = "acgt"
)

var (
	header1 = []byte(">ONE Homo sapiens alu\n")
	header2 = []byte(">TWO IUB ambiguity codes\n")
	header3 = []byte(">THREE Homo sapiens frequency\n")

	seed  uint32 = SEED
	iub_p        = [...]float32{
		0.27,
		0.12,
		0.12,
		0.27,
		0.02,
		0.02,
		0.02,
		0.02,
		0.02,
		0.02,
		0.02,
		0.02,
		0.02,
		0.02,
		0.02,
	}
	homosapiens_p = [...]float32{
		0.3029549426680,
		0.1979883004921,
		0.1975473066391,
		0.3015094502008,
	}
)

func uint32_rand() uint32 {
	seed = (seed*IA + IC) % IM
	return seed
}

func repeat_fasta(seq string, n int) {
	len := len(seq)
	buflen1 := len + LINELEN
	buffer1 := make([]byte, buflen1)
	var i int

	if LINELEN < len {
		copy(buffer1, seq)
		copy(buffer1[len:], seq[:LINELEN])
	} else {
		for i := 0; i < LINELEN/len; i++ {
			copy(buffer1[i*len:], seq)
		}
		copy(buffer1[i*len:], seq[:n-i*n])
	}

	buflen2 := (LINELEN + 1) * len
	buffer2 := make([]byte, buflen2)
	for i := 0; i < len; i++ {
		copy(buffer2[i*(LINELEN+1):], buffer1[((i*LINELEN)%len):((i*LINELEN)%len)+LINELEN])
		buffer2[(i+1)*(LINELEN+1)-1] = '\n'
	}

	whole_buffers := n / (len * LINELEN)
	for i := 0; i < whole_buffers; i++ {
		os.Stdout.Write(buffer2)
	}

	data_remaining := n - whole_buffers*len*LINELEN
	embedded_newlines := data_remaining / LINELEN
	os.Stdout.Write(buffer2[:data_remaining+embedded_newlines])

	if n%LINELEN != 0 {
		os.Stdout.Write([]byte{'\n'})
	}

}

func random_fasta(symb string, probability []float32, n int) {
	hash := build_hash(symb, probability)
	buffer := buffer_with_linebreaks(BUFLINES)

	buffers := n / LINELEN / BUFLINES
	for i := 0; i < buffers; i++ {
		for j := 0; j < BUFLINES; j++ {
			for k := 0; k < LINELEN; k++ {
				v := uint32_rand()
				buffer[j*(LINELEN+1)+k] = hash[v]
			}
		}
		os.Stdout.Write(buffer[:((LINELEN + 1) * BUFLINES)])
	}

	lines := n/LINELEN - buffers*BUFLINES
	for j := 0; j < lines; j++ {
		for k := 0; k < LINELEN; k++ {
			v := uint32_rand()
			buffer[j*(LINELEN+1)+k] = hash[v]
		}
	}
	partials := n - LINELEN*lines - buffers*BUFLINES*LINELEN
	for k := 0; k < partials; k++ {
		v := uint32_rand()
		buffer[lines*(LINELEN+1)+k] = hash[v]
	}
	os.Stdout.Write(buffer[:lines*(LINELEN+1)+partials])

	if n%LINELEN != 0 {
		os.Stdout.Write([]byte{'\n'})
	}

}

func build_hash(symb string, probability []float32) []byte {
	hash := make([]byte, IM)
	if hash == nil {
		os.Exit(-1)
	}

	sum := float32(0.0)
	len := len(symb)
	sum = probability[0]

	for i, j := 0, 0; i < IM && j < len; i++ {
		r := 1.0 * float32(i) / float32(IM)
		if r >= sum {
			j++
			sum += probability[j]
		}
		hash[i] = symb[j]
	}

	return hash
}

func buffer_with_linebreaks(lines int) []byte {

	buffer := make([]byte, (LINELEN+1)*lines)
	if buffer == nil {
		os.Exit(-1)
	}
	for i := 0; i < lines; i++ {
		buffer[i*(LINELEN+1)+LINELEN] = '\n'
	}
	return buffer

}

func main() {
	//n := 1000
	//if (argc>1) n = atoi(argv[1])
	n := LARGE_DATASET

	os.Stdout.Write(header1)
	repeat_fasta(alu, n*2)
	os.Stdout.Write(header2)
	random_fasta(iub, iub_p[:], n*3)

	os.Stdout.Write(header3)
	random_fasta(homosapiens, homosapiens_p[:], n*5)
	os.Stdout.Write([]byte("o input é:" + strconv.Itoa(n) + "\n"))
}
