FROM openjdk:11-slim

RUN apt-get update && apt-get install -y \
        wget \
        unzip \
    && apt-get clean \
    && rm -rf /var/lib/apt

RUN wget https://github.com/mybatis/migrations/releases/download/mybatis-migrations-3.3.9/mybatis-migrations-3.3.9-bundle.zip \
    && unzip mybatis-migrations-3.3.9-bundle.zip -d /usr/local \
    && rm mybatis-migrations-3.3.9-bundle.zip

ENV MYBATIS_MIGRATIONS_HOME /usr/local/mybatis-migrations-3.3.9
ENV PATH $PATH:$MYBATIS_MIGRATIONS_HOME/bin

WORKDIR /mybatis

COPY . .
CMD ["/bin/bash"]
