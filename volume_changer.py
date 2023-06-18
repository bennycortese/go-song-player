from pydub import AudioSegment
import os

song = AudioSegment.from_mp3(os.path.join("video4OkCZoSdYJ0.mp3"))

quieter_song = song - 3

quieter_song.export("video4OkCZoSdYJ0_quieter.mp3", format='mp3')
