edify-web
=========

Simple web frontend to the edify EDIFACT toolkit


Build Status
============

====== ===============
Branch Status
====== ===============
dev    |travis-dev|
====== ===============

.. |travis-dev| image:: https://travis-ci.org/bbiskup/edify-web.svg?branch=dev
        :target: https://travis-ci.org/bbiskup/edify-web

Travis CI Setup
===============

The following environment variables need to be set to pull/push from/to Docker Hub:

- ``DOCKER_USERNAME``
- ``DOCKER_PASSWORD``
- ``DOCKER_EMAIL``

Run server in container
=======================
  
  docker run --restart=always -p 18001:8001 --name edifyweb bbiskup/edifyweb_dev ."/edify-web run -H 0.0.0.0"