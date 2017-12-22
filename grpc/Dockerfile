FROM scratch

ARG wunderground_key

ENV REDIS_HOSTNAME redis.weather.svc
ENV WEATHER_UNDERGROUND_API_KEY $wunderground_key

COPY ./service /

CMD ["/service"]