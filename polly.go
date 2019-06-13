package main
import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/polly"
	"github.com/youpy/go-wav"
	"io/ioutil"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
	"time"
)

func NewTTSPolly(polly *polly.Polly, voice string, frequency int) *TTSPolly {
	return &TTSPolly{
		polly:                  polly,
		voiceId:                voice,
		audioFrequency:         frequency,
	}
}

func NewPolly(credPath string, profile string, region string) (*polly.Polly, error) {
	creds := credentials.NewSharedCredentials(credPath, profile)
	awsSess, err := session.NewSession(&aws.Config{Region: aws.String(region), Credentials:creds})

	if err != nil {
		return nil, err
	}

	return polly.New(awsSess), nil
}

type TTSPolly struct {
	polly                  *polly.Polly
	voiceId                string
	audioFrequency         int
}

func (t *TTSPolly) Speak(message string) (string, error) {
	filepath, err := t.textToWav(message)

	if err != nil {
		log.Println("Failed to save wav file for TTS:", err)
		return "", err
	}

	log.Println(message, "\t", filepath)

	return filepath, nil
}


// Takes a message, downloads an audio file that speaks that message
// and returns the local filename
func (t *TTSPolly) textToWav(message string) (string, error) {
	message = strings.Replace(message, "&", "&amp;", -1)

	openRe := regexp.MustCompile("(?i)<spell>")
	closeRe := regexp.MustCompile("(?i)</spell>")
	noSpk := regexp.MustCompile("(?i)<nospeak>(?:.|\n)+</nospeak>")

	message = noSpk.ReplaceAllString(message, "")
	message = openRe.ReplaceAllString(message, "<say-as interpret-as=\"spell-out\">")
	message = closeRe.ReplaceAllString(message, "</say-as>")
	message = "<speak>" + message + "</speak>"

	input := &polly.SynthesizeSpeechInput{
		OutputFormat: aws.String("pcm"),
		SampleRate:   aws.String(strconv.Itoa(t.audioFrequency)),
		Text:         aws.String(message),
		TextType:     aws.String("ssml"),
		VoiceId:      aws.String(t.voiceId),
	}

	result, err := t.polly.SynthesizeSpeech(input)

	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			case polly.ErrCodeTextLengthExceededException:
				log.Println(polly.ErrCodeTextLengthExceededException, aerr.Error())
			case polly.ErrCodeInvalidSampleRateException:
				log.Println(polly.ErrCodeInvalidSampleRateException, aerr.Error())
			case polly.ErrCodeInvalidSsmlException:
				log.Println(polly.ErrCodeInvalidSsmlException, aerr.Error())
			case polly.ErrCodeLexiconNotFoundException:
				log.Println(polly.ErrCodeLexiconNotFoundException, aerr.Error())
			case polly.ErrCodeServiceFailureException:
				log.Println(polly.ErrCodeServiceFailureException, aerr.Error())
			case polly.ErrCodeMarksNotSupportedForFormatException:
				log.Println(polly.ErrCodeMarksNotSupportedForFormatException, aerr.Error())
			case polly.ErrCodeSsmlMarksNotSupportedForTextTypeException:
				log.Println(polly.ErrCodeSsmlMarksNotSupportedForTextTypeException, aerr.Error())
			default:
				log.Println(aerr.Error())
			}
		} else {
			// Print the error, cast err to awserr.Error to get the Code and
			// Message from an error.
			log.Println(err.Error())
		}
		return "", err
	}

	defer result.AudioStream.Close()

	b, err := ioutil.ReadAll(result.AudioStream)

	if err != nil {
		log.Println("Failed to read the audio stream", err)
		return "", err
	}
	filepath := "/tmp/" + timeNow + ".wav"

	outFile, err := os.Create(filepath)
	// handle err
	defer outFile.Close()
	w := wav.NewWriter(
		outFile,
		uint32(len(b)),
		1,
		uint32(t.audioFrequency),
		16,
	)

	w.Write(b)

	if err != nil {
		log.Println("Error writing audio to file:", err)
		return "", err
	}

	return filepath, nil
}
