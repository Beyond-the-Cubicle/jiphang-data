FROM gradle:8-jdk21 AS builder
WORKDIR /builder
COPY . .
RUN gradle clean bootJar

FROM openjdk:21
WORKDIR /application
COPY --from=builder ./builder/build/libs/route-information-0.0.1-SNAPSHOT.jar app.jar
RUN ls -al
ENTRYPOINT ["java","-jar","app.jar"]