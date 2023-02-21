FROM alpine
ADD app.py /apps/
ADD start.sh /root/
WORKDIR /apps
RUN mkdir /apps ; apk update; apk add python-3; apk add py3-pip ; pip install flask; chmod +x /root/start.sh
CMD ["flask", "run", "--host=0.0.0.0"]