import whisper
import time

print("Loading model...")
model = whisper.load_model("turbo")
print("Model loaded")

start = time.time()
print("Running task...")
result = model.transcribe("./uploads/up_file")
print(result["text"])
end = time.time()

print("Finished task")
print("Run time:%d", end - start)

from whisper.utils import get_writer

writer = get_writer("srt", "converted")
writer(result, "cv.srt")

print("Wrote to cv.srt")