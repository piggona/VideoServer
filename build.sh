#! /bin/bash

#build web UI
cd /Users/haohao/Documents/go/src/video_server/web
go install
cp /Users/haohao/Documents/go/bin/web /Users/haohao/Documents/go/bin/video_server_web_ui/web
cp -R /Users/haohao/Documents/go/src/video_server/templates /Users/haohao/Documents/go/bin/video_server_web_ui/templates