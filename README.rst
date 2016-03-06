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

Running the server
==================
 
  There is a `Demo server <http://ec2-54-194-191-17.eu-west-1.compute.amazonaws.com:18001/
>`_.
  
  docker run --restart=always -p 18001:8001 --name edifyweb bbiskup/edifyweb_dev ."/edify-web run -H 0.0.0.0"

Development
===========

Travis CI Setup
+++++++++++++++

The following environment variables need to be set to pull/push from/to Docker Hub:

- ``DOCKER_USERNAME``
- ``DOCKER_PASSWORD``
- ``DOCKER_EMAIL``
