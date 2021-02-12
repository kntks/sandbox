FROM fluent/fluentd:v1.12-1

# https://github.com/fluent/fluentd-docker-image#how-to-build-your-own-image
# Use root account to use apk
USER root

RUN apk add --no-cache --update --virtual .build-deps \
  sudo build-base ruby-dev && \
  sudo gem install \
  fluent-plugin-elasticsearch \
  fluent-plugin-cloudwatch-logs \
  fluent-plugin-s3 \
  fluent-plugin-kinesis \
  && sudo gem sources --clear-all \
  && apk del .build-deps \
  && rm -rf /tmp/* /var/tmp/* /usr/lib/ruby/gems/*/cache/*.gem

# COPY fluent.conf /fluentd/etc/
COPY entrypoint.sh /bin/

USER fluent