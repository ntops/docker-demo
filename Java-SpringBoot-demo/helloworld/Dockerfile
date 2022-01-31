FROM openjdk
ADD helloworld.jar helloworld.jar
RUN bash -c 'touch /helloworld.jar'
ENTRYPOINT ["java","-jar","/helloworld.jar"]
