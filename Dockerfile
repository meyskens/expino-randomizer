ARG arch
FROM multiarch/alpine:${arch}-edge

COPY ./expino-randomizer /expino-randomizer

ENV ENDPOINT="http://api;8080"

CMD /expino-randomizer