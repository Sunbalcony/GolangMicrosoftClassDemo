FROM registry.cn-shenzhen.aliyuncs.com/base_repo/java_8u212_cst:v1.0.0

# Add Maintainer Info
MAINTAINER xuyx <xuyuxia@e6yun.com>

# 设置locale
ENV LANG en_US.UTF-8
ENV LANGUAGE en_US:en
ENV LC_ALL en_US.UTF-8
ENV TZ=Asia/Shanghai

RUN mkdir /app 

WORKDIR /app

COPY ./target/javademo.jar /app/javademo.jar

ENTRYPOINT ["java", "-Dserver.port=8080", "-jar","/app/javademo.jar"]





