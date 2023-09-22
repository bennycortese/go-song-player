from pytube import YouTube

# where to save
SAVE_PATH = "."  # to_do

# link of the video to be downloaded
link = "https://www.youtube.com/watch?v=LYRyKvEvVrA"

try:
    # object creation using YouTube
    # which was imported in the beginning
    yt = YouTube(link)
except:
    print("Connection Error")  # to handle exception

# filters out all the files with "mp4" extension
# get the video with the extension and
# resolution passed in the get() function

stream = yt.streams.filter(progressive=True, file_extension='mp4').first()
try:
    # downloading the video
    stream.download(SAVE_PATH)
except:
    print("Some Error!")
print('Task Completed!')
