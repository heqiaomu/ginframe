FROM alpha:latest
LABEL authors="sunyang"

ENTRYPOINT ["top", "-b"]