FROM php:8-fpm-alpine

ADD https://github.com/mlocati/docker-php-extension-installer/releases/latest/download/install-php-extensions /usr/local/bin/

RUN chmod +x /usr/local/bin/install-php-extensions && \
    install-php-extensions exif gd imagick pdo_mysql mysqli zip intl

ARG INI_NAME

RUN mv "$PHP_INI_DIR/$INI_NAME" "$PHP_INI_DIR/php.ini"

RUN wget https://getcomposer.org/composer-stable.phar -O /usr/local/bin/composer && chmod +x /usr/local/bin/composer

RUN install-php-extensions pdo_pgsql pgsql

COPY ./custom.ini /usr/local/etc/php/conf.d/custom.ini
COPY ./zz-docker.conf /usr/local/etc/php-fpm.d/zz-docker.conf

RUN chown -R www-data:www-data /var/www



WORKDIR /var/www

#RUN chown -R 755 /var/www
#RUN chown -R 755 /var/www/courses/bootstrap