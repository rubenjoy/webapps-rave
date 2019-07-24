FROM openjdk:8u131-jre-alpine

ARG JMETER_ZIP
ARG JMETER_DIR

WORKDIR /tools/$JMETER_DIR
ENV JMETER_HOME=/tools/$JMETER_DIR

COPY $JMETER_ZIP /tools/
RUN cd /tools && \
    unzip $JMETER_ZIP && \
    rm -v $JMETER_ZIP

ONBUILD ADD . /tools

ENTRYPOINT ["bin/jmeter.sh"]