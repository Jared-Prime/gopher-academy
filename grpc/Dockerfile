FROM scratch

ARG wunderground_key

ENV REDIS_HOSTNAME redis.weather.svc
ENV WEATHER_UNDERGROUND_API_KEY $wunderground_key

EXPOSE 8000

COPY ./service /

CMD ["/service"]