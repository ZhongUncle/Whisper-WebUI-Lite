# Whisper-WebUI-Lite
This is a whisper Web UI using pure Go, HTML,CSS,JS and Python, rather than Gradio like others structure.

I want to crate a simple tool, no showy UI. So you can use it as a simple tool. Or learn a simple structure of AI application.

This is still in developing. So it may have some bugs or issues. If you find one or want new features, please tell me.

I want to develop a simple tool, no any useless animation, so it looks like not cool.

## Usage
You can delay it in another device.

You need install FFMPEG and Whisper first:

```
sudo apt install ffmpeg
pip install -U openai-whisper
```

And you can use below command to run service:

```
go run main.go
```

Now you can use it in browser, link is http://localhost:8080, or `http://<IP-Address>:8080` when you run it in other device:

https://github.com/user-attachments/assets/b2012347-3ba6-4ccb-b21b-d161ea36c8c4

Yes, you can directly use video to convert.

If you want to upload new video/audio, please refresh page.

## Todos
- add clear button. (Now it via refresh page)
- add menus to output more format subtitles/text.
- develop REST APIs.
- fix Performance Issues.
- add copy button.
- add more informations in log, such as start time....

## Performance Issues
Now I use [Whisper](https://github.com/openai/whisper). it is python script, so every time execution will reload model. It is bad for performance.

