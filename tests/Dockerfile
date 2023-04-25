FROM python:latest

RUN mkdir /minitwit-tests
RUN pip install --upgrade pip
RUN pip install psycopg2
RUN pip install requests

COPY . /tests

WORKDIR /tests


