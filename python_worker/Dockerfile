FROM python:3.12-alpine
RUN apk add --no-cache bash util-linux 
RUN adduser -D -u 1000 judgeuser
WORKDIR /sandbox
USER judgeuser
CMD ["python", "submission.py"]

