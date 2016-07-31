FROM iron/base
WORKDIR /app
# copy binary into image
COPY jwtea /app/
EXPOSE 8000
ENTRYPOINT ["./jwtea"]