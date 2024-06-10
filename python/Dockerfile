FROM python:3.11

ENV PYTHONDONTWRITEBYTECODE 1
ENV PYTHONUNBUFFERED 1

WORKDIR /python

COPY requirements.txt /python/

RUN pip install --no-cache-dir -r requirements.txt

COPY . /python/

CMD [ "python", "main.py" ]