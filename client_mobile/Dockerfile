FROM node:18

RUN apt-get update && apt-get install -y \
    openjdk-17-jdk \
    wget \
    unzip \
    gradle \
    cmake \
    dos2unix

ENV ANDROID_HOME=/opt/android-sdk
ENV PATH=$PATH:$ANDROID_HOME/cmdline-tools/latest/bin:$ANDROID_HOME/platform-tools

RUN mkdir -p ${ANDROID_HOME}/cmdline-tools && cd ${ANDROID_HOME}/cmdline-tools \
    && wget -q https://dl.google.com/android/repository/commandlinetools-linux-8092744_latest.zip \
    && unzip commandlinetools-linux-*_latest.zip -d ${ANDROID_HOME}/cmdline-tools \
    && rm commandlinetools-linux-*_latest.zip \
    && mv ${ANDROID_HOME}/cmdline-tools/cmdline-tools ${ANDROID_HOME}/cmdline-tools/latest

RUN yes | sdkmanager --licenses \
    && sdkmanager "platform-tools" "platforms;android-31" "build-tools;31.0.0"

RUN echo "sdk.dir=$ANDROID_HOME" > /root/local.properties

WORKDIR /app

COPY package*.json ./

COPY . .

RUN find /app -type f -exec dos2unix {} +

RUN npm install

RUN npm install -g expo-cli

RUN npx expo prebuild

RUN cp /root/local.properties ./android/local.properties

RUN cd android && ./gradlew assembleRelease --stacktrace --info

CMD ["sh", "-c", "cp /app/android/app/build/outputs/apk/release/app-release.apk /output/client.apk"]
