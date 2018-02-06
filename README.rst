===================================
Web Scrape Instagram Stories in Go_
===================================

.. image:: https://img.shields.io/badge/Language-Go-blue.svg
   :target: https://golang.org/

.. image:: https://godoc.org/github.com/siongui/goigstorydl?status.png
   :target: https://godoc.org/github.com/siongui/goigstorydl

.. image:: https://api.travis-ci.org/siongui/goigstorydl.png?branch=master
   :target: https://travis-ci.org/siongui/goigstorydl

.. image:: https://goreportcard.com/badge/github.com/siongui/goigstorydl
   :target: https://goreportcard.com/report/github.com/siongui/goigstorydl

.. image:: https://img.shields.io/badge/license-Unlicense-blue.svg
   :target: https://raw.githubusercontent.com/siongui/goigstorydl/master/UNLICENSE

.. image:: https://img.shields.io/badge/Status-Beta-brightgreen.svg

.. image:: https://img.shields.io/twitter/url/https/github.com/siongui/goigstorydl.svg?style=social
   :target: https://twitter.com/intent/tweet?text=Wow:&url=%5Bobject%20Object%5D


Download Instagram_ stories in Go_. Currently only Linux systems are supported
because wget_ is called via ``os/exec`` package.


Obtain Cookies
++++++++++++++

The following three values are must to access the Instagram story API.

- ``ds_user_id``
- ``sessionid``
- ``csrftoken``

First login to Instagram_ from Chrome browser, and there are two ways to get the
above information:

1. From `Chrome Developer Tools`_: See this `SO answer`_ or `Obtain cookies`_
   section in `instastories-backup`_ repo.
2. From Chrome extension: Use EditThisCookie_ or `cookie-txt-export`_ or other
   cookie tools.


Example
+++++++

.. code-block:: go

  package main

  import (
  	"github.com/siongui/goigstorydl"
  )

  func main() {
  	igstorydl.MonitorAndDownload("your user id", "your session id", "your csrftoken")
  }


UNLICENSE
+++++++++

Released in public domain. See UNLICENSE_.


References
++++++++++

.. [1] `GitHub - siongui/goigstorylink: Get Links (URL) of Instagram Stories in Go <https://github.com/siongui/goigstorylink>`_


.. _Go: https://golang.org/
.. _UNLICENSE: http://unlicense.org/
.. _Web Scrape: https://www.google.com/search?q=Web+Scrape
.. _EditThisCookie: https://www.google.com/search?q=EditThisCookie
.. _cookie-txt-export: https://github.com/siongui/cookie-txt-export.go
.. _Obtain cookies: https://github.com/hoschiCZ/instastories-backup#obtain-cookies
.. _instastories-backup: https://github.com/hoschiCZ/instastories-backup
.. _Chrome Developer Tools: https://developer.chrome.com/devtools
.. _SO answer: https://stackoverflow.com/a/44773079
.. _Instagram: https://www.instagram.com/
.. _wget: https://www.gnu.org/software/wget/
