FROM python:3.7-alpine3.11

COPY . /app

WORKDIR /app

RUN apk update && apk add gcc python3-dev musl-dev

RUN pip install -r requirements.txt

EXPOSE 5500

CMD [ "python", "./app.py" ]