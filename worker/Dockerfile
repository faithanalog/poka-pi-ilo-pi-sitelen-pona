FROM "ubuntu:18.04"

RUN apt-get update \
        && apt-get install -y phantomjs \
        && rm -r /var/lib/apt/lists/

# Required to make phantomjs run properly
ENV QT_QPA_PLATFORM=offscreen
WORKDIR /app
COPY . .
CMD phantomjs render.js

