FROM python:3.6.10-buster
COPY requirements.txt /
RUN pip3 install -r /requirements.txt
COPY . /
EXPOSE 8081
ENTRYPOINT ["./gunicorn.sh"]