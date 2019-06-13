package main

import (
	"bufio"
	"github.com/namsral/flag"
	"log"
	"os"

)

func main() {
	awsCredPath := flag.String("aws-credentials", "config/credentials", "Path to AWS credentials file")
	awsCredProfile := flag.String("aws-credentials-profile", "default", "Profile in AWS credentials to use")
	awsRegion := flag.String("aws-region", "us-east-1", "AWS region to send requests to")
	awsVoice := flag.String("aws-voice", "Miguel", "AWS region to send requests to")
	ttsFrequency := flag.Int("tts-frequency", 16000, "AWS region to send requests to")

	polly, err := NewPolly(*awsCredPath, *awsCredProfile, *awsRegion)

	file, err := os.Open("input_files/spanish_data")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	file1, err := os.Create("output_files/output_spanish_data.txt")
	if err != nil {
		log.Fatal(err)
	}
	w := bufio.NewWriter(file1)

	if err == nil {
		ttsPolly := NewTTSPolly(polly, *awsVoice, *ttsFrequency)
		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			text := scanner.Text()
			filePath, err := ttsPolly.Speak(text)
			if err == nil {
				w.WriteString(text + "\t\t" + filePath)
				w.WriteByte('\n')
			}
		}

		if err := scanner.Err(); err != nil {
			log.Fatal(err)
		}
	} else {
		log.Println("Error in creating Polly",err)
	}
	w.Flush()
	file1.Close()
}

