import os
from moviepy.editor import *

video = VideoFileClip(os.path.join("video4OkCZoSdYJ0.mp4"))
video.audio.write_audiofile(os.path.join("video4OkCZoSdYJ0.mp3"))
