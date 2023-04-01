FROM ubuntu:latest

COPY ./kubernetes-controller ./kubernetes-controller
RUN chmod +x ./kubernetes-controller

ENTRYPOINT ["./kubernetes-controller"]
