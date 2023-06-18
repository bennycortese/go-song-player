package main

import (
	"io"
	"os"

	"github.com/kkdai/youtube/v2"
	ffmpeg "github.com/u2takey/ffmpeg-go"

	"github.com/hajimehoshi/oto"

	"github.com/hajimehoshi/go-mp3"
)

func videoDownload(videoID string, youtubeClient youtube.Client) {
	video, err := youtubeClient.GetVideo(videoID)

	if err != nil {
		panic(err)
	}

	formats := video.Formats.WithAudioChannels()
	stream, _, err := youtubeClient.GetStream(video, &formats[0])

	if err != nil {
		panic(err)
	}

	file, err := os.Create("video" + videoID + ".mp4")
	if err != nil {
		panic(err)
	}

	defer file.Close()

	_, err = io.Copy(file, stream)
	if err != nil {
		panic(err)
	}
}

func videoFile(testFileLocation string) {
	err := ffmpeg.Input("./"+testFileLocation, ffmpeg.KwArgs{"ss": "1"}).
		Output("./out1.gif", ffmpeg.KwArgs{"s": "320x240", "pix_fmt": "rgb24", "t": "3", "r": "3"}).
		OverWriteOutput().ErrorToStdOut().Run()
	if err != nil {
		panic(err)
	}
}

func playSong() {
	file, err := os.Open("video4OkCZoSdYJ0.mp3")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	decoder, err := mp3.NewDecoder(file)
	if err != nil {
		panic(err)
	}

	context, err := oto.NewContext(decoder.SampleRate(), 2, 2, 8192)
	if err != nil {
		panic(err)
	}
	defer context.Close()

	player := context.NewPlayer()
	defer player.Close()

	if _, err := io.Copy(player, decoder); err != nil {
		panic(err)
	}

}

func main() {
	/*videoID := "4OkCZoSdYJ0"
	client := youtube.Client{}
	videoDownload(videoID, client)*/
	//python_converter := exec.Command("python", "mp4_to_mp3.py")
	//python_converter.Start()
	//python_converter.Wait()
	playSong()
}
